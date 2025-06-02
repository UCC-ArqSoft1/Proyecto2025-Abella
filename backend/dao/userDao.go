package dao

import (
	_ "gorm.io/gorm"
)

type UserType struct {
	ID   uint
	Type string
}

type User struct {
	ID              uint `gorm:"primaryKey:type:uuid;autoIncrement"`
	UserTypeID      uint
	Email           string
	HashedPassword  string
	Name            string
	LastName        string
	Documentation   int // DNI
	IsCoach         bool
	CoachActivities []Activity `gorm:"foreignKey:CoachID"`
}
