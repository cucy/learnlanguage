package main

import "fmt"

func delete_slice() {
	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice) // [Monday Tuesday Wednesday Thursday Friday]

	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice) // [Monday Tuesday Thursday Friday]

}

func main() {
	delete_slice()
}
