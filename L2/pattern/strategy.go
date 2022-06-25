package pattern

type Executor struct {
	Strategy
}

type Strategy interface {
	Execute()
}

type Strategy1 struct{}
type Strategy2 struct{}

func (s *Strategy1) Execute() {
	println("executing strategy №1...")
}

func (s *Strategy2) Execute() {
	println("executing strategy №2...")
}

/*
плюсы:
-отделяет выбор алгоритма от его рализации, обособление реализации стратегии
-стратегия может быть какой угодно, мы просто подключим ее путем Executor{Strategy1} и запустим Executor.Execute()
*/

