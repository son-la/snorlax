package cmd

import (
	"github.com/son-la/snorlax/internal/kafka"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"log"
	"time"
)

func serveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "serve",
		Short:   "Start a inifite loop and send current server timestamp to Kafka queue every 5 seconds",
		Aliases: []string{"s"},
		//Args:    cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			var appConfig AppConfig
			if err := loadConfigFromFile(&appConfig); err != nil {
				log.Fatal(err)
				return
			}

			kafkaConfig := kafka.KafkaConfig{
				Brokers: appConfig.Kafka.Brokers,
				UseTLS:  appConfig.Kafka.UseTLS,
				CAFile:  appConfig.Kafka.CAFile,
				Topic:   appConfig.Kafka.Topic,
				Authentcation: kafka.Authentcation{
					Username:  appConfig.Kafka.Authentcation.Username,
					Password:  appConfig.Kafka.Authentcation.Password,
					Algorithm: appConfig.Kafka.Authentcation.Algorithm,
				},
				Version: appConfig.Kafka.Version,
			}

			producer, err := kafka.InitKafka(&kafkaConfig)
			if err != nil {
				log.Fatal(err)
			}

			d := time.Duration(5 * time.Second)
			for {
				kafka.SendMessage(producer, appConfig.Kafka.Topic, time.Now().String())
				time.Sleep(d)
			}
		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(serveCmd())
}

func loadConfigFromFile(appConfig *AppConfig) error {
	// TODO: Load authentication data from env variable

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
		return err
	}

	if err := viper.Unmarshal(&appConfig); err != nil {
		return err
		log.Println(err)
	}

	return nil
}
