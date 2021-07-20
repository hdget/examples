//指令相关模块
package cmd

import (
	"examples/alidts/g"
	"github.com/hdget/sdk"
	"github.com/hdget/sdk/types"
	"github.com/hdget/sdk/utils"
	"github.com/spf13/cobra"
	"log"
	"os"
)

const (
	APP = "alidts"
)

var (
	env        string // 命令行指定的环境
	configFile string
)

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   APP,
	Short: "sync something",
	Long:  `sync something from db`,
	Run: func(cmd *cobra.Command, args []string) {
		err := sdk.Initialize(g.Config)
		if err != nil {
			utils.Fatal("sdk initialize", "err", err)
		}

		c, err := sdk.Kafka.By("aliyun").CreateConsumer("syncsomething", dtsHandler)
		if err != nil {
			sdk.Logger.Fatal("create consumer", "err", err)
		}

		c.Consume()
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&env, "env", "e", "", "running environment, e,g: [prod, sim, pre, test, dev, local]")
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "f", "", "config file, default: config.toml")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// 尝试捕获panic并保存到错误中
	defer func() {
		if r := recover(); r != nil {
			utils.RecordErrorStack(APP)
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	// 尝试从各种源加载配置信息
	v := sdk.LoadConfig(APP, env, configFile)

	// 将配置信息转换成对应的数据结构
	err := v.Unmarshal(&g.Config)
	if err != nil {
		utils.Fatal("msg", "unmarshal config", "err", err)
	}
}

func dtsHandler(data []byte) types.MqMsgAction {
	log.Println(utils.BytesToString(data))
	return types.Ack
}
