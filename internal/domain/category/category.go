package category

import "errors"

type Category struct {
	Id          int8   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func New(name, description string) (*Category, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	return &Category{
		Id:          -1,
		Name:        name,
		Description: description,
	}, nil
}

func Build(id int8, name string, description string) (*Category, error) {
	return &Category{
		Id:          id,
		Name:        name,
		Description: description,
	}, nil
}
