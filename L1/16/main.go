package main

import (
	"fmt"
	"math/rand"
)

func quickSort(slice []int) []int {
	length := len(slice)

	if length <= 1 {
		sliceCopy := make([]int, length)
		copy(sliceCopy, slice)
		return sliceCopy
	}

	m := slice[rand.Intn(length)]

	less := make([]int, 0, length)// делим массив на 3 группы: больше данного элемента, меньше, или равно
	middle := make([]int, 0, length)
	more := make([]int, 0, length)

	for _, item := range slice {
		switch {
		case item < m:
			less = append(less, item)//заполняем здесь эти массивы
		case item == m:
			middle = append(middle, item)
		case item > m:
			more = append(more, item)
		}
	}

	less, more = quickSort(less), quickSort(more)// рекурсивный вызов,
	//в вызываемой функции так же может быть сделан рекурсивный вызов, и получится дерево вызовов
	// ветки завершатся там, где less и more массивы будут пустыми,
	// проверка на пустоту расположена в начале функции, там просто выходим через return
	// особенность этого метода в том, что расчет не контролируется заранее заданной степенью точности,
	// метод сортировки будет выполнятся всегда, но мы не сможем ускорить его за счет уменьшения степени точности

	less = append(less, middle...)
	less = append(less, more...)

	return less // это не группа меньших элементов, это уже пересобранный через append() массив
}

func main() {
	array := []int{98, 51, 45, 56, 58, 728, 11}
	result := quickSort(array)
	fmt.Println(result)
}
