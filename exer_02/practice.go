package main

import "fmt"

func main() {
	fmt.Println("1:")
	hobbies := [3]string{"Programming", "Watch TV series", "Be with my girlfriend!"}
	fmt.Println(hobbies)

	fmt.Println("2:")
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	fmt.Println("3:")
	sliceHobbies := hobbies[:2]
	fmt.Println(sliceHobbies)

	fmt.Println("4:")
	sliceHobbies3 := sliceHobbies[1:3]
	fmt.Println(sliceHobbies3)

	fmt.Println("5:")
	courseGoals := []string{"Programming in golang", "Make apis in golang"}
	fmt.Println(courseGoals)

	fmt.Println("6:")
	courseGoals[1] = "Be a GoLang Developer"
	courseGoals = append(courseGoals, "Make a lot of money")
	fmt.Println(courseGoals)

	fmt.Println("7:")
	type Product struct {
		id    string
		title string
		price float64
	}
	products := []Product{
		{id: "1", title: "Orange", price: 4.99},
		{id: "2", title: "Banana", price: 2.50},
	}
	products = append(products, Product{id: "3", title: "Jaca", price: 9.80})
	fmt.Println(products)
}
