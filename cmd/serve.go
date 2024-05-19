package cmd

import (
	"github.com/son-la/snorlax/internal/kafka"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"log"

	"github.com/gin-gonic/gin"


    "github.com/son-la/snorlax/internal/handlers"
    "github.com/son-la/snorlax/internal/middleware"
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

			_, err := kafka.InitKafka(&kafkaConfig)
			if err != nil {
				log.Println(err)
			}

			// d := time.Duration(5 * time.Second)
			// for {
			// 	kafka.SendMessage(producer, appConfig.Kafka.Topic, time.Now().String())
			// 	time.Sleep(d)
			// }

			r := gin.Default()

			// Public routes (do not require authentication)
			publicRoutes := r.Group("/public")
			{
				publicRoutes.POST("/login", handlers.Login)
				publicRoutes.POST("/register", handlers.Register)
			}

			// Protected routes (require authentication)
			protectedRoutes := r.Group("/protected")
			protectedRoutes.Use(middleware.AuthenticationMiddleware())
			{
				// Protected routes here
			}

			r.Run(":8080")
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
