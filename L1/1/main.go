package main

//Child struct
type Action struct {
	Human
}
//Parent struct
type Human struct {}
//Parent method
func (h Human) Dosmthn() {
	print("Do something")
}

func main() {
	//Define child
	a := Action{}
	//Call Parent method from child
	a.Dosmthn()
}