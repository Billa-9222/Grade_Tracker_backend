package entities

type Grades struct {
	ID int          `json:"id"`
	Student_ID int     `json:"student_id"`
	Subject string  `json:"subject"`
	Score int    `json:"score"`

}