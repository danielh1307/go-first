package main

import (
	"fmt"
	"bufio"
	"os"
)

type user struct {
	vorname string
	nachname string
}

func main() {
	// read program arguments
	name := os.Args[1]

	// read value from stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("How old are you?")
	age, _ := reader.ReadString('\n')

	// print something
	fmt.Println("So your name is " + name + " and you are " + age + " years old")

	callSimpleMethod()

	text, value := callMethodWithMultipleReturns(1, 1)
	fmt.Printf("The result of the calculcation is %d and the returned text is "+text+"\n", value)

	var a int
	a = 5
	fmt.Printf("The value of a is %d and its address is %d \n", a, &a)

	var b *int
	b = &a
	fmt.Printf("The value of b is %d and its address is %d \n", *b, b)

	var c = new(int)
	var d = new(int)

	fmt.Printf("%d, %d, %d, %d\n", c, d, &c, &d)

	slice := []int{10, 20, 30}
	fmt.Printf("Length of the slice is %d and capacity is %d\n", len(slice), cap(slice))
	for index, value := range slice {
		fmt.Printf("Index is %d and value is %d \n", index, value)
	}

	for index := 2; index < len(slice); index++ {
		fmt.Printf("Value is %d \n", slice[index])
	}

	// create a slice with make - length is 3 and capacity is 5
	strings := make([]string, 3, 5)
	for _, value := range strings {
		fmt.Printf("The value of the string is %v\n", value)
	}
	// assign values
	strings[0] = "Null"
	strings[1] = "Eins"
	strings[2] = "Zwei"
	for _, value := range strings {
		fmt.Printf("The value of the string is %v\n", value)
	}

	// this call would crash since the length is only 3
	// strings[3] = "Drei"

	strings = append(strings, "Drei")
	strings = append(strings, "Vier")
	fmt.Printf("Now the length is %d and the capacity is %d\n", len(strings), cap(strings))
	strings = append(strings, "Fünf")
	fmt.Printf("Now the length is %d and the capacity is %d\n", len(strings), cap(strings))

	// make a slice of three integers
	numbers := []int{1, 2, 3}
	fmt.Println("Before passing the slice, the values are: ")
	for _, value := range numbers {
		fmt.Printf("%d\n", value)
	}
	newNumbers := changeSlice(numbers)
	fmt.Println("After passing the slice, the values are: ")
	for _, value := range newNumbers {
		fmt.Printf("%d\n", value)
	}
	fmt.Println("The original values are: ")
	for _, value := range numbers {
		fmt.Printf("%d\n", value)
	}

	daniel := user{"daniel", "hamm"}
	fmt.Printf("%v\n", daniel)

	var hugo user;
	hugo.nachname = "Hase"
	hugo.vorname = "Hugo"
	fmt.Printf("%v\n", hugo)

}

func callSimpleMethod() (string) {
	return "This method just returns a string"
}

func callMethodWithMultipleReturns(a int, b int) (string, int) {
	c := a + b
	return "The sum is " + fmt.Sprintf("%d", c), c
}

func changeSlice(slice [] int) [] int {
	// we create a new slice
	innerSlice := slice

	// we change a value in the new slice (innerSlice)
	innerSlice[1] = 7

	// we change a value in the given slice
	slice[1] = 5
	return slice
}
