package seed

import (
	"github.com/SemmiDev/go-combo/api/models"
	"github.com/SemmiDev/go-combo/api/utils/random"
	"github.com/jinzhu/gorm"
	"log"
)

var users = []models.User{
	models.User{
		Username: random.RandomUsername(),
		FullName: random.RandomFullName(10),
		Email:    random.RandomEmail(),
		Password: random.RandomPassword(),
	},
	models.User{
		Username: random.RandomUsername(),
		FullName: random.RandomFullName(10),
		Email:    random.RandomEmail(),
		Password: random.RandomPassword(),
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*
		err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
		if err != nil {
			log.Fatalf("attaching foreign key error: %v", err)
		}
	*/

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
