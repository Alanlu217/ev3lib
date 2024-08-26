package ev3lib

type MenuPages []CommandPage

func NewMenuPages() MenuPages {
	temp := make(MenuPages, 0)
	return temp
}

func (m MenuPages) AddPage(name string, pages ...Command) MenuPages {
	return append(m, CommandPage{name, pages})
}

type CommandPage struct {
	N string
	C []Command
}

type MenuConfig interface {
	GetCommandPages() MenuPages
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
