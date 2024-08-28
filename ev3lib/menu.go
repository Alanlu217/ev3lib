package ev3lib

type NamedCommand struct {
	Name string
	CommandInterface
}

type CommandPage struct {
	menu *CommandMenu

	Name     string
	Commands []NamedCommand
}

func (c *CommandPage) AddCommand(name string, command CommandInterface) *CommandPage {
	c.Commands = append(c.Commands, NamedCommand{name, command})

	return c
}

func (c *CommandPage) Add() {
	c.menu.Pages = append(c.menu.Pages, c)
}

type CommandMenu struct {
	Pages []*CommandPage
}

func NewCommandMenu() *CommandMenu {
	return &CommandMenu{Pages: make([]*CommandPage, 0)}
}

func (c *CommandMenu) AddPage(name string) *CommandPage {
	return &CommandPage{menu: c, Name: name, Commands: make([]NamedCommand, 0)}
}

type MenuConfig interface {
	GetCommandPages() CommandMenu
}

////////////////////////////////////////////////////////////////////////////////
// Menu                                                                       //
////////////////////////////////////////////////////////////////////////////////

type Menu struct {
	config MenuConfig
}

// NewMenu takes a MenuConfig and returns a new menu
func NewMenu(config MenuConfig) *Menu {
	return &Menu{config}
}
