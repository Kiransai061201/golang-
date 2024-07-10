package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	kiran := User{"Kiran", "kiran@go.dev", true, 16}
	fmt.Println(kiran)
	fmt.Printf("kiran details are: %+v\n", kiran)
	fmt.Printf("Name is %v and email is %v.", kiran.Name, kiran.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
