package postal

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleAddress_fr() {
	address := Address{
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
	fmt.Println(address.Short())
	fmt.Println(address.Long())
	fmt.Println(address.Full())
	// Output:
	// 96 Boulevard Bessières, Paris
	// 96 Boulevard Bessières 75017 Paris, France
	// 96, Boulevard Bessières, Épinettes, 17th Arrondissement, Paris, Ile-de-France, 75017, France
}

func ExampleAddress_Short_fr() {
	address := Address{
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
	fmt.Println(address.Short())
	// Output: 96 Boulevard Bessières, Paris
}

func ExampleAddress_Long_fr() {
	address := Address{
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
	fmt.Println(address.Long())
	// Output: 96 Boulevard Bessières 75017 Paris, France
}

func ExampleAddress_Full_fr() {
	address := Address{
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
	fmt.Println(address.Full())
	// Output: 96, Boulevard Bessières, Épinettes, 17th Arrondissement, Paris, Ile-de-France, 75017, France
}

func ExampleAddress_Letter_fr() {
	address := Address{
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
	fmt.Println(address.Letter())
	// Output:
	// 96, Boulevard Bessières
	// 75017 - Paris
	// France
}

func Test(t *testing.T) {
	type NominatimResponse struct {
		DisplayName string  `json:"display_name"`
		Address     Address `json:"address"`
	}
	Convey("Testing package", t, FailureContinues, func() {
		adds := make([]NominatimResponse, 9)

		// https://nominatim.openstreetmap.org/reverse?format=json&lat=52.5487429714954&lon=-1.81602098644987&zoom=18&addressdetails=1
		So(json.Unmarshal([]byte(`{"place_id":"91015268","licence":"Data © OpenStreetMap contributors, ODbL 1.0. http:\/\/www.openstreetmap.org\/copyright","osm_type":"way","osm_id":"90394420","lat":"52.54877605","lon":"-1.81627033283164","display_name":"137, Pilkington Avenue, Sutton Coldfield, Birmingham, West Midlands Combined Authority, West Midlands, England, B72 1LH, United Kingdom","address":{"house_number":"137","road":"Pilkington Avenue","suburb":"Sutton Coldfield","city":"Birmingham","county":"West Midlands Combined Authority","state_district":"West Midlands","state":"England","postcode":"B72 1LH","country":"United Kingdom","country_code":"gb"},"boundingbox":["52.5487321","52.5488299","-1.8163514","-1.8161885"]}`), &adds[0]), ShouldBeNil)
		// https://nominatim.openstreetmap.org/reverse.php/?format=json&lat=38.45547070594816&lon=-107.02606201171875&addressdetails=1
		So(json.Unmarshal([]byte(`{"place_id":"68990930","licence":"Data © OpenStreetMap contributors, ODbL 1.0. http:\/\/www.openstreetmap.org\/copyright","osm_type":"way","osm_id":"17064934","lat":"38.458648","lon":"-107.032343","display_name":"Big Mesa Road, Gunnison County, Colorado, United States of America","address":{"road":"Big Mesa Road","county":"Gunnison County","state":"Colorado","country":"United States of America","country_code":"us"},"boundingbox":["38.4473483","38.472006","-107.047445","-107.0128572"]}`), &adds[1]), ShouldBeNil)
		// https://nominatim.openstreetmap.org/reverse.php?format=json&lat=42&lon=42&zoom=
		So(json.Unmarshal([]byte(`{"place_id":"91453566","licence":"Data © OpenStreetMap contributors, ODbL 1.0. http:\/\/www.openstreetmap.org\/copyright","osm_type":"way","osm_id":"101162640","lat":"41.9972094","lon":"41.9903725","display_name":"Silauri, Ozurgeti Municipality, Guria, Georgia","address":{"village":"Silauri","county":"Ozurgeti Municipality","state":"Guria","country":"Georgia","country_code":"ge"},"boundingbox":["41.9862842","42.0008499","41.9795922","42.0015527"]}`), &adds[2]), ShouldBeNil)
		// 42: https://nominatim.openstreetmap.org/reverse?format=json&lat=48.8966358&lon=2.31843371362269&zoom=18&addressdetails=1
		So(json.Unmarshal([]byte(`{"place_id":"179826025","licence":"Data © OpenStreetMap contributors, ODbL 1.0. http:\/\/www.openstreetmap.org\/copyright","osm_type":"relation","osm_id":"3957506","lat":"48.896618","lon":"2.31838131502914","display_name":"96, Boulevard Bessières, Épinettes, 17th Arrondissement, Paris, Ile-de-France, Metropolitan France, 75017, France","address":{"house_number":"96","road":"Boulevard Bessières","suburb":"Épinettes","city_district":"17th Arrondissement","county":"Paris","state":"Ile-de-France","country":"France","postcode":"75017","country_code":"fr"},"boundingbox":["48.8964122","48.8969086","2.3181569","2.3188611"]}`), &adds[3]), ShouldBeNil)
		// st laz: https://nominatim.openstreetmap.org/reverse?format=json&lat=48.87653&lon=2.3252374&zoom=18&addressdetails=1
		So(json.Unmarshal([]byte(`{"place_id":"85927328","licence":"Data © OpenStreetMap contributors, ODbL 1.0. http:\/\/www.openstreetmap.org\/copyright","osm_type":"way","osm_id":"67825345","lat":"48.87714825","lon":"2.3247649469037","display_name":"Gare Saint-Lazare, Parking Saint-Lazare, Europe, 8th Arrondissement of Paris, Paris, Ile-de-France, Metropolitan France, 75008, France","address":{"building":"Gare Saint-Lazare","road":"Parking Saint-Lazare","suburb":"Europe","city_district":"8th Arrondissement of Paris","county":"Paris","state":"Ile-de-France","country":"France","postcode":"75008","country_code":"fr"},"boundingbox":["48.876261","48.8780666","2.3232934","2.3266729"]}`), &adds[4]), ShouldBeNil)
		// https://nominatim.openstreetmap.org/reverse?format=json&lat=48.8785578&lon=2.3640929&zoom=18&addressdetails=1
		So(json.Unmarshal([]byte(`{"place_id":"14662916","licence":"Data © OpenStreetMap contributors, ODbL 1.0. http:\/\/www.openstreetmap.org\/copyright","osm_type":"node","osm_id":"1279610949","lat":"48.8785578","lon":"2.3640929","display_name":"15, Rue Eugène Varlin, Hôpital-St-Louis, 10th Arrondissement, Campagne à Paris, Paris, Ile-de-France, Metropolitan France, 75010, France","address":{"house_number":"15","road":"Rue Eugène Varlin","suburb":"Hôpital-St-Louis","city_district":"10th Arrondissement","hamlet":"Campagne à Paris","county":"Paris","state":"Ile-de-France","country":"France","postcode":"75010","country_code":"fr"},"boundingbox":["48.8784578","48.8786578","2.3639929","2.3641929"]}`), &adds[5]), ShouldBeNil)
		// https://nominatim.openstreetmap.org/reverse?format=json&lat=43.3369407&lon=1.1449023&zoom=18&addressdetails=1
		So(json.Unmarshal([]byte(`{"place_id":"107350828","licence":"Data © OpenStreetMap contributors, ODbL 1.0. http:\/\/www.openstreetmap.org\/copyright","osm_type":"way","osm_id":"156074292","lat":"43.3340195","lon":"1.1404921","display_name":"D 73a, Morin, Gratens, Muret, Haute-Garonne, Occitania, Metropolitan France, 31430, France","address":{"road":"D 73a","suburb":"Morin","village":"Gratens","county":"Muret","state":"Occitania","country":"France","postcode":"31430","country_code":"fr"},"boundingbox":["43.3214383","43.3436739","1.1349576","1.1498949"]}`), &adds[6]), ShouldBeNil)
		// https://nominatim.openstreetmap.org/reverse?format=json&lat=43.8251492&lon=5.7885946&zoom=18&addressdetails=1
		So(json.Unmarshal([]byte(`{"place_id":"19952168","licence":"Data © OpenStreetMap contributors, ODbL 1.0. http:\/\/www.openstreetmap.org\/copyright","osm_type":"node","osm_id":"1897135018","lat":"43.8251492","lon":"5.7885946","display_name":"Lattre de Tassigny, Avenue Maréchal de Lattre de Tassigny, Zone commerciale Mistral, La Baronne, Manosque, Forcalquier, Alpes-de-Haute-Provence, Provence-Alpes-Côte d'Azur, Metropolitan France, 04100, France","address":{"bus_stop":"Lattre de Tassigny","road":"Avenue Maréchal de Lattre de Tassigny","commercial":"Zone commerciale Mistral","suburb":"La Baronne","town":"Manosque","county":"Forcalquier","state":"Provence-Alpes-Côte d'Azur","country":"France","postcode":"04100","country_code":"fr"},"boundingbox":["43.8250492","43.8252492","5.7884946","5.7886946"]}`), &adds[7]), ShouldBeNil)
		// https://nominatim.openstreetmap.org/reverse?format=json&lat=48.8635833&lon=2.2871285&zoom=18&addressdetails=1
		So(json.Unmarshal([]byte(`{"place_id":"8984951","licence":"Data © OpenStreetMap contributors, ODbL 1.0. http:\/\/www.openstreetmap.org\/copyright","osm_type":"node","osm_id":"932277394","lat":"48.8637222","lon":"2.2870818","display_name":"Le Malakoff, Avenue Raymond Poincaré, Chaillot, 16e, Paris, Île-de-France, France métropolitaine, 75116, France","address":{"cafe":"Le Malakoff","road":"Avenue Raymond Poincaré","suburb":"Chaillot","city_district":"16e","city":"Paris","county":"Paris","state":"Île-de-France","country":"France","postcode":"75116","country_code":"fr"},"boundingbox":["48.8636222","48.8638222","2.2869818","2.2871818"]}`), &adds[8]), ShouldBeNil)

		So(adds[0].DisplayName, ShouldEqual, "137, Pilkington Avenue, Sutton Coldfield, Birmingham, West Midlands Combined Authority, West Midlands, England, B72 1LH, United Kingdom")
		So(adds[0].Address.Full(), ShouldEqual, "137, Pilkington Avenue, Sutton Coldfield, Birmingham, West Midlands Combined Authority, West Midlands, England, B72 1LH, United Kingdom")
		So(adds[0].Address.Short(), ShouldEqual, "137, Pilkington Avenue, Birmingham")
		So(adds[0].Address.Long(), ShouldEqual, "137, Pilkington Avenue, Sutton Coldfield, Birmingham, West Midlands Combined Authority, West Midlands, England, B72 1LH, UK")
		So(adds[0].Address.Letter(), ShouldEqual, "137, Pilkington Avenue\nB72 1LH - Birmingham\nUnited Kingdom")

		So(adds[1].DisplayName, ShouldEqual, "Big Mesa Road, Gunnison County, Colorado, United States of America")
		So(adds[1].Address.Full(), ShouldEqual, "Big Mesa Road, Gunnison County, Colorado, United States of America")
		So(adds[1].Address.Short(), ShouldEqual, "Big Mesa Road, Gunnison County, Colorado")
		So(adds[1].Address.Long(), ShouldEqual, "Big Mesa Road, Gunnison County, Colorado, USA")
		So(adds[1].Address.Letter(), ShouldEqual, "Big Mesa Road\nGunnison County\nUnited States of America")

		So(adds[2].DisplayName, ShouldEqual, "Silauri, Ozurgeti Municipality, Guria, Georgia")
		So(adds[2].Address.Full(), ShouldEqual, "Silauri, Ozurgeti Municipality, Guria, Georgia")
		So(adds[2].Address.Short(), ShouldEqual, "Silauri")
		So(adds[2].Address.Long(), ShouldEqual, "Silauri, Ozurgeti Municipality, Guria, Georgia")
		So(adds[2].Address.Letter(), ShouldEqual, "Silauri\nGeorgia")

		So(adds[3].DisplayName, ShouldEqual, "96, Boulevard Bessières, Épinettes, 17th Arrondissement, Paris, Ile-de-France, Metropolitan France, 75017, France")
		So(adds[3].Address.Full(), ShouldEqual, "96, Boulevard Bessières, Épinettes, 17th Arrondissement, Paris, Ile-de-France, 75017, France")
		So(adds[3].Address.Short(), ShouldEqual, "96 Boulevard Bessières, Paris")
		So(adds[3].Address.Long(), ShouldEqual, "96 Boulevard Bessières 75017 Paris, France")
		So(adds[3].Address.Letter(), ShouldEqual, "96, Boulevard Bessières\n75017 - Paris\nFrance")

		So(adds[4].DisplayName, ShouldEqual, "Gare Saint-Lazare, Parking Saint-Lazare, Europe, 8th Arrondissement of Paris, Paris, Ile-de-France, Metropolitan France, 75008, France")
		So(adds[4].Address.Full(), ShouldEqual, "Gare Saint-Lazare, Parking Saint-Lazare, Europe, 8th Arrondissement of Paris, Paris, Ile-de-France, 75008, France")
		So(adds[4].Address.Short(), ShouldEqual, "Gare Saint-Lazare, Paris")
		So(adds[4].Address.Long(), ShouldEqual, "Gare Saint-Lazare 75008 Paris, France")
		So(adds[4].Address.Letter(), ShouldEqual, "Gare Saint-Lazare\nParking Saint-Lazare\n75008 - Paris\nFrance")

		So(adds[5].DisplayName, ShouldEqual, "15, Rue Eugène Varlin, Hôpital-St-Louis, 10th Arrondissement, Campagne à Paris, Paris, Ile-de-France, Metropolitan France, 75010, France")
		So(adds[5].Address.Full(), ShouldEqual, "15, Rue Eugène Varlin, Hôpital-St-Louis, 10th Arrondissement, Campagne à Paris, Paris, Ile-de-France, 75010, France")
		So(adds[5].Address.Short(), ShouldEqual, "15 Rue Eugène Varlin, Paris")
		So(adds[5].Address.Long(), ShouldEqual, "15 Rue Eugène Varlin 75010 Paris, France")
		So(adds[5].Address.Letter(), ShouldEqual, "15, Rue Eugène Varlin\n75010 - Paris\nFrance")

		So(adds[6].DisplayName, ShouldEqual, "D 73a, Morin, Gratens, Muret, Haute-Garonne, Occitania, Metropolitan France, 31430, France")
		So(adds[6].Address.Full(), ShouldEqual, "D 73a, Morin, Gratens, Muret, Occitania, 31430, France")
		So(adds[6].Address.Short(), ShouldEqual, "D 73a, Gratens")
		So(adds[6].Address.Long(), ShouldEqual, "D 73a 31430 Gratens, France")
		So(adds[6].Address.Letter(), ShouldEqual, "D 73a\n31430 - Gratens\nFrance")

		So(adds[7].DisplayName, ShouldEqual, "Lattre de Tassigny, Avenue Maréchal de Lattre de Tassigny, Zone commerciale Mistral, La Baronne, Manosque, Forcalquier, Alpes-de-Haute-Provence, Provence-Alpes-Côte d'Azur, Metropolitan France, 04100, France")
		So(adds[7].Address.Full(), ShouldEqual, "Lattre de Tassigny, Avenue Maréchal de Lattre de Tassigny, La Baronne, Zone commerciale Mistral, Manosque, Forcalquier, Provence-Alpes-Côte d'Azur, 04100, France")
		So(adds[7].Address.Short(), ShouldEqual, "Avenue Maréchal de Lattre de Tassigny, Manosque")
		So(adds[7].Address.Long(), ShouldEqual, "Avenue Maréchal de Lattre de Tassigny 04100 Manosque, France")
		So(adds[7].Address.Letter(), ShouldEqual, "Avenue Maréchal de Lattre de Tassigny\n04100 - Manosque\nFrance")

		So(adds[8].DisplayName, ShouldEqual, "Le Malakoff, Avenue Raymond Poincaré, Chaillot, 16e, Paris, Île-de-France, France métropolitaine, 75116, France")
		So(adds[8].Address.Full(), ShouldEqual, "Avenue Raymond Poincaré, Chaillot, 16e, Paris, Paris, Île-de-France, 75116, France")
		So(adds[8].Address.Short(), ShouldEqual, "Avenue Raymond Poincaré, Paris")
		So(adds[8].Address.Long(), ShouldEqual, "Avenue Raymond Poincaré 75116 Paris, France")
		So(adds[8].Address.Letter(), ShouldEqual, "Avenue Raymond Poincaré\n75116 - Paris\nFrance")

		testMinimalCompare := func(aIdx, bIdx int) string {
			a := adds[aIdx]
			b := adds[bIdx]
			retA, retB := a.Address.MinimalCompare(&b.Address)
			return fmt.Sprintf("%s|%s", retA, retB)
		}

		So(testMinimalCompare(0, 0), ShouldEqual, "137, Pilkington Avenue, Birmingham|137, Pilkington Avenue, Birmingham")
		So(testMinimalCompare(0, 1), ShouldEqual, "Birmingham, UK|Gunnison County, USA")
		So(testMinimalCompare(0, 2), ShouldEqual, "Birmingham, UK|Silauri, Georgia")
		So(testMinimalCompare(0, 3), ShouldEqual, "Birmingham, UK|Paris, France")
		So(testMinimalCompare(0, 4), ShouldEqual, "Birmingham, UK|Paris, France")
		So(testMinimalCompare(0, 5), ShouldEqual, "Birmingham, UK|Paris, France")
		So(testMinimalCompare(0, 6), ShouldEqual, "Birmingham, UK|Gratens, France")
		So(testMinimalCompare(0, 7), ShouldEqual, "Birmingham, UK|Manosque, France")

		So(testMinimalCompare(1, 0), ShouldEqual, "Gunnison County, USA|Birmingham, UK")
		So(testMinimalCompare(1, 1), ShouldEqual, "Big Mesa Road, Gunnison County, Colorado|Big Mesa Road, Gunnison County, Colorado")
		So(testMinimalCompare(1, 2), ShouldEqual, "Gunnison County, USA|Silauri, Georgia")
		So(testMinimalCompare(1, 3), ShouldEqual, "Gunnison County, USA|Paris, France")
		So(testMinimalCompare(1, 4), ShouldEqual, "Gunnison County, USA|Paris, France")
		So(testMinimalCompare(1, 5), ShouldEqual, "Gunnison County, USA|Paris, France")
		So(testMinimalCompare(1, 6), ShouldEqual, "Gunnison County, USA|Gratens, France")
		So(testMinimalCompare(1, 7), ShouldEqual, "Gunnison County, USA|Manosque, France")

		So(testMinimalCompare(2, 0), ShouldEqual, "Silauri, Georgia|Birmingham, UK")
		So(testMinimalCompare(2, 1), ShouldEqual, "Silauri, Georgia|Gunnison County, USA")
		So(testMinimalCompare(2, 2), ShouldEqual, "Silauri|Silauri")
		So(testMinimalCompare(2, 3), ShouldEqual, "Silauri, Georgia|Paris, France")
		So(testMinimalCompare(2, 4), ShouldEqual, "Silauri, Georgia|Paris, France")
		So(testMinimalCompare(2, 5), ShouldEqual, "Silauri, Georgia|Paris, France")
		So(testMinimalCompare(2, 6), ShouldEqual, "Silauri, Georgia|Gratens, France")
		So(testMinimalCompare(2, 7), ShouldEqual, "Silauri, Georgia|Manosque, France")

		So(testMinimalCompare(3, 0), ShouldEqual, "Paris, France|Birmingham, UK")
		So(testMinimalCompare(3, 1), ShouldEqual, "Paris, France|Gunnison County, USA")
		So(testMinimalCompare(3, 2), ShouldEqual, "Paris, France|Silauri, Georgia")
		So(testMinimalCompare(3, 3), ShouldEqual, "96 Boulevard Bessières, Paris|96 Boulevard Bessières, Paris")
		So(testMinimalCompare(3, 4), ShouldEqual, "96 Boulevard Bessières, Paris|Gare Saint-Lazare, Paris")
		So(testMinimalCompare(3, 5), ShouldEqual, "96 Boulevard Bessières, Paris|15 Rue Eugène Varlin, Paris")
		So(testMinimalCompare(3, 6), ShouldEqual, "96 Boulevard Bessières, Paris|D 73a, Gratens")
		So(testMinimalCompare(3, 7), ShouldEqual, "96 Boulevard Bessières, Paris|Avenue Maréchal de Lattre de Tassigny, Manosque")

		So(testMinimalCompare(4, 0), ShouldEqual, "Paris, France|Birmingham, UK")
		So(testMinimalCompare(4, 1), ShouldEqual, "Paris, France|Gunnison County, USA")
		So(testMinimalCompare(4, 2), ShouldEqual, "Paris, France|Silauri, Georgia")
	})
}
