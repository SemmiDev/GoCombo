package models

import (
	ErrMsg "github.com/SemmiDev/go-combo/api/errors_messages"
	"gorm.io/gorm"
	"html"
	"strings"
	"time"
)

type Village struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement; not null" json:"id"`
	Name       string    `gorm:"type:varchar;size:255;not null" json:"name"`
	PostalCode string    `gorm:"type:varchar;size:5;not null" json:"postal_code"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (v *Village) PrepareVillage() {
	v.ID = 0
	v.Name = html.EscapeString(strings.TrimSpace(v.Name))
	v.CreatedAt = time.Now()
	v.UpdatedAt = time.Now()
}

func (v *Village) ValidateVillage() error {
	if v.Name == "" {
		return ErrMsg.ErrVillageNameRequired
	}
	if v.PostalCode == "" {
		return ErrMsg.ErrPostalCodeRequired
	}
	return nil
}

func (v *Village) SaveVillage(db *gorm.DB) (*Village, error) {
	var err error
	err = db.Debug().Create(&v).Error
	if err != nil {
		return &Village{}, err
	}
	return v, nil
}

func (v *Village) FindAllVillages(db *gorm.DB) (*[]Village, error) {
	var err error
	var users []Village
	err = db.Debug().Model(&Village{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]Village{}, err
	}
	return &users, err
}

func (v *Village) FindVillageByID(db *gorm.DB, uid uint64) (*Village, error) {
	var err error
	err = db.Debug().Model(Village{}).Where("id = ?", uid).Take(&v).Error
	if err != nil {
		return &Village{}, err
	}
	if gorm.ErrRecordNotFound == err {
		return &Village{}, ErrMsg.ErrVillageNotFound
	}
	return v, err
}

func (v *Village) UpdateAVillage(db *gorm.DB, uid uint64) (*Village, error) {
	db = db.Debug().Model(&Village{}).Where("id = ?", uid).Take(&Village{}).UpdateColumns(
		map[string]interface{}{
			"name":        v.Name,
			"postal_code": v.PostalCode,
			"updated_at":  time.Now(),
		},
	)
	if db.Error != nil {
		return &Village{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&Village{}).Where("id = ?", uid).Take(&v).Error
	if err != nil {
		return &Village{}, err
	}
	return v, nil
}

func (v *Village) DeleteAVillage(db *gorm.DB, uid uint64) (int64, error) {
	db = db.Debug().Model(&Village{}).Where("id = ?", uid).Take(&Village{}).Delete(&Village{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
