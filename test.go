package main

import "fmt"

func main() {
	newDeliverymen := NewDeliverymen(NewExpressDelivery("20/30", 30))
	newDeliverymen.DeliverPack()
	newDeliverymen.SetVariant(new(courierDelivery))
	newDeliverymen.DeliverPack()
	newDeliverymen.SetVariant(new(pickupDelivery))
	newDeliverymen.DeliverPack()

}

type Delivery interface {
	Deliver() error
}

type expressDelivery struct {
	time   string
	weight float32
}

func NewExpressDelivery(time string, weight int) Delivery {
	return &expressDelivery{"20/30", 30}
}

func (d *expressDelivery) Deliver() error {
	fmt.Println(d.time + " bystro otnesem vam")

	return nil
}

type courierDelivery struct {
	name string
	size float32
}

func NewCourierDelivery(name string, size int) Delivery {
	return &courierDelivery{
		name: "Jason",
		size: 30,
	}
}

func (d *courierDelivery) Deliver() error {
	fmt.Println(d.name + " Kurer otneset vam")

	return nil
}

type pickupDelivery struct {
	address  string
	password int
}

func NewPickupDelivery(address string, password int) Delivery {
	return &pickupDelivery{
		address:  "Masanchi 22/5",
		password: 223366,
	}
}

func (d *pickupDelivery) Deliver() error {
	fmt.Println("budew zabirat sam")

	return nil
}

type deliverymen struct {
	Delivery Delivery
}

func NewDeliverymen(dl Delivery) *deliverymen {
	return &deliverymen{
		Delivery: dl,
	}
}

func (d *deliverymen) DeliverPack() {
	d.Delivery.Deliver()

}

func (d *deliverymen) SetVariant(dl Delivery) {
	d.Delivery = dl
}
