package postal

import (
	"fmt"
	"strings"
)

// warning: be careful when adding new fields, today the compatibility is only done for Nominatim
// if you want to add more drivers (i.e., gmap), please do not add fields, but create a mapping
// function that transforms a gmap struct to Address
// if unsure, ask @moul

type Address struct {
	BusStop       string  `json:"bus_stop"`
	Building      string  `json:"building"`
	HouseNumber   string  `json:"house_number"`
	Road          string  `json:"road"`
	Suburb        string  `json:"suburb"`
	CityDistrict  string  `json:"city_district"`
	City          string  `json:"city"`
	Commercial    string  `json:"commercial"`
	County        string  `json:"county"`
	Hamlet        string  `json:"hamlet"`
	State         string  `json:"state"`
	StateDistrict string  `json:"state_district"`
	Country       Country `json:"country"`
	Town          string  `json:"town"`
	Village       string  `json:"village"`
	PostCode      string  `json:"postcode"`
	CountryCode   string  `json:"country_code"`
}

// MinimalCompare returns a minimal address based on another one
// if a and b doesn't share the same city, it's not a mandatory
// to detail smaller objects than the city
//
// This function should be used when we display from + to addresses on the same screen
func (a *Address) MinimalCompare(b *Address) (string, string) {
	if string(a.Country) != string(b.Country) {
		aRet := fmt.Sprintf("%s, %s", a.SmartCity(), a.Country.Short())
		bRet := fmt.Sprintf("%s, %s", b.SmartCity(), b.Country.Short())
		return aRet, bRet
	}
	// FIXME: add more checks: states, cities, etc
	return a.Short(), b.Short()
}

// Full returns a complete (cannot be more precise) address
func (a *Address) Full() string {
	parts := []string{
		a.BusStop,
		a.Building,
		a.HouseNumber,
		a.Road,
		a.Suburb,
		a.Village,
		a.CityDistrict,
		a.Commercial,
		a.Hamlet,
		a.City,
		a.Town,
		a.County,
		a.StateDistrict,
		a.State,
		a.PostCode,
		string(a.Country),
	}
	return strings.Join(nonEmpty(parts), ", ")
}

// Short returns a minimalist representation of the address
func (a *Address) Short() string {
	var parts []string
	switch a.CountryCode {
	case "fr":
		parts = []string{
			a.SmartRoad(),
			a.SmartCity(),
		}
	case "us":
		parts = []string{
			a.SmartRoad(),
			a.SmartCity(),
			a.State,
		}
	default:
		parts = []string{
			a.SmartRoad(),
			a.SmartCity(),
		}
	}
	return strings.Join(nonEmpty(parts), ", ")
}

func (a *Address) SmartRoad() string {
	switch a.CountryCode {
	case "fr":
		return first(
			a.Building,
			fmt.Sprintf("%s %s", a.HouseNumber, a.Road),
		)
	}
	return first(
		a.Building,
		strings.Join(nonEmpty([]string{a.HouseNumber, a.Road}), ", "),
	)
}

func (a *Address) SmartCity() string {
	return first(
		a.Village,
		a.Town,
		a.City,
		a.County,
		a.CityDistrict,
		a.State,
		a.StateDistrict,
	)
}

// Long returns the a long but not full address
func (a *Address) Long() string {
	switch a.CountryCode {
	case "fr":
		parts := []string{}
		parts = append(parts, a.SmartRoad())
		parts = append(parts, a.PostCode)
		parts = append(parts, a.SmartCity())
		return fmt.Sprintf("%s, %s", strings.Join(nonEmpty(parts), " "), a.Country.Short())
	}
	parts := []string{
		// a.BusStop,
		a.Building,
		a.HouseNumber,
		a.Road,
		a.Suburb,
		a.Village,
		a.CityDistrict,
		// a.Commercial,
		a.Hamlet,
		a.City,
		a.Town,
		a.County,
		a.StateDistrict,
		a.State,
		a.PostCode,
		a.Country.Short(),
	}
	return strings.Join(nonEmpty(parts), ", ")
}

// Letter returns an address in a letter-ready format (multiline)
func (a *Address) Letter() string {
	lines := []string{
		a.Building,
		strings.Join(nonEmpty([]string{a.HouseNumber, a.Road}), ", "),
		strings.Join(nonEmpty([]string{a.PostCode, a.SmartCity()}), " - "),
		string(a.Country),
	}
	return strings.Join(nonEmpty(lines), "\n")
}
