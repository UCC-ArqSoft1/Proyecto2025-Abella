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

type NewActivity struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Duration       int    `json:"duration"`
	CoachId        uint   `json:"coachid"`
	ActivityTypeID uint   `json:"activitytypeid"`
	Capacity       int    `json:"capacity"`
}

type NewHour struct {
	ActivityID    uint   `json:"id"`
	Day           string `json:"day"`
	Starting_Hour int    `json:"starting_hour"`
	Finish_hour   int    `json:"finish_hour"`
}

type UpdateActivity struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Duration       int    `json:"duration"`
	CoachId        uint   `json:"coachid"`
	ActivityTypeID uint   `json:"activitytypeid"`
	Capacity       int    `json:"capacity"`
}

type ActivityType struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type NewCategory struct {
	Name string `json:"name"`
}

type UpdateActivityHours struct {
	Id          uint   `json:"id"`
	ActivityId  uint   `json:"idactividad"`
	Day         string `json:"day"`
	Hour_start  int    `json:"starting_hour"`
	Hour_finish int    `json:"finish_hour"`
}

type Activities []Activity

type ActivityTypes []ActivityType
