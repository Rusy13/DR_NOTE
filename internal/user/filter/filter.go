package filter

type Filter struct {
	OrderID uint64
}

func New(orderID uint64) Filter {
	return Filter{
		OrderID: orderID,
	}
}
