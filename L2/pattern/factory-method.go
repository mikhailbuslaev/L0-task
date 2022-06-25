package pattern

type Product struct {
	Style string
}

type Fabric interface {
	Create() *Product
}

type ChineseFabric struct {}
type EuropeanFabric struct {}

func (f *ChineseFabric) Create() *Product{
	println("create something chinese...")
	return &Product{Style:"chinese"}
}

func (f *EuropeanFabric) Create() *Product{
	println("create something european...")
	return &Product{Style:"european"}
}

/*
плюсы: 
-можно создавать разные обьекты
минусы:
-под каждый обьект придется писать фабрику 
*/