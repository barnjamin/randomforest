package randomforest

import (
	"log"
	"testing"
)

type TestData struct {
	Data  []float64
	Label string
}

var data = []TestData{
	{[]float64{1, 1, 1}, "ben"},
	{[]float64{2, 2, 2}, "cassie"},
	{[]float64{3, 3, 3}, "floki"},
	{[]float64{4, 4, 4}, "luna"},
}

func TestRandomForest(t *testing.T) {

	rf := New(100)

	datas := [][]float64{}
	labels := []string{}
	for _, val := range data {
		datas = append(datas, val.Data)
		labels = append(labels, val.Label)
	}
	rf.Fit(datas, labels)
	log.Printf("%s", rf.Predict([]float64{3, 3, 3}))
}
