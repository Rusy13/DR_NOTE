package fixtures

import (
	"time"

	"WB/internal/order/model"
	"WB/tests/states"
)

type OrderBuilder struct {
	instance *model.Order
}

func Order() *OrderBuilder {
	return &OrderBuilder{instance: &model.Order{}}
}

func (b *OrderBuilder) OrderUID(v string) *OrderBuilder {
	b.instance.OrderUID = v
	return b
}

func (b *OrderBuilder) Items(v []model.Item) *OrderBuilder {
	b.instance.Items = v
	return b
}

func (b *OrderBuilder) Shardkey(v string) *OrderBuilder {
	b.instance.Shardkey = v
	return b
}

func (b *OrderBuilder) Entry(v string) *OrderBuilder {
	b.instance.Entry = v
	return b
}

func (b *OrderBuilder) DateCreated(v time.Time) *OrderBuilder {
	b.instance.DateCreated = v
	return b
}

func (b *OrderBuilder) Delivery(v model.Delivery) *OrderBuilder {
	b.instance.Delivery = v
	return b
}

func (b *OrderBuilder) CustomerID(v string) *OrderBuilder {
	b.instance.CustomerID = v
	return b
}

func (b *OrderBuilder) DeliveryService(v string) *OrderBuilder {
	b.instance.DeliveryService = v
	return b
}

func (b *OrderBuilder) InternalSignature(v string) *OrderBuilder {
	b.instance.InternalSignature = v
	return b
}

func (b *OrderBuilder) Locale(v string) *OrderBuilder {
	b.instance.Locale = v
	return b
}

func (b *OrderBuilder) OofShard(v string) *OrderBuilder {
	b.instance.OofShard = v
	return b
}

func (b *OrderBuilder) Payment(v model.Payment) *OrderBuilder {
	b.instance.Payment = v
	return b
}

func (b *OrderBuilder) SmID(v int) *OrderBuilder {
	b.instance.SmID = v
	return b
}

func (b *OrderBuilder) TrackNumber(v string) *OrderBuilder {
	b.instance.TrackNumber = v
	return b
}

func (b *OrderBuilder) Ptr() *model.Order {
	return b.instance
}

func (b *OrderBuilder) Val() model.Order {
	return *b.instance
}

func (b *OrderBuilder) Valid1() *OrderBuilder {
	return Order().OrderUID(states.OrderUID1).SmID(states.Sm_id).Items(states.SLItems).Locale(states.Locale).OofShard(states.Oof_shard).TrackNumber(states.Track_number).DeliveryService(states.Delivery_service).Delivery(states.Delivery1).Payment(states.Payment1).InternalSignature(states.Internal_signature).CustomerID(states.Customer_id).Shardkey(states.Shardkey).DateCreated(states.Date_created).Entry(states.Entry)
}
