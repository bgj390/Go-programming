package main

import "fmt"
import "flag"


func main() {

	var ruoka string
	flag.StringVar(&ruoka, "name", "vettä ja leipää", "Mitä tahtoisit syödä?")
	flag.Parse()

	fmt.Println("Tänään syödään", ruoka)
}
