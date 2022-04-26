package main

import "github.com/go_introduction/countries"

func main() {
	countries.Add("United States")
	countries.Add("Canada")
	countries.Add("Mexico")
	countries.Add("Brazil")
	err := countries.SelectCountry("Brazil")

	if err != nil {
		panic(err)
	}
	countries.List()

}
