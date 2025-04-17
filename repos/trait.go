package repos

import (
	"go-web-project/core"
)

type CuteDBParameter interface {
	GetKey() string
	GetValue() interface{}
	GetSQL() string
}

type CuteDBOutput interface {
	SetOutput(row map[string]interface{})
}

type CuteDB interface {
	Create(config CuteDBConfig) *core.CuteError
	Close() *core.CuteError
	Set(input CuteDBParameter) *core.CuteError
	Get(input CuteDBParameter, refData *[]CuteDBOutput, factory func() CuteDBOutput) *core.CuteError
	Delete(input CuteDBParameter) *core.CuteError
}
