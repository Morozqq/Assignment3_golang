package main

import "fmt"

// Guitar представляет структуру для хранения информации о гитаре.
type Guitar struct {
	Brand  string
	Model  string
	Active bool
}

// Command интерфейс определяет метод Execute для выполнения команд.
type Command interface {
	Execute(guitar *Guitar)
}

// PlayCommand реализует интерфейс Command и представляет команду "Играть на гитаре".
type PlayCommand struct{}

func (c *PlayCommand) Execute(guitar *Guitar) {
	fmt.Printf("Играем на гитаре: %s %s\n", guitar.Brand, guitar.Model)
}

// ActivateCommand реализует интерфейс Command и представляет команду "Активировать гитару".
type ActivateCommand struct{}

func (c *ActivateCommand) Execute(guitar *Guitar) {
	guitar.Active = true
	fmt.Printf("Гитара активирована: %s %s\n", guitar.Brand, guitar.Model)
}

// DeactivateCommand реализует интерфейс Command и представляет команду "Деактивировать гитару".
type DeactivateCommand struct{}

func (c *DeactivateCommand) Execute(guitar *Guitar) {
	guitar.Active = false
	fmt.Printf("Гитара деактивирована: %s %s\n", guitar.Brand, guitar.Model)
}

// Invoker хранит команду и выполняет её.
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand(guitar *Guitar) {
	i.command.Execute(guitar)
}

func main() {
	// Создаем экземпляр гитары
	guitar := &Guitar{
		Brand:  "Fender",
		Model:  "Stratocaster",
		Active: false,
	}

	// Создаем вызывающий объект
	invoker := &Invoker{}

	// Назначаем команды и выполняем их
	playCommand := &PlayCommand{}
	activateCommand := &ActivateCommand{}
	deactivateCommand := &DeactivateCommand{}

	invoker.SetCommand(playCommand)
	invoker.ExecuteCommand(guitar)

	invoker.SetCommand(activateCommand)
	invoker.ExecuteCommand(guitar)

	invoker.SetCommand(deactivateCommand)
	invoker.ExecuteCommand(guitar)
}
