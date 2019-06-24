package main

import "fmt"

func minMax(x, y int) (min, max int) {
	if x > y {
		min, max = y, x
	} else {
		min, max = x, y
	}
	return
}

func funcReturnFunc() (fun func() int) {
	i := 0
	fun = func() (v int) {
		i++
		v = i * i
		return
	}
	return
}

func argFuncPlus(n int) (v int) {
	v = n + n
	return
}

func argFuncMul(n int) (v int) {
	v = n * n
	return
}

func funcArgFunc(f func(int) int, n int) (v int) {
	v = f(n)
	return
}

func main() {
	fmt.Println(minMax(32, 16))
	fmt.Println(minMax(8, 4))
	fmt.Println(minMax(2, 1))
	fmt.Println(minMax(0, 1))

	fun := funcReturnFunc()
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())

	fmt.Println(funcArgFunc(argFuncPlus, 4))
	fmt.Println(funcArgFunc(argFuncMul, 4))
}
