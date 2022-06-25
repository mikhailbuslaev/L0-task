package pattern

type Condition interface {
	Eat()
	Work()
	Sleep()
}

type Human struct {
	Name string
	Condition
}

type TiredCondition struct{}

type RestedCondition struct{}

func (t TiredCondition) Eat() {
	println("eats a lot")
}

func (t TiredCondition) Work() {
	println("works lazily")
}

func (t TiredCondition) Sleep() {
	println("sleep a lot")
}

func (r RestedCondition) Eat() {
	println("doesn't want to eat")
}

func (r RestedCondition) Work() {
	println("works a lot")
}

func (r RestedCondition) Sleep() {
	println("doesn't want to sleep")
}

func (h Human) SwitchState(c Condition) {
	h.Condition = c
}

/*
плюсы:
-позволяет менять поведение обьекта в зависимости от его состояния, 
так что от паттерна можно извлечь практическую пользу
*/