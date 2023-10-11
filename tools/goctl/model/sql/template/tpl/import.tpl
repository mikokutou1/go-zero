import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	{{if .containsPQ}}"github.com/lib/pq"{{end}}
	"github.com/mikokutou1/go-zero-m/core/stores/builder"
	"github.com/mikokutou1/go-zero-m/core/stores/cache"
	"github.com/mikokutou1/go-zero-m/core/stores/sqlc"
	"github.com/mikokutou1/go-zero-m/core/stores/sqlx"
	"github.com/mikokutou1/go-zero-m/core/stringx"
)
