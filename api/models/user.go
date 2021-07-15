package models

import (
	"context"
	ErrMsg "github.com/SemmiDev/go-combo/api/errors_messages"
	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"log"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar;size:255;not null;unique" json:"username"`
	FullName  string    `gorm:"type:varchar;size:255;not null" json:"full_name"`
	Email     string    `gorm:"type:varchar;size:255;not null;unique" json:"email"`
	Password  string    `gorm:"type:varchar;size:255;not null" json:"password"`
	VillageID uint64    `gorm:"not null" json:"village_id"`
	Village   *Village  `json:"village"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSaveUser() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.FullName = html.EscapeString(strings.TrimSpace(u.FullName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Username == "" {
			return ErrMsg.ErrUsernameRequired
		}
		if u.FullName == "" {
			return ErrMsg.ErrFullNameRequired
		}
		if u.Password == "" {
			return ErrMsg.ErrPasswordRequired
		}
		if u.Email == "" {
			return ErrMsg.ErrInvalidEmail
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return ErrMsg.ErrInvalidEmail
		}
		return nil

	case "login":
		if u.Password == "" {
			return ErrMsg.ErrPasswordRequired
		}
		if u.Email == "" {
			return ErrMsg.ErrEmailRequired
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return ErrMsg.ErrInvalidEmail
		}
		return nil

	default:
		if u.Username == "" {
			return ErrMsg.ErrUsernameRequired
		}
		if u.Password == "" {
			return ErrMsg.ErrPasswordRequired
		}
		if u.Email == "" {
			return ErrMsg.ErrEmailRequired
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return ErrMsg.ErrInvalidEmail
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	ctx := context.Background()
	tx := db.WithContext(ctx).
		Where("users.id = ?", u.ID).
		Joins("Village").
		First(u)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	var users []User

	err = db.Debug().Model(&User{}).Limit(100).Joins("Village").Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}

	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint64) (*User, error) {
	var err error

	ctx := context.Background()
	var user User
	tx := db.WithContext(ctx).
		Where("users.id = ?", uid).
		Joins("Village").
		First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if gorm.ErrRecordNotFound == err {
		return &User{}, ErrMsg.ErrUserNotFound
	}

	return &user, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint64) (*User, error) {
	// To hash the password
	err := u.BeforeSaveUser()
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":   u.Password,
			"username":   u.Username,
			"full_name":  u.FullName,
			"email":      u.Email,
			"village_id": u.VillageID,
			"updated_at": time.Now(),
		},
	)

	if db.Error != nil {
		return &User{}, db.Error
	}

	// This is the display the updated user
	ctx := context.Background()

	tx := db.WithContext(ctx).
		Select("*").
		Joins("LEFT JOIN villages on villages.id = users.village_id").
		Where("users.id = ?", uid).
		First(&u)

	if tx.Error != nil {
		log.Println("------------------------")
		log.Println(tx.Error)
		return nil, tx.Error
	}

	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint64) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
