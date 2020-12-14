package influx

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/zoechirico/septa/pkg/config"
	"github.com/zoechirico/septa/pkg/septa"
	"os"
)

func Write() {

	client := Client()
	defer client.Close()

	bucket := os.Getenv("SEPTA_BUCKET")
	org := os.Getenv("SEPTA_ORG")

	writeAPI := client.WriteAPI(org, bucket)

	values := septa.GetTrainno()
	for idx, v := range values {
		writeAPI.WriteRecord(fmt.Sprintf("septa,unit=train trainno=%v", v))
		fmt.Println(idx, v)
	}

	// Flush writes
	writeAPI.Flush()
}

func Client() influxdb2.Client {
	config.SetEnv()
	token := os.Getenv("SEPTA_TOKEN")
	url := os.Getenv("SEPTA_URL")

	client := influxdb2.NewClient(url, token)
	return client

}

func Read() {

	client := Client()
	defer client.Close()

	bucket := os.Getenv("SEPTA_BUCKET")
	org := os.Getenv("SEPTA_ORG")

	query := fmt.Sprintf("from(bucket:\"%v\")|> range(start: -1h) |> filter(fn: (r) => r._measurement == \"septa\")", bucket)
	// Get query client
	queryAPI := client.QueryAPI(org)
	// get QueryTableResult
	result, err := queryAPI.Query(context.Background(), query)
	if err == nil {
		// Iterate over query response
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// Access data
			fmt.Printf("value: %v\n", result.Record().Value())
		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %v\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
}
