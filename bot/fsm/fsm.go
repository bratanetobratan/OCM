package fsm

type Storage interface {
	SetState()
	State()
}
