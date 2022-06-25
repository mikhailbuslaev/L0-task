package pattern

type Meal struct {
	Salad string
	Soup string
	Dessert string
}

type MealBuilder interface {
	BuildSalad() string
	BuildSoup() string
	BuildDessert() string
	BuildMeal() *Meal
}

type Director {
	*MealBuilder
}

type Builder1 struct{}

type Builder2 struct{}

func (d *Director) ExecuteOrder() {
	println("Your order id done!")
	return d.b.BuildMeal()
}

func (b *Buider1) BuildSalad() string{
	return "salad type 1"
}

func (b *Buider1) BuildSoup() string{
	return "soup type 1"
}

func (b *Buider1) BuildDessert() string{
	return "dessert type 1"
}

func (b *Buider1) BuildMeal() *Meal{
	m := &Meal{}
	m.Salad = b.BuildSalad()
	m.Soup = b.BuildSoup()
	m.Dessert = b.BuildDessert()
	return m
}

func (b *Buider2) BuildSalad() string{
	return "salad type 2"
}

func (b *Buider2) BuildSoup() string{
	return "soup type 2"
}

func (b *Buider2) BuildDessert() string{
	return "dessert type 2"
}

func (b *Buider2) BuildMeal() *Meal{
	m := &Meal{}
	m.Soup = b.BuildSoup()
	m.Salad = b.BuildSalad()
	m.Dessert = b.BuildDessert()
	return m
}


/*
плюсы:
-позволяет создавать сложные и вариативные обьекты
-инкапсулирует создание обьекта от области, где мы его применим
-можно для каждого конкретбилдера написать свой способ создания

минусы:
-много кода, каждыый метод конкретбилдера пишется самостоятельно
 */