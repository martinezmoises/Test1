//Filename: main.go

package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// package level variable that points to a template definition from the provided files.
var tpl = template.Must(template.ParseFiles("index.html")) //parses the index.html file in the root of our project directory and validates it.

// Create a global struct named
// struct to store data and pass it to your template for display on your final page.
type Webpage struct {
	Headline string
	Body     string //contain the
	Time     string
	Quotes   string
}

// "K" for "Key" and "V" for "Value". These types are used to define the key and value types for the map that will be used later in the program.
type K string
type V string

func randomQuote(m map[K]V) V { // Takes a map of type map[K]V as its argument and returns a value of type V.
	k := rand.Intn(len(m)) // uses the "rand.Intn" function from the "math/rand" package to generate a random integer between 0 and the length of the map.
	for _, x := range m {  // It then iterates over the values in the map using a range loop, decrementing the random integer each time it encounters a new value.
		if k == 0 {
			return x // When the random integer reaches 0, the function returns the current value.
		}
		k--
	}
	panic("unreachable") // If the loop completes without finding a value to return, the function panics with an "unreachable" error message.
}

func home(w http.ResponseWriter, r *http.Request) { //output to the ResponseWriter

	//Populating the struct Webpage with the appropriate data
	parts := Webpage{
		Headline: "Welcome to my Test 2",
		Body:     " My name is Moises Martinez. I enjoy swimming and excersing. I also enjoying spending time with my family and friends. I dislike when people whisper for some reason. My email address is: 2021154670@ub.edu.bz",
	}

	//tpl template is executed by providing two arguments: where we want to write the output to, and the data we want to pass to the template.
	err := tpl.Execute(w, parts)

	if err != nil {
		return
	}

	w.Write([]byte("<center><p>This is my teacher, his name is Dalwin lewis and unfortunatley he is leaving UB this year :( </p></center>"))
	w.Write([]byte("<center><img src= https://doit.ub.edu.bz/pluginfile.php/237/user/icon/eguru/f1?rev=87989 alt=Dalwin></center>"))
}

func greeting(w http.ResponseWriter, r *http.Request) {

	//Variable that holds the current time using .Now function
	currentTime := time.Now()

	//Populating struct with adequate data
	parts := Webpage{
		Time: " The current time right now is: " + currentTime.Format("15:04:05 ") + " continue having a good day",
	}

	err := tpl.Execute(w, parts) //Execute the template

	if err != nil {
		return
	}

	w.Write([]byte("<center><img src= https://thumbs.dreamstime.com/b/funny-man-eating-break-time-message-211128379.jpg alt=meme ></center>"))

}

func random(w http.ResponseWriter, r *http.Request) {

	//Map that holds 5 quotes
	m := map[K]V{
		"quote1": "The most difficult thing is the decision to act. The rest is merely tenacity. The fears are paper tigers. You can do anything you decide to do. You can act to change and control your life; and the procedure, the process, is its own reward.",
		"quote2": "Spit on your hands and get busy. Your blood will start circulating; your mind will start ticking—and pretty soon this whole positive upsurge of life in your body will drive worry from your mind. Get busy. Keep busy. It’s the cheapest kind of medicine there is on this earth—and one of the best. - Dale Carnegie",
		"quote3": "To forgive is to set a prisoner free and discover that the prisoner was you. - Lewis B. Smedes",
		"quote4": "Where you are is a result of who you were, but where you go depends entirely on who you choose to be. - Hal Elrod (check source)",
		"quote5": "The man who asks a question is a fool for a minute. The man who does not ask is a fool for a minute",
	}

	//Creating a variable and passing in the function to get random quote and passing in the map for the parameter
	quote := randomQuote(m)

	//Populating the struct with the quotes
	parts := Webpage{
		Quotes: string(quote), //Used string() since Quotes in struct is type string
	}

	err := tpl.Execute(w, parts) //Execute the template

	if err != nil {
		return
	}

	w.Write([]byte("<center><img src= https://media4.giphy.com/media/Rlwz4m0aHgXH13jyrE/200w.gif?cid=6c09b952c4p30z2oa04zrwf8evy7t80ar9e336f92fb52w0k&rid=200w.gif&ct=g alt=meme ></center>"))

}

func main() {

	//The client sends a request to the web server the sever then sends the request to the multiplexer which then sends it to the handler
	//This is how the request is processed
	//Client->Server->Multiplexer->Handler
	//This is how the response is sent back to the server
	//Handler->Multiplexer->Server->Client

	//Multiplexer(map)
	mux := http.NewServeMux() //method used to create an HTTP request multiplexer which is assigned to the mux variable.
	mux.HandleFunc("/", home) //a request multiplexer matches the URL of incoming requests against a list of registered patterns,
	// and calls the associated handler for the pattern whenever a match is found.
	mux.HandleFunc("/greeting", greeting) //Bascially where to go
	mux.HandleFunc("/random", random)
	log.Print("starting server on :4000")    //Just printing an the server address
	err := http.ListenAndServe(":4000", mux) //method which starts the server on the port defined
	log.Fatal(err)

}
