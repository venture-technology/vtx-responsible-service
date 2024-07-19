package service

import (
	"context"
	"log"

	"github.com/venture-technology/vtx-responsible-service/models"
)

type ResponsibleService struct {
	responsiblerepository repository.IResponsibleRepository
	kafkarepository       repository.IKafkaRepository
}

func NewResponsibleService(driverrepository repository.IResponsibleRepository, kafkarepository repository.IKafkaRepository) *ResponsibleService {
	return &ResponsibleService{
		responsiblerepository: responsiblerepository,
		kafkarepository:       kafkarepository,
	}
}

func (d *ResponsibleService) CreateResponsible(ctx context.Context, driver *models.Responsible) error {

	driver.Password = utils.HashPassword(driver.Password)

	// err := d.PublishKafkaMessage(ctx,
	// 	driver.Email,
	// 	fmt.Sprintf("Verification Email - %s", driver.Name),
	// 	fmt.Sprintf("Greetings %s, thank you very much for choosing us, we will be with you today, tomorrow and always. Venture, fast and safe.", driver.Name),
	// )

	// if err != nil {
	// 	return err
	// }

	return d.responsiblerepository.CreateResponsible(ctx, driver)
}

func (d *DriverService) GetDriver(ctx context.Context, cnh *string) (*models.Driver, error) {
	log.Printf("param read school -> cnh: %s", *cnh)
	return d.responsiblerepository.GetDriver(ctx, cnh)
}

func (d *DriverService) UpdateDriver(ctx context.Context, driver *models.Driver) error {
	log.Printf("input received to update school -> name: %s, cnh: %s, email: %s", driver.Name, driver.CNH, driver.Email)
	return d.responsiblerepository.UpdateDriver(ctx, driver)
}

func (d *DriverService) DeleteDriver(ctx context.Context, cnh *string) error {
	log.Printf("trying delete your infos --> %v", *cnh)
	return d.responsiblerepository.DeleteDriver(ctx, cnh)
}
