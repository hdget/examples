package g

import "github.com/hdget/sdk"

type XxxServiceConfig struct {
	sdk.Config `mapstructure:",squash"`
}

// 全局配置信息
var (
	Config *XxxServiceConfig
)
