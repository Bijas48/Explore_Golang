package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("hai")
	data.PushBack("boom")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
