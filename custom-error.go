package main

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message

}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message

}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "id tidak ditemukan"}
	}
	if id != "test" {
		return &notFoundError{Message: "data tidak ditemukan"}
	}
	return nil
}

// func main() {
// 	err := SaveData("test", nil)
// 	if err != nil {
// 		// if validationErr, ok := err.(*validationError); ok {
// 		// 	fmt.Println("validation error :", validationErr.Error())
// 		// } else if notFoundErr, ok := err.(*notFoundError); ok {
// 		// 	fmt.Println("notfound error :", notFoundErr.Error())
// 		// } else {
// 		// 	fmt.Println("unknown", err.Error())
// 		// }

// 		switch finalError := err.(type) {
// 		case *validationError:
// 			fmt.Println("validation error :", finalError.Error())

// 		case *notFoundError:
// 			fmt.Println("notfound error :", finalError.Error())
// 		default:
// 			fmt.Println("unknown error :", finalError.Error())

// 		}

// 	} else {
// 		fmt.Println("sukses")
// 	}
// }
