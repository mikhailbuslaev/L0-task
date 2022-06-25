package pattern

type Something struct {}

type Visitor interface {
	Visit(*Something)
}

type ConcreteVisitor1 struct {}
type ConcreteVisitor2 struct {}

func (v *ConcreteVisitor1) Visit(s *Something) {
	println("do something crazy...")
}

func (v *ConcreteVisitor2) Visit(s *Something) {
	println("do something ...")
}

func (s *Something) Accept(v *Visitor) {
	v.Visit(s)
}

/*
плюсы:
- можно расширять функционал класса, не добавляя к нему методов
- достаточно легкий для понимания паттерн
минусы:
- если visitor посещает несколько классов, но мы хотим поменять его работу, 
нам придется переписать чуть больше кода
*/