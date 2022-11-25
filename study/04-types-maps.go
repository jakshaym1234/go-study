package main

import "fmt"

func main() {
	var countryCapital map[string]string

	countryCapital = make(map[string]string)

	countryCapital["France"] = "Paris"
	countryCapital["India"] = "Delhi"
	countryCapital["Italy"] = "Rome"

	for country := range countryCapital {
		fmt.Printf("Captial of %s is %s\n", country, countryCapital[country])
	}
}
