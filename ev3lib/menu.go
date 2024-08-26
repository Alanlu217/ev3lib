package ev3lib

type CommandPage struct {
	Name     string
	Commands []Command
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
