package fixtures

import (
	"WB/internal/order/model"
	"WB/tests/states"
)

type ItemBuilder struct {
	instance *model.Item
}

func Item() *ItemBuilder {
	return &ItemBuilder{instance: &model.Item{}}
}

func (b *ItemBuilder) OrderUID(v string) *ItemBuilder {
	b.instance.OrderUID = v
	return b
}

func (b *ItemBuilder) ChrtID(v int) *ItemBuilder {
	b.instance.ChrtID = v
	return b
}

func (b *ItemBuilder) TrackNumber(v string) *ItemBuilder {
	b.instance.TrackNumber = v
	return b
}

func (b *ItemBuilder) Price(v int) *ItemBuilder {
	b.instance.Price = v
	return b
}

func (b *ItemBuilder) RID(v string) *ItemBuilder {
	b.instance.Rid = v
	return b
}

func (b *ItemBuilder) Name(v string) *ItemBuilder {
	b.instance.Name = v
	return b
}

func (b *ItemBuilder) Sale(v int) *ItemBuilder {
	b.instance.Sale = v
	return b
}

func (b *ItemBuilder) Size(v string) *ItemBuilder {
	b.instance.Size = v
	return b
}

func (b *ItemBuilder) TotalPrice(v int) *ItemBuilder {
	b.instance.TotalPrice = v
	return b
}

func (b *ItemBuilder) NMID(v int) *ItemBuilder {
	b.instance.NmID = v
	return b
}

func (b *ItemBuilder) Brand(v string) *ItemBuilder {
	b.instance.Brand = v
	return b
}

func (b *ItemBuilder) Status(v int) *ItemBuilder {
	b.instance.Status = v
	return b
}

func (b *ItemBuilder) Valid1() *ItemBuilder {
	return Item().OrderUID(states.OrderUID1).ChrtID(states.Item1.ChrtID).TrackNumber(states.Item1.TrackNumber).Price(states.Item1.Price).RID(states.Item1.Rid).Name(states.Item1.Name).Sale(states.Item1.Sale).Size(states.Item1.Size).TotalPrice(states.Item1.TotalPrice).NMID(states.Item1.NmID).Brand(states.Item1.Brand).Status(states.Item1.Status)
}

func (b *ItemBuilder) Ptr() *model.Item {
	return b.instance
}

func (b *ItemBuilder) Val() model.Item {
	return *b.instance
}
