package seed

import (
	"github.com/SemmiDev/go-combo/api/models"
	"github.com/SemmiDev/go-combo/api/utils/random"
	"gorm.io/gorm"
	"log"
)

var village = &models.Village{
	Name:       "tinggam",
	PostalCode: "11111",
}

var user = &models.User{
	Username:  random.RandomUsername(),
	FullName:  random.RandomFullName(10),
	Email:     random.RandomEmail(),
	Password:  random.RandomPassword(),
	VillageID: 1,
}

func Load(db *gorm.DB) {

	err := db.Migrator().DropTable(&models.User{}, &models.Village{})
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = user.BeforeSaveUser()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(&models.User{}, &models.Village{})
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}

	err = db.Debug().Model(&models.Village{}).Create(village).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}

	err = db.Debug().Model(&models.User{}).Create(user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
}
