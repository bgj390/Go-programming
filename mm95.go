// This is the final assignment for http://terokarvinen.com/2020/go-programming-course-2020-w22/#h2
// In this program a stick figure dude is celebrating

package main

import (
	"fmt"
	"time"
)

var p = fmt.Println

func tuuletus1(jee chan bool) {
	
	time.Sleep(400*time.Millisecond)
	p("   _🙂_")
	p(" 🏆|__|⎞")
	p("    ⎭⎩")
	jee <- true
}

func tuuletus2(joo chan bool) {
	
	time.Sleep(200*time.Millisecond)
	p(" 🏆")
	p("  ⎝_😄_⎠")
	p("   |__|")
	p("   〈 〉")
	joo <- true
}

func mm() {
	
	p("**MM95🥇**")
}

func rajahtaa(poks chan bool) {

	fmt.Print("🔅")
	time.Sleep(200*time.Millisecond)
	fmt.Print("🔆")
	time.Sleep(200*time.Millisecond)
	fmt.Print("⭐💥")
	poks <- true	
}

func main() {
	
	p("")	
	mm()
	p("")

	var i int
	for i < 6 {
		jee := make(chan bool)
		go tuuletus1(jee)
		<-jee
		joo := make(chan bool)
		go tuuletus2(joo)
		<-joo
		i++
	}

	for i < 10 {
		poks := make(chan bool)
		go rajahtaa(poks)
		<-poks
		i++
	}

	p("")
	p("")
	mm()
	p("")
}
