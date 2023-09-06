package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Environment string

func LoadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

type Config struct {
	Database struct {
		Dsn     string `mapstructure:"dsn"`
		DB_Name string `mapstructure:"db_name"`
		MY_SQL  string `mapstructure:"my_sql"`
	} `mapstructure:"database"`
	Redis struct {
		Dsn string `mapstructure:"dsn"`
	} `mapstructure:"database"`
	JWT struct {
		Secret           string `mapstructure:"secret"`
		ExpiryMinAccess  int    `mapstructure:"expiry"`
		ExpiryMinRefresh int    `mapstructure:"expiry_refresh"`
		AccessSecret     string `mapstructure:"access_secret"`
		RefreshSecret    string `mapstructure:"refresh_secret"`
		AnotherSecret    string `mapstructure:"another_secret"`
	} `mapstructure:"jwt"`
	Pagination struct {
		Offset int `mapstructure:"offset"`
		Limit  int `mapstructure:"limit"`
	} `mapstructure:"pagination"`
	App struct {
		Name            string `mapstructure:"name"`
		Version         string `mapstructure:"version"`
		Env             string `mapstructure:"env"`
		Port            string `mapstructure:"port"`
		UIEndpoint      string `mapstructure:"ui_endpoint"`
		PageUrl         string `mapstructure:"page_url"`
		Hostname        string `mapstructure:"hostname"`
		WebhookHostName string `mapstructure:"webhook_hostname"`
		CustomerPWD     string `mapstructure:"customer_pwd"`
		LocationID      string `mapstructure:"location_id"`
		Issuer          string `mapstructure:"issuer"`
	} `mapstructure:"app"`
	Service struct {
		User struct {
			Url string `mapstructure:"url"`
		} `mapstructure:"user"`
		Product struct {
			Url string `mapstructure:"url"`
		} `mapstructure:"product"`
		Order struct {
			Url string `mapstructure:"url"`
		} `mapstructure:"order"`
		Seller struct {
			Url string `mapstructure:"url"`
		} `mapstructure:"seller"`
		Email struct {
			Url string `mapstructure:"url"`
		} `mapstructure:"email"`
		Connector struct {
			Url string `mapstructure:"url"`
		} `mapstructure:"connector"`
		Document struct {
			Url string `mapstructure:"url"`
		} `mapstructure:"document"`
	} `mapstructure:"services"`
	ServiceBus struct {
		ConnStr string `mapstructure:"conn_str"`
	} `mapstructure:"service_bus"`
	SendGrid struct {
		SenderName     string `mapstructure:"sender_name"`
		SenderEmailID  string `mapstructure:"sender_email_id"`
		SendGridAPIKey string `mapstructure:"sendgrid_api_key"`
	} `mapstructure:"sendgrid"`
	Storage struct {
		AccountName   string `mapstructure:"account_name"`
		AccountKey    string `mapstructure:"account_key"`
		Container     string `mapstructure:"container"`
		MaxFileSizeKb int    `mapstructure:"max_file_size_kb"`
	} `mapstructure:"storage"`
	Email struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		From     string `mapstructure:"from"`
	} `mapstructure:"email"`
	Sentry struct {
		Dsn string `mapstructure:"dsn"`
	} `mapstructure:"sentry"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	// env := os.Getenv("GOENV")
	once.Do(func() {
		instance = &Config{}
		if err := initConfig(Environment(os.Getenv("GOENV"))); err != nil {
			log.Fatalf("error initializing config: %v", err)
		}
	})
	return instance
}

func initConfig(env Environment) error {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>> Environment: ", env)
	if env == "" {
		return errors.New("environment not set")
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(string(env))
	v.AddConfigPath("./config/")
	v.AddConfigPath("config")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("configuration file not found")
		}
		return errors.New("error on parsing configuration file")
	}

	if err := v.Unmarshal(instance); err != nil {
		return errors.New("error on parsing configuration file to struct")
	}

	v.WatchConfig()

	return nil
}
