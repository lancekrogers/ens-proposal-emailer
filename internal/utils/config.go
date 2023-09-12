package utils

import (
	"strings"
	"tally-takehome/internal/email"
	"tally-takehome/internal/tally"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const DB_PATH = "./internal/store/store.db"

type Config struct {
	TallyApi              tally.TallyApi
	EmailSettings         email.EmailSettings
	ENSGovernanceContract Contract
	MainnetRPC            string
}

type Contract struct {
	Address string
	AbiPath string
}

func LoadConfig[T interface{}](path string, fileName string) (*T, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var c *T
	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot decode into struct")
	}
	return c, nil
}
