package model

import (
	"errors"

	"github.com/mikokutou1/go-zero-m/core/stores/mon"
)

var (
	ErrNotFound        = mon.ErrNotFound
	ErrInvalidObjectId = errors.New("invalid objectId")
)
