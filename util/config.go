package util

import(
	"github.com/spf13/viper"
	"time"
)
// will hold all the config of the applicatio
// the values are read by viper from the config file or environment variable
type Config struct{
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey string `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

// loadconfig reads the configuration from the file or env variable 
func LoadConfig(path string) (config Config, err error){
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") 
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil{
		return
	}

	err = viper.Unmarshal(&config)
	return 
}