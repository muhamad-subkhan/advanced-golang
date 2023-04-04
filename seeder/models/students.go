package models

type Student struct {
	Student_id      uint64 `json:"student_id" binding:"Required"`
	Student_name    string `json:"student_name" binding:"Required"`
	Student_age     uint64 `json:"student_age" binding:"Required"`
	Student_address string `json:"student_address" binding:"Required"`
	Student_phone   string `json:"student_phone" binding:"Required"`
}