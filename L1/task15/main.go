package main

var justString string

func someFunc() {
	// в памяти хранится большой массив
	v := createHugeString(1 << 10)
	justString = v[:100]
	// при создании нового массива не происходит копирования значений массива в новой области памяти
	// происходит лишь копирование ссылки на кусок массива
	// после выполнения функции большой массив так и останется в памяти
	// и будет создавать ненужную нагрузку на ресурсы памяти, покуда мы продолжим работать с justString
	// мы можем сделать более экономный способ копирования этого куска:
	justString1 := make([]byte, 0, 100)
	justString2 := make([]byte, 0, 100)
	//this
	copy(justString1, v)
	//or this
	justString2 = append(justString2, v[:100]...)
}

func main() {
	someFunc()
}
