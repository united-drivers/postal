# postal
:postbox: international postal address normalizer and renderer

[![CircleCI](https://circleci.com/gh/united-drivers/postal.svg?style=svg)](https://circleci.com/gh/united-drivers/postal)
[![GoDoc](https://godoc.org/github.com/united-drivers/postal?status.svg)](https://godoc.org/github.com/united-drivers/postal)
[![Go Report Card](https://goreportcard.com/badge/github.com/united-drivers/postal)](https://goreportcard.com/report/github.com/united-drivers/postal)
[![License](https://img.shields.io/github/license/united-drivers/postal.svg)](https://github.com/united-drivers/postal/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/united-drivers/postal.svg)](https://github.com/united-drivers/postal/releases)

## Usage

Install with `go get github.com/united-drivers/postal`

```go
import "github.com/united-drivers/postal"

address := postal.Address{
    HouseNumber:  "96",
	Road:         "Boulevard Bessières",
	Suburb:       "Épinettes",
	CityDistrict: "17th Arrondissement",
	County:       "Paris",
	State:        "Ile-de-France",
	Country:      "France",
	PostCode:     "75017",
	CountryCode:  "fr",
}

fmt.Println(address.Short()) // 96 Boulevard Bessières, Paris
fmt.Println(address.Long())  // 96 Boulevard Bessières 75017 Paris, France
fmt.Println(address.Full())  // 96, Boulevard Bessières, Épinettes, 17th Arrondissement, Paris, Ile-de-France, 75017, France
```

More usage on [GoDoc](https://godoc.org/github.com/united-drivers/postal), or in the [test file](https://github.com/united-drivers/postal/blob/master/address_test.go).

## License

[Apache 2.0](https://github.com/united-drivers/postal/blob/master/LICENSE)
