package main

import (
	"fmt"

	"belajar-go/pkg/database"
	"belajar-go/repository"
)

type Myfriends struct {
	Name    string
	Age     int
	Address string
	Friends []string
}

func main() {

	conn := database.InitPostgres()

	_ = repository.NewTodo(conn)

	var user = Myfriends{}

	user.Name = "myname"
	user.Age = 20
	user.Address = "jl jalan"
	user.Friends = append(user.Friends, "apple")
	user.Friends = append(user.Friends, "grape")
	user.Friends = append(user.Friends, "banana")
	user.Friends = append(user.Friends, "melon")

	fmt.Println(user)

	u := builderName(&Myfriends{Name: "myname", Age: 12, Address: "jl jalan"})

	fmt.Println(u)

}

func builderName(user *Myfriends) string {

	u := fmt.Sprintf("hello: %v, umur kamu : %v", user.Name, user.Age)

	return u
}
