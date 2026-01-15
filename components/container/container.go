package container

import "github.com/bwmarrin/discordgo"

type Container struct {
	discordgo.Container
}

func newContainer() Container {
	return Container{discordgo.Container{}}
}

func NewBasicContainer() Container {
	return newContainer()
}

func (container Container) SetColor(color int) Container {
	container.AccentColor = &color

	return container
}

func (container Container) AddComponent(component discordgo.MessageComponent) Container {
	container.Components = append(container.Components, component)

	return container
}
