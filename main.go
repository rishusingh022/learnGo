package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type bot interface {
	getGreeting() string
}

type user struct {
	name string
}

type bot2 interface {
	getGreeting(string, int) (string, error)
	getBotVersion() float64
	respondToUser(user) string
}

// all the requirement should be qualified to become a member of bot2 family

type englishBot struct{}
type spanishBot struct{}

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	// cards := newDeck()
	// cards.print()
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// Assignment solution
	arrayNumber := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, num := range arrayNumber {
		if num%2 == 0 {
			fmt.Printf("Number %v is even\n", num)
		} else {
			fmt.Printf("Number %v is odd\n", num)
		}
	}

	//1st method
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	//2nd method
	alex1 := person{"Alex", "Anderson", contactInfo{}}
	fmt.Println(alex1)

	//3rd method
	var alex3 person
	// fmt.Println(alex3)

	alex3.firstName = "Alex"
	alex3.lastName = "Anderson"
	fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()

	// map

	//var color map[string]string

	//color:= make(map[string]string)

	color := map[string]string{
		"red":  "#00000",
		"blue": "#00000",
	}

	color["white"] = "#00000"

	delete(color, "white")
	printMap(color)

	// interface
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	// make http request
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body)

	// assignment solution
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)

	// channel and routines

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// go keyword is used to create a new routine and run the function in the background and main routine will not wait for the function to complete and it will continue to execute the next line of code

	for _, link := range links {
		go checkLink(link)
	}

}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		return
	}
	fmt.Println(link, "is up")
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
