package sub_category

import "fmt"

type SubCategory struct {
	ID       int8   `json:id`
	ParentID int8   `json:parent_category_id`
	Name     string `json:name`
}

func New(parentID int8, name string) (*SubCategory, error) {
	if parentID < 0 {
		return nil, fmt.Errorf("parent id must not be negative")
	}

	if name == "" {
		return nil, fmt.Errorf("name must not be empty")
	}

	return &SubCategory{
		ID:       -1,
		ParentID: parentID,
		Name:     name,
	}, nil
}

func Build(id, parentID int8, name string) *SubCategory {
	return &SubCategory{
		ID:       id,
		ParentID: parentID,
		Name:     name,
	}
}
