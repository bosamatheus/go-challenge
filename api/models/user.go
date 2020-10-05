package models

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:200;not null" json:"name"`
	Cellphone string    `gorm:"size:20;not null" json:"cellphone"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func FormatCellphone(str string) string {
	return "+" + str[0:2] + " (" + str[2:4] + ") " + str[4:9] + "-" + str[9:len(str)]
}

func (u *User) Prepare(client string) {
	u.ID = 0
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	if client == os.Getenv("API_SECRET_MACAPA") {
		u.Name = strings.ToUpper(strings.TrimSpace(u.Name))
		u.Cellphone = FormatCellphone(strings.TrimSpace(u.Cellphone))
	}
	if client == os.Getenv("API_SECRET_VAREJAO") {
		u.Name = strings.TrimSpace(u.Name)
		u.Cellphone = strings.TrimSpace(u.Cellphone)
	}
}

func (u *User) Validate(client string) error {
	if client == os.Getenv("API_SECRET_MACAPA") && len(u.Cellphone) != 19 {
		return errors.New("Invalid length of cellphone")
	}
	if client == os.Getenv("API_SECRET_VAREJAO") && len(u.Cellphone) != 13 {
		return errors.New("Invalid length of cellphone")
	}
	return nil
}

func (u *User) CreateUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("ID = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User not found.")
	}
	return u, err
}

func (u *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {
	db = db.Debug().Model(&User{}).Where("ID = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"name":      u.Name,
			"cellphone": u.Cellphone,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}

	var err error
	err = db.Debug().Model(&User{}).Where("ID = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&User{}).Where("ID = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
