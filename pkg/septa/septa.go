package septa

import (
	"encoding/json"
	"github.com/Jeffail/gabs"
	"io/ioutil"
	"net/http"
)

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

const septaURL string = "https://www3.septa.org/hackathon/TrainView/"

type TrainView []struct {
	Lat         string  `json:"lat"`
	Lon         string  `json:"lon"`
	Trainno     string  `json:"trainno"`
	Service     string  `json:"service"`
	Dest        string  `json:"dest"`
	Currentstop string  `json:"currentstop"`
	Nextstop    string  `json:"nextstop"`
	Line        string  `json:"line"`
	Consist     string  `json:"consist"`
	Heading     float64 `json:"heading"`
	Late        int     `json:"late"`
	SOURCE      string  `json:"SOURCE"`
	TRACK       string  `json:"TRACK"`
	TRACKCHANGE string  `json:"TRACK_CHANGE"`
}

func GetTrainno() []*gabs.Container {

	septaData, err := doGetRequest(septaURL)
	if err != nil {
		panic(err)
	}

	septaDataParsed, _ := gabs.ParseJSON(septaData)
	trainno, _ := septaDataParsed.Path("trainno").Children()
	///fmt.Println(trainno)
	return trainno

}

func GetTrainView() (TrainView, error) {

	var trainView TrainView

	septaData, err := doGetRequest(septaURL)
	if err != nil {
		return trainView, err
	}

	if err := json.Unmarshal([]byte(septaData), &trainView); err != nil {
		return trainView, err
	}

	return trainView, err

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
