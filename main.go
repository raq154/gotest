package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var colors = map[string]string{
	"red": "#xxeasd",
}

type bot interface {
	getGreeting() string
}

type SpanishBot struct{}
type EnglishBot struct{}
type JapaneseBot struct{}

type FuncType func(int) int

func (SpanishBot) getGreeting() string {
	return "Hola"
}

func (EnglishBot) getGreeting() string {
	return "Hi There"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

type Triangle struct {
	height float64
	base   float64
}

type Square struct {
	sideLength float64
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (t Square) getArea() float64 {
	return t.sideLength * 2
}

type Shape interface {
	getArea() float64
}

func main() {

	//cards := newDeck()
	// filepath := "/Users/raheelarif/Desktop/Cards3.txt"
	// err := cards.saveToFile(filepath)
	// if err != nil {
	// 	fmt.Println("Unable to save file ", err)
	// 	os.Exit(1)
	// }

	// cardsFromFile, err := readFromFile(filepath)
	// if err != nil {
	// 	fmt.Println("Unable to read from file ", err)
	// 	os.Exit(1)
	// }
	//cards.shuffle()
	//cards.print()

	// person := Person{
	// 	firstName: "raheel",
	// 	lastName:  "arif",
	// 	contactInfo: ContactInfo{
	// 		email:   "raheelarif@gmail.com",
	// 		zipcode: "L970X7",
	// 	},
	// }
	// person.print()
	// person.updateName("RAHEEL")
	// person.print()

	// personPointer := &person
	// fmt.Printf("pinter to person %p \n", personPointer)
	// fmt.Printf("value of pinter to person %s \n", *personPointer)

	//colors["orange"] = "#ORANGE"

	// for k, v := range colors {
	// 	fmt.Printf("colors %s %s \n", k, v)
	// }

	//printGreeting(SpanishBot{})
	//printGreeting(EnglishBot{})

	// helloHandler := func(w http.ResponseWriter, req *http.Request) {
	// 	io.WriteString(w, "Hello, world!\n")
	// }

	// http.HandleFunc("/hello", helloHandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// res, err := http.Get("http://www.google.com")
	// if err != nil {
	// 	log.Fatal("Unable to connect to google.com", err)
	// 	os.Exit(1)
	// }
	// resStr, errReadingBody := io.ReadAll(res.Body)
	// if errReadingBody != nil {
	// 	log.Fatal("Unable to read body of response", errReadingBody)
	// 	os.Exit(1)
	// }

	// log.Printf("Got response %s", string(resStr))

	// squareArea := Square{sideLength: 10}.getArea()
	// log.Println("area is ", squareArea)

	// triangleArea := Triangle{height: 10, base: 15}.getArea()
	// log.Println("area is ", triangleArea)

	// arguments := os.Args
	// if len(arguments) != 2 {
	// 	log.Fatal("Unexpected number of arguments")
	// 	os.Exit(1)
	// }

	// filePath := os.Args[1]
	// file, err := os.Open(filePath)
	// if err != nil {
	// 	log.Fatal("unable to open file at path", filePath)
	// 	os.Exit(1)
	// }

	// content, readingErr := io.ReadAll(io.Reader(file))
	// if readingErr != nil {
	// 	log.Fatal("unable to open file at path", filePath)
	// 	os.Exit(1)
	// }
	// log.Println("filecontents ", string(content))

	// var websitesToCheck = [...]string{"http://google.com", "http://facebook.com", "http://stackoverflow.com"}

	// c := make(chan string)
	// sleepFunc := func() {
	// 	time.Sleep(time.Duration(time.Second))
	// }

	// for _, website := range websitesToCheck {
	// 	go checkWebsite(website, c, sleepFunc)
	// }

	// for link := range c {
	// 	go checkWebsite(link, c, sleepFunc)
	// }

	// cmd := &cobra.Command{
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		fmt.Println("Hello, Modules!")

	// 		mypackage.PrintHello()
	// 	},
	// }
	// cmd.Execute()

	size, status := niceFuncWithMultipleReturns()
	fmt.Println(size, status)

	var v interface{} /** similar to object in java **/
	v = "Hello"
	v = []int{1, 2, 3}
	v = 42
	fmt.Print(v)

	strValue, ok := v.(string)
	if ok {
		fmt.Println(strValue)
	}

	if intValue, ok := v.(int); ok {
		fmt.Println("v is an integer:", intValue)
	} else if floatValue, ok := v.(float64); ok {
		fmt.Println("v is an float:", floatValue)
	} else if strVal, ok := v.(string); ok {
		fmt.Println("v is an float:", strVal)
	}

	switch value := v.(type) {
	case int:
		fmt.Println("v is an integer:", value)
	case float64:
		fmt.Println("v is a float:", value)
	case string:
		fmt.Println("v is a string:", value)
	default:
		fmt.Println("v is of unknown type")
	}

}

const HTTP_SUCESS = 200

func checkWebsite(website string, c chan string, onComplete func()) {
	res, err := http.Get(website)
	if err != nil {
		fmt.Println("Unable to read response " + website)
	} else if res.StatusCode != HTTP_SUCESS {
		fmt.Println("Website not running :( " + website)
	} else {
		fmt.Println("Website looks good :) " + website)
	}
	onComplete()
	c <- website
}

func niceFuncWithMultipleReturns() (size int64, status string) {
	return 100, "Thank you for your purchase"
}

func init() {
	http.HandleFunc("/", errorHandler(betterHandler))
}

func errorHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("handling %q: %v", r.RequestURI, err)
		}
	}
}

func betterHandler(w http.ResponseWriter, r *http.Request) error {
	if err := doThis(); err != nil {
		return fmt.Errorf("doing this: %v", err)
	}

	if err := doThat(); err != nil {
		return fmt.Errorf("doing that: %v", err)
	}
	return nil
}

func doThis() error {
	return nil
}

func doThat() error {
	return nil
}

type Server struct {
	quit chan bool
}

func NewServer() *Server {
	s := &Server{make(chan bool)}
	go s.run()
	return s
}

func (s *Server) run() {
	for {
		select {
		case <-s.quit:
			fmt.Println("finishing task")
			time.Sleep(time.Second)
			fmt.Println("task done")
			s.quit <- true
			return
		case <-time.After(time.Second):
			fmt.Println("running task")
		}
	}
}
