package domain

type ActivityHours struct {
	Id          uint   `json:"id"`
	Day         string `json:"day"`
	Hour_start  int    `json:"hour_start"`
	Hour_finish int    `json:"hour_finish"`
}

type Activity struct {
	ID            uint            `json:"id"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	Duration      uint            `json:"duration"`
	CoachName     string          `json:"coach_name"`
	ActivityType  string          `json:"activitytype"`
	ActivityHours []ActivityHours `json:"activity_hours"`
}

type Activities []Activity

type MakeInscription struct {
	UserId      uint   `json:"userid"`
	ActivityId  uint   `json:"activityid"`
	Day         string `json:"day"`
	Hour_start  int    `json:"hour_start"`
	Hour_finish int    `json:"hour_finish"`
}

type Inscription struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Duration    uint   `json:"duration"`
	CoachName   string `json:"coach_name"`
	Day         string `json:"day"`
	Hour_start  int    `json:"hour_start"`
	Hour_finish int    `json:"hour_finish"`
}

type Inscriptions []Inscription
