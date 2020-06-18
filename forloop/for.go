package main

import "fmt"

type person struct {
	firstname  string
	secondname string
	thirdname  string
	age        int
}

type secret struct {
	person
	experience bool
}

func (sa secret) opt() {
	fmt.Println(sa.person.firstname, `experience is`, sa.experience)
}

func main() {
	// for loop
	/* for i := 0; i < 10; i++ {
		fmt.Println(i)
	} */

	i := 0

	for i <= 10 {
		fmt.Println(i)
		i++
	}
	/*
		working with arrays
		var arr [5]int
		arr[0] = 10
		arr[1] = 20
		arr[2] = 30
		arr[3] = 40
		fmt.Println(arr)
		fmt.Println(len(arr))

		// 2d arrays
		var array [2][4]int
		fmt.Println(array[0])
	*/
	person1 := person{
		"lumuli",
		"ken",
		"reagan",
		20,
	}
	person1.fullname()

	agent1 := secret{person{
		"ignatius",
		"wasike",
		"otwala",
		34,
	},
		true,
	}
	agent1.opt()
}

func (p person) fullname() {
	fmt.Println(`agent one fisrt name is`, p.firstname, p.secondname, p.thirdname)
}
