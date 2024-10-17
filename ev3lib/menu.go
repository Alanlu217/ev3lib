package ev3lib

type NamedCommand struct {
	Name string
	CommandInterface
}

////////////////////////////////////////////////////////////////////////////////
// MenuPage                                                                       //
////////////////////////////////////////////////////////////////////////////////

type MenuPage struct {
	menu *Menu

	Name     string
	Commands []NamedCommand
}

func (c *MenuPage) AddCommand(name string, command CommandInterface) *MenuPage {
	c.Commands = append(c.Commands, NamedCommand{name, command})

	return c
}

func (c *MenuPage) Add() {
	c.menu.Pages = append(c.menu.Pages, c)
}

////////////////////////////////////////////////////////////////////////////////
// Menu                                                                       //
////////////////////////////////////////////////////////////////////////////////

type Menu struct {
	Pages []*MenuPage
}

func NewCommandMenu() *Menu {
	return &Menu{Pages: make([]*MenuPage, 0)}
}

func (c *Menu) AddPage(name string) *MenuPage {
	return &MenuPage{menu: c, Name: name, Commands: make([]NamedCommand, 0)}
}

type MenuConfig interface {
	GetCommandPages() Menu
}
