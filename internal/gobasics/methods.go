package gobasics

type Gift struct {
	Name   string
	Author string
}

func (gift Gift) GetName() string {
	return gift.Name
}

func (gift *Gift) GetNameLink() string {
	//
	return gift.Name
}
