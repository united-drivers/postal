package postal

type Country string

func (c Country) String() string {
	return string(c)
}

func (c Country) Short() string {
	switch c {
	case "United States of America":
		return "USA"
	case "United Kingdom":
		return "UK"
	}
	return string(c)
}
