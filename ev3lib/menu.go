package ev3lib

import (
	"time"
)

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

////////////////////////////////////////////////////////////////////////////////
// Main Menu                                                                  //
////////////////////////////////////////////////////////////////////////////////

type MainMenu struct {
	i MainMenuInterface
	m *Menu

	commandIdx, pageIdx int
}

func NewMainMenu(i MainMenuInterface, m *Menu) *MainMenu {
	return &MainMenu{i, m, 0, 0}
}

func (m *MainMenu) Start() {
	t := time.NewTicker(time.Millisecond * 50)

main:
	for {
		// Check if program should exit
		if m.i.Exit() {
			break main
		}

		if m.i.NextCommand() {
			m.commandIdx += 1
		}

		if m.i.PreviousCommand() {
			m.commandIdx -= 1
		}

		f, idx := m.i.SetCommand()
		if f {
			m.commandIdx = idx
		}

		if m.i.NextPage() {
			m.pageIdx += 1
			m.commandIdx = 0
		}

		if m.i.PreviousPage() {
			m.pageIdx -= 1
			m.commandIdx = 0
		}

		f, idx = m.i.SetPage()
		if f {
			m.pageIdx = idx
			m.commandIdx = 0
		}

		m.pageIdx = Clamp(m.pageIdx, 0, len(m.m.Pages))
		m.commandIdx = Clamp(m.commandIdx, 0, len(m.m.Pages[m.pageIdx].Commands))

		if m.i.RunSelected() {
			RunTimedCommand(m.m.Pages[m.pageIdx].Commands[m.commandIdx], time.Millisecond*20)
		}

		m.i.Display(m.m, m.commandIdx, m.pageIdx)

		<-t.C
	}
}
