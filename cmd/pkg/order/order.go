package order

import "github.com/vmihailenco/msgpack/v4"

type Order map[string]interface{}

func (o Order) SetData(data *Data) Order {
	o["order_data"] = data
	return o
}
func (o Order) Data() (*Data, error) {
	d := &Data{}
	err := d.UnmarshalBinary([]byte(o["order_data"].(string)))
	return d, err
}

// Data --
type Data struct {
	CustomerID     string
	ItemIDs        []string
	ItemQuantities []int
}

func (d Data) MarshalBinary() ([]byte, error) {
	return msgpack.Marshal(d)
}

func (d *Data) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, d)
}
