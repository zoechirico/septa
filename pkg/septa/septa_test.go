package septa

import (
	"fmt"
	"testing"
)

func TestGetData(t *testing.T) {
	values := GetTrainno()
	for idx, v := range values {
		fmt.Println(idx, v)
	}
}

func TestGetT(t *testing.T) {
	r, e := GetTrainView()
	if e != nil {
		t.Fatalf("Error: %v\n", e)
	}
	_ = r
}
