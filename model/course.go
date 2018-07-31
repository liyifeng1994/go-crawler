package model

type JobCourse struct {
	Job            string
	Title          string
	Time           string
	StudentsNumber int
	Score          float32
	Price          float32
}

type CodingCourse struct {
	Title          string
	OriginalTitle  string
	Teacher        string
	Level          string
	Time           string
	StudentsNumber int
	Score          float32
	Price          float32
	Deleted        bool
}
