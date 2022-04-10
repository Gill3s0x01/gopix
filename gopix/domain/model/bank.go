package model
import (
	uuidn "github.com/satori/go.uuid"
	Base "gopix/domain/model/base"
)


type Bank struct {
	Base `valid: "required"`
	Code 			string		`json:"code" valid:"notnull"`
	Name 			string 		`json:"name" valid:"notnull"`
}

func (bank *Bank) isValid() error{
	_, err := govalidator.ValidateStruct(bank)
	if err != nil {
		return err
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