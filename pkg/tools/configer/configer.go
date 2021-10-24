package configer

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type OAuth2AppData struct {
	Keycloak KeycloakAppData `mapstructure:"keycloak-app-data"`
}

type KeycloakAppData struct {
	ClientID     string `mapstructure:"client-id"`
	ClientSecret string `mapstructure:"client-secret"`
	RedirectURL  string `mapstructure:"redirect-url"`
	ConfigURL  string `mapstructure:"config-url"`
}

type ServerData struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type RedisData struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
}

type PostgresqlData struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db-name"`
	Port     string `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	Sslmode  string `mapstructure:"sslmode"`
}

type SecretData struct {
	AccessSecret  string `mapstructure:"access-secret"`
	RefreshSecret string `mapstructure:"refresh-secret"`
}

type Config struct {
	OAuth2App  OAuth2AppData  `mapstructure:"oauth2-app-data"`
	Server     ServerData     `mapstructure:"server-data"`
	Redis      RedisData      `mapstructure:"redis-data"`
	Postgresql PostgresqlData `mapstructure:"postgresql-data"`
	Secret     SecretData     `mapstructure:"secrets-data"`
}

var AppConfig Config

func Init(configPath string) {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatal(err)
	}

	fmt.Println(AppConfig)
}
