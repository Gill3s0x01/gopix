package model
import (
	"time"
	uuidn "github.com/satori/go.uuid"
)


type Bank struct {
	ID 				string 		`json:"id"`
	Code 			string		`json:"code"`
	Name 			string 		`json:"name"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}

func (bank *Bank) isValid() error{
	if bank.Code == "" {
		return errors.New("code is required")
	}
	if bank.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

func NewBank(code string, name string) (*Bank, error) {
	bank := Bank {
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()
	if err != nil {
		return nil, err
	}
	
	return &bank, nil
	
}