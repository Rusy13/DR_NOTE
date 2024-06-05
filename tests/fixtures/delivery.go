package fixtures

import (
	"WB/internal/order/model"
	"WB/tests/states"
)

type DeliveryBuilder struct {
	instance *model.Delivery
}

func Delivery() *DeliveryBuilder {
	return &DeliveryBuilder{instance: &model.Delivery{}}
}

func (b *DeliveryBuilder) OrderUID(v string) *DeliveryBuilder {
	b.instance.OrderUID = v
	return b
}

func (b *DeliveryBuilder) Name(v string) *DeliveryBuilder {
	b.instance.Name = v
	return b
}

func (b *DeliveryBuilder) Phone(v string) *DeliveryBuilder {
	b.instance.Phone = v
	return b
}

func (b *DeliveryBuilder) Zip(v string) *DeliveryBuilder {
	b.instance.Zip = v
	return b
}

func (b *DeliveryBuilder) City(v string) *DeliveryBuilder {
	b.instance.City = v
	return b
}

func (b *DeliveryBuilder) Address(v string) *DeliveryBuilder {
	b.instance.Address = v
	return b
}

func (b *DeliveryBuilder) Region(v string) *DeliveryBuilder {
	b.instance.Region = v
	return b
}

func (b *DeliveryBuilder) Email(v string) *DeliveryBuilder {
	b.instance.Email = v
	return b
}

func (b *DeliveryBuilder) Valid1() *DeliveryBuilder {
	return Delivery().
		OrderUID(states.OrderUID1).
		Name(states.Delivery1.Name).
		Phone(states.Delivery1.Phone).
		Zip(states.Delivery1.Zip).
		City(states.Delivery1.City).
		Address(states.Delivery1.Address).
		Region(states.Delivery1.Region).
		Email(states.Delivery1.Email)
}

func (b *DeliveryBuilder) Ptr() *model.Delivery {
	return b.instance
}

func (b *DeliveryBuilder) Val() model.Delivery {
	return *b.instance
}
