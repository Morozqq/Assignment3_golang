package main

import (
	"fmt"
	"sync"
)

// Guitar представляет собой структуру для хранения информации о гитаре.
type Guitar struct {
	Brand  string
	Model  string
	Year   int
	Active bool
}

// guitarSingleton представляет собой структуру Singleton для хранения единственного экземпляра гитары.
type guitarSingleton struct {
	guitar *Guitar
	sync.Mutex
}

var instance *guitarSingleton

// GetGuitarInstance возвращает экземпляр Singleton гитары.
func GetGuitarInstance() *Guitar {
	if instance == nil {
		instance = &guitarSingleton{
			guitar: &Guitar{
				Brand:  "Fender",
				Model:  "Stratocaster",
				Year:   2020,
				Active: true,
			},
		}
	}
	return instance.guitar
}

func main() {
	// Получаем экземпляр гитары с использованием Singleton
	guitar1 := GetGuitarInstance()
	fmt.Printf("Гитара 1: %v %v, выпущена в %v, активна: %v\n", guitar1.Brand, guitar1.Model, guitar1.Year, guitar1.Active)

	// Попробуем получить другой экземпляр гитары
	guitar2 := GetGuitarInstance()
	fmt.Printf("Гитара 2: %v %v, выпущена в %v, активна: %v\n", guitar2.Brand, guitar2.Model, guitar2.Year, guitar2.Active)

	// Оба экземпляра гитары ссылаются на один и тот же Singleton, поэтому они идентичны.
	fmt.Println("guitar1 и guitar2 являются одним и тем же объектом:", guitar1 == guitar2)
}
