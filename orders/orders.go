package orders

type Order struct {
	ID int
	Customer string
	Status string
}

func SetStatusCopy(o Order, status string) {
	o.Status = status
}

func SetStatusPtr(o *Order, status string) {
	o.Status = status
}