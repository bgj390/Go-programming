// This is the final assignment for http://terokarvinen.com/2020/go-programming-course-2020-w22/#h2
// In this program a stick figure dude is celebrating
// Sources:
// https://gobyexample.com/http-servers
// https://gobyexample.com/timers

package main

import (
	"fmt"
//	"time"
	"net/http"
)


func tuuletus1(w http.ResponseWriter, req *http.Request) {
	
	fmt.Fprintf(w, "   _🙂_\n")
	fmt.Fprintf(w, " 🏆|__|⎞\n")
	fmt.Fprintf(w, "    ⎭⎩\n")
	
}

func tuuletus2(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, " 🏆\n")
	fmt.Fprintf(w, "  ⎝_😄_⎠\n")
	fmt.Fprintf(w, "   |__|\n")
	fmt.Fprintf(w, "   〈 〉\n")
	
}

func mm(w http.ResponseWriter, req *http.Request) {
	
	fmt.Fprintf(w, "**MM95🥇**\n")
}

func rajahtaa(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "🔅")
	fmt.Fprintf(w, "🔆")
	fmt.Fprintf(w, "⭐💥")
		
}

func main() {

	fmt.Println("")	

	http.HandleFunc("/mm", mm)

	fmt.Println("")

	http.HandleFunc("/mmt1", tuuletus1)
	http.HandleFunc("/mmt2", tuuletus2)		
	http.HandleFunc("/mmrajah", rajahtaa)	

	fmt.Println("")

	http.HandleFunc("/mm2", mm)

	http.ListenAndServe(":8090", nil)
}
