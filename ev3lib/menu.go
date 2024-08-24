package ev3lib

type Menu[T any] struct {
	Config T
	Runs   []func(T)
}

func NewMenu[T any](config T, runs ...func(T)) *Menu[T] {
	return &Menu[T]{Config: config, Runs: runs}
}
