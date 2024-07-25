package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/customer"
	"github.com/stripe/stripe-go/v79/paymentmethod"
	"github.com/venture-technology/vtx-responsible-service/config"
	"github.com/venture-technology/vtx-responsible-service/internal/repository"
	"github.com/venture-technology/vtx-responsible-service/models"
	"github.com/venture-technology/vtx-responsible-service/utils"
)

type ResponsibleService struct {
	responsiblerepository repository.IResponsibleRepository
	kafkarepository       repository.IKafkaRepository
}

func NewResponsibleService(responsiblerepository repository.IResponsibleRepository, kafkarepository repository.IKafkaRepository) *ResponsibleService {
	return &ResponsibleService{
		responsiblerepository: responsiblerepository,
		kafkarepository:       kafkarepository,
	}
}

func (rs *ResponsibleService) CreateResponsible(ctx context.Context, responsible *models.Responsible) error {

	responsible.Password = utils.HashPassword(responsible.Password)

	return rs.responsiblerepository.CreateResponsible(ctx, responsible)
}

func (rs *ResponsibleService) GetResponsible(ctx context.Context, cpf *string) (*models.Responsible, error) {
	log.Printf("param read school -> cpf: %s", *cpf)
	return rs.responsiblerepository.GetResponsible(ctx, cpf)
}

func (rs *ResponsibleService) UpdateResponsible(ctx context.Context, currentResponsible, responsible *models.Responsible) error {
	log.Printf("input received to update school -> name: %s, cpf: %s, email: %s", responsible.Name, responsible.CPF, responsible.Email)
	return rs.responsiblerepository.UpdateResponsible(ctx, currentResponsible, responsible)
}

func (rs *ResponsibleService) DeleteResponsible(ctx context.Context, cpf *string) error {
	log.Printf("trying delete your infos --> %v", *cpf)
	return rs.responsiblerepository.DeleteResponsible(ctx, cpf)
}

func (rs *ResponsibleService) AuthResponsible(ctx context.Context, responsible *models.Responsible) (*models.Responsible, error) {
	responsible.Password = utils.HashPassword((responsible.Password))
	return rs.responsiblerepository.AuthResponsible(ctx, responsible)
}

func (rs *ResponsibleService) ParserJwtResponsible(ctx *gin.Context) (interface{}, error) {

	cpf, found := ctx.Get("cpf")

	if !found {
		return nil, fmt.Errorf("error while veryfing token")
	}

	return cpf, nil

}

func (rs *ResponsibleService) CreateTokenJWTResponsible(ctx context.Context, responsible *models.Responsible) (string, error) {

	conf := config.Get()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cpf": responsible.CPF,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	jwt, err := token.SignedString([]byte(conf.Server.Secret))

	if err != nil {
		return "", err
	}

	return jwt, nil

}

func (rs *ResponsibleService) CreateCustomer(ctx context.Context, responsible *models.Responsible) (*stripe.Customer, error) {

	conf := config.Get()

	log.Print(conf.StripeEnv.SecretKey)

	stripe.Key = conf.StripeEnv.SecretKey

	params := &stripe.CustomerParams{
		Name:  stripe.String(responsible.Name),
		Email: stripe.String(responsible.Email),
		Phone: stripe.String(responsible.Phone),
	}

	resp, err := customer.New(params)

	if err != nil {
		return nil, err
	}

	return resp, nil

}

func (rs *ResponsibleService) UpdateCustomer(ctx context.Context, customerId, email, phone string) (*stripe.Customer, error) {

	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	params := &stripe.CustomerParams{
		Email: &email,
		Phone: &phone,
	}

	updatedCustomer, err := customer.Update(customerId, params)

	if err != nil {
		return nil, err
	}

	return updatedCustomer, nil

}

func (rs *ResponsibleService) CreatePaymentMethod(ctx context.Context, customerId, cardToken *string) (*stripe.PaymentMethod, error) {

	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	params := &stripe.PaymentMethodParams{
		Type: stripe.String(string(stripe.PaymentMethodTypeCard)),
		Card: &stripe.PaymentMethodCardParams{
			Token: stripe.String(*cardToken),
		},
	}

	pm, err := paymentmethod.New(params)
	if err != nil {
		fmt.Println("Erro ao criar método de pagamento:", err)
		return nil, err
	}

	return pm, nil

}
