package countries

import "fmt"

var myCountry string
var collection []string

// Add adds a country to the collection
func Add(country string) {
	collection = append(collection, country)
}

// SelectCountry selects a country from the collection
func SelectCountry(country string) error {
	if !isInCollection(country) {
		return fmt.Errorf("Country %s not found", country)
	}
	myCountry = country
	return nil
}

// List lists all countries in the collection
func List() {
	for i, country := range collection {
		myCountryMsg := ""
		if country == myCountry {
			myCountryMsg = " (my country)"
		}

		fmt.Printf("%d: %s %s\n", i+1, country, myCountryMsg)
	}
}

// isInCollection checks if a country is in the collection
func isInCollection(country string) bool {
	for _, c := range collection {
		if c == country {
			return true
		}
	}
	return false
}
