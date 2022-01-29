package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Host string
	Port int
}

var cfgFile string
var appConfig AppConfig

func GetAppConfig() AppConfig {
	return appConfig
}

var rootCmd = &cobra.Command{
	Use:   "condog",
	Short: "コマンドの説明文（短い）",
	Long:  "コマンドの説明文（長文）",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	// フラグと構成設定を定義
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "設定ファイル (デフォルト値は $HOME/.config.yaml)")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		// 設定ファイルが存在するならセット
		viper.SetConfigFile(cfgFile)
	} else {
		// Homeディレクトリを取得
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// ホームディレクトリから「config.yaml」という設定ファイルを検索
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	// 環境変数の読み込み
	viper.AutomaticEnv() 

  // 設定ファイルが見つからなかった場合はエラーを返却
  if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
