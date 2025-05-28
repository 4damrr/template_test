package template

import "time"

type BaseResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UserData struct {
	Name        string       `json:"name"`
	Email       string       `json:"email"`
	Phone       string       `json:"phone"`
	UserName    string       `json:"username"`
	Summary     string       `json:"summary"`
	Skills      []string     `json:"skills"`
	Experiences []Experience `json:"experiences"`
	Educations  []Education  `json:"educations"`
}

type Experience struct {
	Name         string    `json:"name"`
	Descriptions []string  `json:"descriptions"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
}

type Education struct {
	Major  string `json:"major"`
	School string `json:"school"`
	Year   int    `json:"year"`
}

type ExprEnv struct {
	Price    float64   `expr:"price"`
	Discount float64   `expr:"discount"`
	Date     time.Time `expr:"date"`
}
