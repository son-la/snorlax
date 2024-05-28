package cmd

import (
	"fmt"

	"github.com/son-la/snorlax/internal/controllers"
	"github.com/son-la/snorlax/internal/database"
	middleware "github.com/son-la/snorlax/internal/middlewares"
	"github.com/son-la/snorlax/internal/repositories"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"log"

	"github.com/gin-gonic/gin"
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

			// kafkaConfig := kafka.KafkaConfig{
			// 	Brokers: appConfig.Kafka.Brokers,
			// 	UseTLS:  appConfig.Kafka.UseTLS,
			// 	CAFile:  appConfig.Kafka.CAFile,
			// 	Topic:   appConfig.Kafka.Topic,
			// 	Authentcation: kafka.Authentcation{
			// 		Username:  appConfig.Kafka.Authentcation.Username,
			// 		Password:  appConfig.Kafka.Authentcation.Password,
			// 		Algorithm: appConfig.Kafka.Authentcation.Algorithm,
			// 	},
			// 	Version: appConfig.Kafka.Version,
			// }

			// _, err := kafka.InitKafka(&kafkaConfig)
			// if err != nil {
			// 	log.Println(err)
			// }

			// d := time.Duration(5 * time.Second)
			// for {
			// 	kafka.SendMessage(producer, appConfig.Kafka.Topic, time.Now().String())
			// 	time.Sleep(d)
			// }

			// Init DB
			connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", appConfig.Database.Username, appConfig.Database.Password, appConfig.Database.Host, appConfig.Database.Port, appConfig.Database.Database)
			db := database.NewMySQLDB(connectionString)

			userRepo := repositories.NewUserRepo(db)
			h := controllers.NewBaseHandler(userRepo)

			r := gin.Default()

			api := r.Group("/api")
			{
				api.POST("/token", h.GenerateToken)
				api.POST("/user/register", h.RegisterUser)
				secured := api.Group("/secured").Use(middleware.AuthenticationMiddleware())
				{
					secured.GET("/ping", controllers.Ping)
				}
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
