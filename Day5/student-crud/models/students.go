package models

type Student struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Grade int    `json:"grade"`
}