package main

import "fmt"

/* Builder Patterns */
/* แยกการสร้างอ็อบเจ็กต์ที่ซับซ้อนออกจากการแทนที่ เพื่อให้สามารถสร้าง */
/* การแทนที่ที่แตกต่างกันได้ */
type (
	Meal struct {
		MainCourse string
		Drink      string
		Dessert    string
	}

	MealBuilder struct {
		meal Meal // 48 byte
	}
)

func (b *MealBuilder) AddMainCourse(mainCourse string) *MealBuilder {
	b.meal.MainCourse = mainCourse
	return b
}

func (b *MealBuilder) AddDrink(drink string) *MealBuilder {
	b.meal.Drink = drink
	return b
}

func (b *MealBuilder) AddDessert(dessert string) *MealBuilder {
	b.meal.Dessert = dessert
	return b
}

func (b *MealBuilder) Build() Meal {
	return b.meal
}

var (
	builder = &MealBuilder{}
	meal    = builder.AddMainCourse("Steak").AddDrink("Wine").AddDessert("Cake").Build()
)

func main() {
	fmt.Printf("Meal: %+v\n", meal)
}
