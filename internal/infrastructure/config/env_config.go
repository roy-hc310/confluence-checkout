package config

import (
	"reflect"

	// "strings"

	"github.com/spf13/viper"
)

var GlobalEnv Env

type Env struct {
	Port           string `mapstructure:"PORT" json:"PORT"`
	Debugging      bool   `mapstructure:"DEBUGGING"`
	ContextTimeOut int    `mapstructure:"CONTEXT_TIMEOUT"`

	GRPCHost string `mapstructure:"GRPC_HOST"`
	GRPCPort string `mapstructure:"GRPC_PORT"`

	GRPCProductHost string `mapstructure:"GRPC_PRODUCT_HOST"`
	GRPCPaymentHost string `mapstructure:"GRPC_PAYMENT_HOST"`

	OtelHttpExporter string `mapstructure:"OTEL_HTTP_EXPORTER"`
	OtelGrpcExporter string `mapstructure:"OTEL_GRPC_EXPORTER"`

	DBReadHost   string `mapstructure:"DB_READ_HOST"`
	DBReadPort   string `mapstructure:"DB_READ_PORT"`
	DBReadName   string `mapstructure:"DB_READ_NAME"`
	DBReadUser   string `mapstructure:"DB_READ_USER"`
	DBReadPass   string `mapstructure:"DB_READ_PASS"`
	DBReadSchema string `mapstructure:"DB_READ_SCHEMA"`

	DBWriteHost   string `mapstructure:"DB_WRITE_HOST"`
	DBWritePort   string `mapstructure:"DB_WRITE_PORT"`
	DBWriteName   string `mapstructure:"DB_WRITE_NAME"`
	DBWriteUser   string `mapstructure:"DB_WRITE_USER"`
	DBWritePass   string `mapstructure:"DB_WRITE_PASS"`
	DBWriteSchema string `mapstructure:"DB_WRITE_SCHEMA"`

	KafkaHost          string `mapstructure:"KAFKA_HOST"`
	KafkaConsumerGroup string `mapstructure:"KAFKA_CONSUMER_GROUP"`

	RedisHost    string `mapstructure:"REDIS_HOST_PORT"`
	RedisPass    string `mapstructure:"REDIS_PASS"`
	RedisDB      string `mapstructure:"REDIS_DB"`
	RedisTimeOut int64  `mapstructure:"REDIS_TIMEOUT"`

	ElasticHost string `mapstructure:"ELASTIC_HOST"`
	ElasticUser string `mapstructure:"ELASTIC_USER"`
	ElasticPass string `mapstructure:"ELASTIC_PASS"`
}

func LoadGlobalEnv(path string) (err error) {

	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	globalEnvType := reflect.TypeOf(Env{})
	for i := 0; i < globalEnvType.NumField(); i++ {
		field := globalEnvType.Field(i)
		fieldTag := field.Tag.Get("mapstructure")
		if fieldTag != "" {
			viper.BindEnv(fieldTag)
		}
	}

	viper.Unmarshal(&GlobalEnv)
	return nil
}
