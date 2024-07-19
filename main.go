package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"github.com/venture-technology/vtx-responsible-service/config"
	"github.com/venture-technology/vtx-responsible-service/internal/service"
)

func main() {

	config, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := sql.Open("postgres", newPostgres(config.Database))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = migrate(db, config.Database.Schema)
	if err != nil {
		log.Fatalf("failed to execute migrations: %v", err)
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Cloud.Region),
		Credentials: credentials.NewStaticCredentials(config.Cloud.AccessKey, config.Cloud.SecretKey, config.Cloud.Token),
	})
	if err != nil {
		log.Fatalf("failed to create session at aws: %v", err)
	}

	router := gin.Default()
	producer := kafka.NewWriter(kafka.WriterConfig{Brokers: []string{config.Messaging.Brokers}, Topic: config.Messaging.Topic, Balancer: &kafka.LeastBytes{}})

	kafkaRepository := repository.NewKafkaRepository(producer)

	driverRepository := repository.NewDriverRepository(db)
	driverService := service.NewDriverService(driverRepository, kafkaRepository)
	driverController := controller.NewDriverController(driverService)

	driverController.RegisterRoutes(router)

	fmt.Println(driverController)
	router.Run(fmt.Sprintf(":%d", config.Server.Port))
}

func newPostgres(dbconfig config.Database) string {
	return "user=" + dbconfig.User +
		" password=" + dbconfig.Password +
		" dbname=" + dbconfig.Name +
		" host=" + dbconfig.Host +
		" port=" + dbconfig.Port +
		" sslmode=disable"
}

func migrate(db *sql.DB, filepath string) error {
	schema, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return err
	}

	return nil
}
