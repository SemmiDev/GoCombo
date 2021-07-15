package models

//
//import (
//	ErrMsg "github.com/SemmiDev/go-combo/api/errors_messages"
//	"github.com/jinzhu/gorm"
//	"html"
//	"strings"
//	"time"
//)
//
//type Province struct {
//	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
//	Name      string    `gorm:"size:255;not null;unique" json:"name"`
//	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
//	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
//}
//
//func (p *Province) Prepare() {
//	p.ID = 0
//	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
//	p.CreatedAt = time.Now()
//	p.UpdatedAt = time.Now()
//}
//
//func (p *Province) Validate() error {
//	if p.Name == "" {
//		return ErrMsg.ErrNameRequired
//	}
//	return nil
//}
//
//func (p *Province) SaveProvince(db *gorm.DB) (*Province, error) {
//	var err error
//	err = db.Debug().Create(&p).Error
//	if err != nil {
//		return &Province{}, err
//	}
//	return p, nil
//}
//
//func (p *Province) FindAllProvinces(db *gorm.DB) (*[]Province, error) {
//	var err error
//	var province []Province
//	err = db.Debug().Model(&Province{}).Limit(36).Find(&province).Error
//	if err != nil {
//		return &[]Province{}, err
//	}
//	return &province, err
//}
//
//func (p *Province) FindProvinceByID(db *gorm.DB, uid uint32) (*Province, error) {
//	var err error
//	err = db.Debug().Model(Province{}).Where("id = ?", uid).Take(&p).Error
//	if err != nil {
//		return &Province{}, err
//	}
//	if gorm.IsRecordNotFoundError(err) {
//		return &Province{}, ErrMsg.ErrProvinceNotFound
//	}
//	return p, err
//}
//
//func (p *Province) UpdateAProvince(db *gorm.DB, uid uint32) (*Province, error) {
//	db = db.Debug().Model(&Province{}).Where("id = ?", uid).Take(&Province{}).UpdateColumns(
//		map[string]interface{}{
//			"name":       p.Name,
//			"updated_at": time.Now(),
//		},
//	)
//	if db.Error != nil {
//		return &Province{}, db.Error
//	}
//	// This is the display the updated user
//	err := db.Debug().Model(&Province{}).Where("id = ?", uid).Take(&p).Error
//	if err != nil {
//		return &Province{}, err
//	}
//	return p, nil
//}
//
//func (p *Province) DeleteAProvince(db *gorm.DB, uid uint64) (int64, error) {
//	db = db.Debug().Model(&Province{}).Where("id = ?", uid).Take(&Province{}).Delete(&Province{})
//	if db.Error != nil {
//		return 0, db.Error
//	}
//	return db.RowsAffected, nil
//}
