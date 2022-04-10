package model
import "time"

type Bank struct {
	ID 				string
	Code 			string
	Name 			string
	CreatedAt time.Time
	UpdatedAt time.Time
}