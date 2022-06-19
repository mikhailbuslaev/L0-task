package main

type CafeVisitor interface {
	BuyFood()
	EatFood()
}

type NormalVisitor struct {}

func (n *NormalVisitor) BuyFood() {
	fmt.Println("Buy something from menu")
}

func (n *NormalVisitor) EatFood() {
	fmt.Println("Eat their order")
}

// Some visitors just need toilet, not food or drinks
type ManWhoWantsToilet struct {
	Piss()
}

func (m *ManWhoWantsToilet) Piss() {
	fmt.Println("Pissing...")
}

type ToiletToCafeAdapter struct {
	ManWhoWantsToilet
	BuyFood()
	EatFood()
}

func (t *ToiletToCafeAdapter) BuyFood() {
	fmt.Println("Buy cheap drink")
}

func (t *ToiletToCafeAdapter) EatFood() {
	fmt.Println("Not drink it")
}
