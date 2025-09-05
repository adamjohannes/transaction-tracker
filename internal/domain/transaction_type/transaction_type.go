package transaction_type

import "errors"

type TransactionType struct {
	ID   int8   `json:"id"`
	Name string `json:"name"`
}

func New(name string) (*TransactionType, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	return &TransactionType{
		ID:   -1,
		Name: name,
	}, nil
}

func Build(id int8, name string) *TransactionType {
	return &TransactionType{
		ID:   id,
		Name: name,
	}
}
