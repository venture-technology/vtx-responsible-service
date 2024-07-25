package service

import (
	"context"
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/segmentio/kafka-go"
	"github.com/venture-technology/vtx-responsible-service/config"
	"github.com/venture-technology/vtx-responsible-service/internal/repository"
	"github.com/venture-technology/vtx-responsible-service/models"
)

func newPostgres(dbConfig config.Database) string {
	return "user=" + dbConfig.User +
		" password=" + dbConfig.Password +
		" dbname=" + dbConfig.Name +
		" host=" + dbConfig.Host +
		" port=" + dbConfig.Port +
		" sslmode=disable"
}

func mockResponsible() *models.Responsible {
	return &models.Responsible{
		Name:          "John Doe",
		Email:         "johndoe@example.com",
		Password:      "123teste",
		CPF:           "79887042064",
		Street:        "Rua Itamerendiba",
		Number:        "26",
		ZIP:           "12345678",
		Complement:    "Apto 32",
		Status:        "ACTIVE",
		CardToken:     "tok_visa",
		PaymentMethod: "",
		CustomerId:    "",
	}
}

func setupTesteDb(t *testing.T) (*sql.DB, *ResponsibleService) {

	t.Helper()

	config, err := config.Load("../../config/config.yaml")
	if err != nil {
		t.Fatalf("falha ao carregar a configuração: %v", err)
	}

	db, err := sql.Open("postgres", newPostgres(config.Database))
	if err != nil {
		t.Fatalf("falha ao conectar ao banco de dados: %v", err)
	}

	if err != nil {
		log.Fatalf("failed to create session at aws: %v", err)
	}

	producer := kafka.NewWriter(kafka.WriterConfig{Brokers: []string{config.Messaging.Brokers}, Topic: config.Messaging.Topic, Balancer: &kafka.LeastBytes{}})

	responsibleRepository := repository.NewResponsibleRepository(db)
	kafkaRepository := repository.NewKafkaRepository(producer)

	responsibleService := NewResponsibleService(responsibleRepository, kafkaRepository)

	return db, responsibleService

}

func TestCreateResponsible(t *testing.T) {

}

func TestGetResponsible(t *testing.T) {

}

func TestUpdateResponsible(t *testing.T) {

}

func TestDeleteResponsible(t *testing.T) {

}

func TestCreateCustomer(t *testing.T) {

	db, responsibleService := setupTesteDb(t)
	defer db.Close()

	responsibleMock := mockResponsible()

	_, err := responsibleService.CreateCustomer(context.Background(), responsibleMock)

	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestUpdateCustomer(t *testing.T) {

	db, responsibleService := setupTesteDb(t)
	defer db.Close()

	responsibleMock := mockResponsible()

	responsibleMock.CustomerId = "cus_QXDV6WVFkYAe4a"
	responsibleMock.Email = "johndoetesteupdate@example.com"
	responsibleMock.Phone = "+55 11 91234 5678"

	_, err := responsibleService.UpdateCustomer(context.Background(), responsibleMock.CustomerId, responsibleMock.Email, responsibleMock.Phone)

	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestDeleteCustomer(t *testing.T) {

}
