package common

type Course struct {
    ID        int `json:"id"`
    TeacherID int `json:"teacher_id"`
    Name      string `json:"name"`
}