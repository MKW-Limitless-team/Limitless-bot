package modal

import "github.com/bwmarrin/discordgo"

type SelectMenu struct {
	discordgo.SelectMenu
}

func newSelectMenu() *SelectMenu {
	return &SelectMenu{discordgo.SelectMenu{}}
}

func NewStringSelectMenu(id string, minValues int, maxValues int) *SelectMenu {
	selectMenu := newSelectMenu().
		SetMenuType(discordgo.StringSelectMenu).
		SetCustomID(id).
		SetDisabled(false)

	return selectMenu
}

func (selectMenu *SelectMenu) SetMinSelect(minValues int) *SelectMenu {
	selectMenu.MinValues = &minValues

	return selectMenu
}

func (selectMenu *SelectMenu) SetDisabled(disabled bool) *SelectMenu {
	selectMenu.Disabled = disabled

	return selectMenu
}

func (selectMenu *SelectMenu) SetMenuType(menuType discordgo.SelectMenuType) *SelectMenu {
	selectMenu.MenuType = menuType

	return selectMenu
}

func (selectMenu *SelectMenu) SetCustomID(customID string) *SelectMenu {
	selectMenu.CustomID = customID

	return selectMenu
}

func (selectMenu *SelectMenu) AddMenuOption(menuOption SelectMenuOption) *SelectMenu {
	selectMenu.Options = append(selectMenu.Options, menuOption.SelectMenuOption)

	return selectMenu
}
