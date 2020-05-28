// This is the final assignment for http://terokarvinen.com/2020/go-programming-course-2020-w22/#h2
// In this program a stick figure dude is celebrating
// Sources:
// https://gobyexample.com/http-servers
// https://gobyexample.com/timers

package main

import (
	"fmt"
	"time"
	"net/http"
)

var p = fmt.Println

func ohjelma() {
	
	p("**MM95🥇**")

	var i int
	for i < 6 {
		time.Sleep(200*time.Millisecond)
		p("")
		p("   _🙂_")
		p(" 🏆(__)⎞")
		p("    ⎭⎩")
		time.Sleep(400*time.Millisecond)
		p(" 🏆")
        	p("  ⎝_😄_⎠")
        	p("   (__)")
        	p("   〈 〉")
		i++
	}
	
	for i < 10 {
		time.Sleep(200*time.Millisecond)
		fmt.Print("🔅")
        	time.Sleep(200*time.Millisecond)
		fmt.Print("🔆")
		time.Sleep(200*time.Millisecond)
        	p(" 💥")
	}
	
	p("**MM95🥇**")
}

func net(w http.ResponseWriter, req *http.Request) {

	ohjelma()
}

func main() {

	http.HandleFunc("/mm95", net)

	http.ListenAndServe(":8090", nil)
}
