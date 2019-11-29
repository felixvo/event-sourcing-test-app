package event

type Base struct {
	ID string
}

func (o *Base) GetID() string {
	return o.ID
}
func (o *Base) SetID(id string) {
	o.ID = id
}
