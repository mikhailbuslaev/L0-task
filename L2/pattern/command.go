package pattern

type Command interface {
	Execute()
}

type SwitchCommand struct{
	receiver *Receiver
}

type OnCommand struct{}
type OffCommand struct{}

type Sender struct{}
type Receiver struct{}

func (s *Sender) Send(c Command) {
	println("sending command...")
	c.Execute
}

func (r *Receiver) SwitchState() {
	println("switching state of receiver...")
}

func (r *Receiver) On() {
	println("start receiver...")
}

func (r *Receiver) Off() {
	println("breaking receiver...")
}

func (c *OnCommand) Execute() {
	c.receiver.On()
}

func (c *OffCommand) Execute() {
	c.receiver.Off()
}

func (c *SwitchCommand) Execute() {
	c.receiver.Switch()
}

/*
плюсы:
-отделяет отправителя комманды от получателя
*/