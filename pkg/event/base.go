package event

import "fmt"

type Base struct {
	ID   string
	Type Type
}

func (o *Base) GetID() string {
	return o.ID
}

func (o *Base) SetID(id string) {
	o.ID = id
}

func (o *Base) GetType() Type {
	return o.Type
}
func (o *Base) String() string {

	return fmt.Sprintf("id:%s type:%s", o.ID, o.Type)
}
