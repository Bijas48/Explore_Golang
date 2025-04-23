package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `label:"Nama Lengkap"`
	Age  int    `label:"Usia"`
	City string `label:"Kota Tinggal"`
}

func readField(v any) {
	vType := reflect.TypeOf(v)
	fmt.Println("Type Name:", vType.Name())

	for i := 0; i < vType.NumField(); i++ {
		vField := vType.Field(i)
		label := vField.Tag.Get("label")
		fmt.Printf("%s (type: %v) → label: %q\n", vField.Name, vField.Type, label)
	}
}

func main() {
	p := Person{Name: "Adrian", Age: 25, City: "Jakarta"}
	readField(p)

	// Yang ini akan error kalau tidak dihindari, karena pointer
	// v := &Person{Name: "Abi", Age: 30, City: "Bandung"}
	// readField(v)
}
