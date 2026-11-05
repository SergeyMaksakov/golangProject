package personaldata

import "fmt"

// Personal содержит базовые персональные данные пользователя.
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print выводит информацию о пользователе на экран.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.2f\n", p.Weight)
	fmt.Printf("Рост: %.2f\n", p.Height)
}
