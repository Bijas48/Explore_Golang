package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// data := ring.New()
	var data *ring.Ring = ring.New(5)
	data.Value = "VALUE satu"
	data = data.Next()

	data.Value = "VALUE dua"
	data = data.Next()

	data.Value = "VALUE tiga"
	data = data.Next()

	data.Value = "VALUE empat"
	data = data.Next()

	data.Value = "VALUE lima"
	data = data.Next()

	data.Value = "VALUE enam"
	data = data.Next()

	data.Do(func(a any) {
		fmt.Println(a)
	})

}
