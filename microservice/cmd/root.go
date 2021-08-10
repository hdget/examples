//指令相关模块
package cmd

import (
	"context"
	"examples/alidts/g"
	"examples/microservice/autogen"
	"examples/microservice/service"
	"examples/microservice/service/handler"
	"github.com/hdget/sdk"
	"github.com/hdget/sdk/utils"
	"github.com/spf13/cobra"
	"os"
)

const (
	APP = "microservice"
)

type searchServiceImpl struct{}

func (s searchServiceImpl) Search(ctx context.Context, request *autogen.SearchRequest) (*autogen.SearchResponse, error) {
	return &autogen.SearchResponse{
		Response: "hello world",
	}, nil
}

var (
	env        string // 命令行指定的环境
	configFile string
)

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   APP,
	Short: "demonstrate microservice",
	Long:  `demonstrate microservice`,
	Run: func(cmd *cobra.Command, args []string) {
		err := sdk.Initialize(g.Config)
		if err != nil {
			utils.Fatal("sdk initialize", "err", err)
		}

		// 必须手动注册服务实现
		svc := &service.SearchServiceImpl{}
		srv := sdk.MicroService.By("testservice").CreateServer()
		endpoints := &service.Endpoints{
			srv.CreateEndpointServer(svc, &handler.SearchHandler{}),
			srv.CreateEndpointServer(svc, &handler.HelloHandler{}),
		}
		autogen.RegisterSearchServiceServer(srv.GetGrpcServer(), endpoints)
		srv.Run()
	},
}

func register() {

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
