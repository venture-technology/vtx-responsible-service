package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

func (d *ResponsibleService) CreateResponsible(ctx context.Context, responsible *models.Responsible) error {

	responsible.Password = utils.HashPassword(responsible.Password)

	// err := d.PublishKafkaMessage(ctx,
	// 	driver.Email,
	// 	fmt.Sprintf("Verification Email - %s", driver.Name),
	// 	fmt.Sprintf("Greetings %s, thank you very much for choosing us, we will be with you today, tomorrow and always. Venture, fast and safe.", driver.Name),
	// )

	// if err != nil {
	// 	return err
	// }

	return d.responsiblerepository.CreateResponsible(ctx, responsible)
}

func (d *ResponsibleService) GetResponsible(ctx context.Context, cpf *string) (*models.Responsible, error) {
	log.Printf("param read school -> cpf: %s", *cpf)
	return d.responsiblerepository.GetResponsible(ctx, cpf)
}

func (d *ResponsibleService) UpdateResponsible(ctx context.Context, responsible *models.Responsible) error {
	log.Printf("input received to update school -> name: %s, cpf: %s, email: %s", responsible.Name, responsible.CPF, responsible.Email)
	return d.responsiblerepository.UpdateResponsible(ctx, responsible)
}

func (d *ResponsibleService) DeleteResponsible(ctx context.Context, cpf *string) error {
	log.Printf("trying delete your infos --> %v", *cpf)
	return d.responsiblerepository.DeleteResponsible(ctx, cpf)
}

func (d *ResponsibleService) AuthResponsible(ctx context.Context, responsible *models.Responsible) (*models.Responsible, error) {
	responsible.Password = utils.HashPassword((responsible.Password))
	return d.responsiblerepository.AuthResponsible(ctx, responsible)
}

func (d *ResponsibleService) ParserJwtResponsible(ctx *gin.Context) (interface{}, error) {

	cpf, found := ctx.Get("cpf")

	if !found {
		return nil, fmt.Errorf("error while veryfing token")
	}

	return cpf, nil

}

func (d *ResponsibleService) CreateTokenJWTResponsible(ctx context.Context, responsible *models.Responsible) (string, error) {

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
