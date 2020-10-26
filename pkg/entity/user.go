package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	Address      Address `gorm:"embedded;embeddedPrefix:address_"`
}

type Address struct {
	City     string
	District string
}
