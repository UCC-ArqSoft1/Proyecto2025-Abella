package dao

type ActivityHour struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	ActivityID    uint
	Day           string
	Starting_Hour int
	Finish_hour   int
}

type ActivityType struct { // Categorias: 1 Categpria tiene muchas actividades
	ID   uint `gorm:"primaryKey;autoIncrement"`
	Name string
}

type Activity struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	Name           string
	Description    string
	ActivityTypeID uint
	Duration       uint
	Capacity       uint
	CoachID        uint
	Coach          User
	ActivityType   ActivityType
	ActivityHours  []ActivityHour `gorm:"foreignKey:ActivityID"`
}

type ActivityTypes []ActivityType

type Activities []Activity
