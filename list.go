package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()
	data.PushBack("Mohammad")
	data.PushBack("Nawawi")

	var head *list.Element = data.Front()
	fmt.Println(head.Value)

	next := head.Next()
	fmt.Println(next.Value)

	//ATAU

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
