package pattern

type WorkersFacade struct {
	Worker1 *Worker1
	Worker2 *Worker2
	Worker3 *Worker3
}
// some random classes, we need to wrap their work into interface
type Worker1 struct {}
type Worker2 struct {}
type Worker3 struct {}

// some example methods for 1 struct
func (w *Worker1) DoSmthng() {}
func (w *Worker1) DoSmthngElse() {}

// example methods for 2 struct
func (w *Worker2) Set() {}
func (w *Worker2) RunSmthng() {}

// example methods for 3 struct
func (w *Worker3) ListenSmthng() {}

// example of facade methods
func (w *WorkersFacade) Set() {
	w.Worker2.Set()
}

func (w *WorkersFacade) Run() {
	w.Worker2.RunSmthng()
	w.Worker1.Dosmthng()
	w.Worker3.ListenSmthng()
	w.Worker1.DoSmthngElse
}

/*
плюсы: 
-один из самых понятных паттернов, поэтому легко использовать, легкость же чтения будет
 зависить от того, для чего мы будем писать фасад. Одно дело обернуть в фасад работу 
 3 воркеров, другое дело 10 воркеров, там уже будет сложнее
-фасад это интуитивно очевидный способ перейти на более высокий уровень абстракции
-способен инкапсулировать какие то области программы

минусы:
-фасад бывает неинтуитивен и непонятен в тех случаях, когда он перегружен и слишком абстрактен.
*/

