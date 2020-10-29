package database

import (
	_const "github.com/bilalkocoglu/go-crud/pkg/const"
	"github.com/pkg/errors"
	"time"
)

type User struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Age       uint8      `json:"age"`
	Birthday  *time.Time `json:"birthday"`
	Address   Address    `gorm:"foreignKey:AddressId;references:id" json:"address"`
	AddressId uint
	Languages []*Language `gorm:"many2many:user_languages;"`
}

type Address struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	City     string `json:"city"`
	District string `json:"district"`
}

type Language struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `json:"name"`
}

func CreateDefaultUser() {
	var user User

	err := GetUserByUsername(&user, _const.Username)

	if err != nil {
		panic(err)
	}

	if user.ID == 0 {
		user = User{
			Username: _const.Username,
			Password: _const.Password,
			Name:     _const.Name,
			Email:    _const.Email,
			Age:      _const.Age,
			Address: Address{
				City:     _const.City,
				District: _const.District,
			},
			Languages: []*Language{
				{
					Name: "Turkce",
				},
				{
					Name: "Ingilizce",
				},
			},
		}

		err := SaveUser(&user)

		if err != nil {
			errors.Wrap(err, "Can not create default user")
		}
	}
}

func SaveUser(user *User) (err error) {
	if err = DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(user *User, username string) (err error) {
	DB.Preload("Languages").Joins("Address").Where("username = ?", username).First(user)

	return nil
}

func GetUserById(user *User, id uint64) (err error) {
	DB.Joins("Address").Where("id = ?", id).First(user)

	return nil
}
