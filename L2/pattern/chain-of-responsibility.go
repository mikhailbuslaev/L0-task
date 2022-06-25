package pattern

type Message struct {
	Text []byte
}

type Handler interface {
	Handle(m *Message)
	Successor() Handler
}

type ConcreteHandler1 struct {}

type ConcreteHandler2 struct {}

func (h *ConcreteHandler1) Successor() Handler {
	return &ConcreteHandler2{}// classic delegate
}

func (h *ConcreteHandler2) Successor() Handler {
	return &ConcreteHandler2{}// if we want stop chain, return same handler
}

func (h *ConcreteHandler1) Handle(m *Message) {
	if len(m.Text) > 100 {
		println("somehow handle message")
	} else if h.Successor() != h {
		println("delegate message to next handler...")
		h.Successor().Handle(m)
	}
}

func (h *ConcreteHandler2) Handle(m *Message) {
	if len(m.Text) > 0 {
		println("somehow else handle message")
	} else if h.Successor() != h {
		println("delegate message to next handler...")
		h.Successor().Handle(m)
	}
}

func Main_example() {
	m := &Message{Text:[]byte("some example message")}
	h := &ConcreteHandler1{}
	h.Handle(m)
}

/*
плюсы:
-можно подключать большое количество ручек
-упорядочивает обработку сообщения

минусы:
- если писать ручку как интерфейс, появится небольшая помеха с 
механизмом получения наследника, поскольку у интерфесов нет полей данных, 
но ручку можно оформить как структуру с полем наследника и полем функции обработчика(либо же интерфейса),
и добавить метод получения функции обработчика, но так будет больше кода, поэтому не стал писать
- сообщение может не обработаться, но можно написать default-handler как в switch
*/


