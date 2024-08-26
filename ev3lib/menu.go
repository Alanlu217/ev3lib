package ev3lib

type CommandPage struct {
	N string
	C []Command
}

type MenuConfig interface {
	GetCommandPages() []CommandPage
}

type Menu struct {
	config MenuConfig
}

func NewMenu(config MenuConfig) *Menu {
	return &Menu{}
}
