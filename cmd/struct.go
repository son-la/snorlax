package cmd

type AppConfig struct {
	Kafka    KafkaConfig    `mapstructure:"kafka"`
	Database DatabaseConfig `mapstructure:"database"`
}

type KafkaConfig struct {
	Brokers []string `mapstructure:"brokers"`
	UseTLS  bool     `mapstructure:"useTLS"`
	CAFile  string   `mapstructure:"caFile"`
	Version string   `mapstructure:"version"`

	Authentcation Authentcation `mapstructure:"authentcation"`
	Topic         string        `mapstructure:"topic"`
}

type Authentcation struct {
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	Algorithm string `mapstructure:"algorithm"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}
