package main

import "godatatypes/organization"

func main() {
	p := organization.NewPerson( firstname: "Anmol", lastname: "Gupta")
	println(p.ID())
	println(p.fullname())
}
