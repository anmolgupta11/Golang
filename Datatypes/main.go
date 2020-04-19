package main

import "godatatypes/organization"
import "fmt"

func main() {
	p := organization.NewPerson( firstname: "Anmol", lastname: "Gupta")
	err := p.SetTwitterhandler(handler: "@anmol_gupta")
	if err != nil {
		fmt.Printf(format: "An error occurred setting twitter handler: %s\n", err.Error())
	}
	println(p.TwitterHandler())
	println(p.ID())
	println(p.fullname())
}
