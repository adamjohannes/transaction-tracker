package status

import "errors"

type Status struct {
	ID   int8   `json:"id"`
	Name string `json:"name"`
}

func New(name string) (*Status, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	return &Status{
		ID:   -1,
		Name: name,
	}, nil
}

func Build(id int8, name string) *Status {
	return &Status{
		ID:   id,
		Name: name,
	}
}
