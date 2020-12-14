package septa

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"io/ioutil"
	"net/http"
)

type Coordinate struct {
	Latitude  float64
	Longitude float64
}
const septaURL string = "https://www3.septa.org/hackathon/TrainView/"

func GetData() {

	septaData, err := doGetRequest(septaURL)
	if err != nil {
		panic(err)
	}

	septaDataParsed, _ := gabs.ParseJSON(septaData)
	cities, _ := septaDataParsed.Path("trainno").Children()
	fmt.Println(cities)

}



func doGetRequest(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

