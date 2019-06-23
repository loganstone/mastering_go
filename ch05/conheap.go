package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type books []string

func (bs *books) Pop() interface{} {
	old := *bs
	last := old[len(old)-1]
	*bs = old[0 : len(old)-1]
	return last
}

func (bs *books) Push(b interface{}) {
	*bs = append(*bs, b.(string))
}

func (bs books) Len() int {
	return len(bs)
}

func (bs books) Less(a, b int) bool {
	return bs[a] < bs[b]
}

func (bs books) Swap(a, b int) {
	bs[a], bs[b] = bs[b], bs[a]
}

func main() {
	bookshelf := &books{
		"이산수학",
		"코딩 인터뷰 완전분석",
		"모던 C++ 입문",
		"C 기초 플러스",
		"CODE COMPLETE 2",
	}
	heap.Init(bookshelf)
	fmt.Println("books:", bookshelf)
	fmt.Println("books size:", len(*bookshelf))

	fmt.Println("take out a book:", bookshelf.Pop())
	fmt.Println("books:", bookshelf)
	fmt.Println("books size:", len(*bookshelf))

	bookshelf.Push("The Go Programming Language")
	fmt.Println("put a book:", "The Go Programming Language")
	bookshelf.Push("Mastering Go")
	fmt.Println("put a book:", "Mastering Go")
	fmt.Println("books:", bookshelf)
	fmt.Println("books size:", len(*bookshelf))

	heap.Init(bookshelf)
	fmt.Println("books:", bookshelf)
	fmt.Println("books size:", len(*bookshelf))

	fmt.Println(strings.Repeat("-", 10))
	for bookshelf.Len() > 0 {
		fmt.Printf("%s \n", heap.Pop(bookshelf))
	}
}
