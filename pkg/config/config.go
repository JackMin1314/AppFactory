package config

import (
	"path/filepath"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"gopkg.in/yaml.v2"
)

// var once sync.Once
// // InitConfig 初始化读取配置文件,成功会返回 config
// func InitConfig(confName string){
// 	once.Do(func(){
// 		if confName == "" {
// 			confName = "../config/config.yaml"
// 		}
// 		absConfigPath, err := filepath.Abs(confName)
// 		if err != nil {
// 			panic(err)
// 		}
// 		// 获取配置文件
// 		config := config.New(
// 			config.WithSource(
// 				file.NewSource(absConfigPath),
// 			),
// 			config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
// 				return yaml.Unmarshal(kv.Value, v)
// 			}),
// 		)
// 		if err := config.Load(); err != nil {
// 			panic(err)
// 		}

// 	})

// }
type ApplicationYaml struct {
	Name string
	Env string
	Host string
	Port int
}
type ServerYaml struct {
	Http struct {
		Addr    string
		Timeout string
	}
	Grpc struct {
		Addr    string
		Timeout string
	}
}
type DataYaml struct {
	DataBase struct {
		Driver string
		Source string
	}
	Redis struct {
		Addr         string
		DialTimeout  string
		ReadTimeout  string
		WriteTimeout string
	}
}
type LogYaml struct {
	FileFolder string
	FileName   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}
type JwtYaml struct {
	ExpiresAt int
	RefreshAt int
	SecretKey string
}
type CasbinYaml struct {
	ModelFile string
}

type ConfigYaml struct {
	Application ApplicationYaml
	Server      ServerYaml
	Data        DataYaml
	Log         LogYaml
	JWT         JwtYaml
	Casbin      CasbinYaml
}

func NewConfigYaml() *ConfigYaml {
	return new(ConfigYaml)
}

// InitConfigYaml 初始化读取配置文件,成功会返回 *ConfigYaml
func InitConfigYaml(confName string) *ConfigYaml {

	config := InitConfig(confName)
	cfg := NewConfigYaml()
	if err := config.Scan(cfg); err != nil {
		panic(err)
	}
	return cfg

}

// InitConfig 初始化读取配置文件，成功返回config.Config
func InitConfig(confName string) config.Config {
	if confName == "" {
		confName = "../config/config.yaml"
	}
	absConfigPath, err := filepath.Abs(confName)
	if err != nil {
		panic(err)
	}
	// 获取配置文件
	config := config.New(
		config.WithSource(
			file.NewSource(absConfigPath),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)
	if err := config.Load(); err != nil {
		panic(err)
	}
	return config
}
