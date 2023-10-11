package nacos

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/pkg/errors"
	"google.golang.org/grpc/resolver"
)

func init() {
	resolver.Register(&builder{})
}

// schemeName for the urls
// All target URLs like 'nacos://.../...' will be resolved by this resolver
const schemeName = "nacos"

// builder implements resolver.Builder and use for constructing all consul resolvers
type builder struct{}

// Build 构造
func (b *builder) Build(url resolver.Target, conn resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	tgt, err := parseURL(url.URL.String())
	if err != nil {
		return nil, errors.Wrap(err, "Wrong nacos URL")
	}

	host, ports, err := net.SplitHostPort(tgt.Addr)
	if err != nil {
		return nil, fmt.Errorf("failed parsing address error: %v", err)
	}
	port, _ := strconv.ParseUint(ports, 10, 16)

	var gPort uint64
	if tgt.Grpc == "" {
		gPort = port + 1000
	} else {
		gPort, _ = strconv.ParseUint(tgt.Grpc, 10, 16)
	}

	sc := []constant.ServerConfig{
		*constant.NewServerConfig(host, port, constant.WithContextPath("/nacos"), constant.WithGrpcPort(gPort)),
	}

	// create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(tgt.NamespaceID),   // NameSpaceID，显示在 Nacos UI 中
		constant.WithTimeoutMs(uint64(tgt.Timeout)), // 超时
		constant.WithNotLoadCacheAtStart(true),      // 不加载缓存
		constant.WithUsername(tgt.User),             // 账号，Nacos 登录使用
		constant.WithPassword(tgt.Password),         // 密码  Nacos 登录使用
		constant.WithLogLevel("info"),               // 日志级别
		constant.WithAppName(tgt.AppName),           // 订阅者名称，显示在 Nacos UI 中
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
	)
	// cc := &constant.ClientConfig{
	// 	AppName:             tgt.AppName,         // 订阅者名称，显示在 Nacos UI 中
	// 	NamespaceId:         tgt.NamespaceID,     // NameSpaceID，显示在 Nacos UI 中
	// 	Username:            tgt.User,            // 账号，Nacos 登录使用
	// 	Password:            tgt.Password,        // 密码  Nacos 登录使用
	// 	TimeoutMs:           uint64(tgt.Timeout), // timout
	// 	NotLoadCacheAtStart: true,
	// }

	if tgt.CacheDir != "" {
		cc.CacheDir = tgt.CacheDir
	}
	if tgt.LogDir != "" {
		cc.LogDir = tgt.LogDir
	}
	if tgt.LogLevel != "" {
		cc.LogLevel = tgt.LogLevel
	}
	if tgt.Clusters != nil {
		tgt.Clusters = []string{"DEFAULT"}
	}

	cli, err := clients.NewNamingClient(vo.NacosClientParam{
		ServerConfigs: sc,
		ClientConfig:  &cc,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't connect to the nacos API")
	}

	ctx, cancel := context.WithCancel(context.Background())
	pipe := make(chan []string)

	param := &vo.SubscribeParam{
		ServiceName:       tgt.Service,
		Clusters:          tgt.Clusters,
		GroupName:         tgt.GroupName,
		SubscribeCallback: newWatcher(ctx, cancel, pipe).CallBackHandle,
	}

	go cli.Subscribe(param)

	go populateEndpoints(ctx, conn, pipe)

	return &resolvr{cancelFunc: cancel}, nil
}

// Scheme returns the scheme supported by this resolver.
// Scheme is defined at https://github.com/grpc/grpc/blob/master/doc/naming.md.
func (b *builder) Scheme() string {
	return schemeName
}
