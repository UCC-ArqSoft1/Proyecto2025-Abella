package dao

type Inscription struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	UserID        uint
	ActivityID    uint
	Day           string
	Starting_Hour int
	Finish_hour   int
	Activity      Activity
	User          User
}

type Inscriptions []Inscription
