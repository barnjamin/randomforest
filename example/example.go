package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"randomforest"
	"strconv"
)

var iris_url = "https://archive.ics.uci.edu/ml/machine-learning-databases/iris/iris.data"

func main() {

	data, labels := parseIris()

	rf := randomforest.New(100)

	log.Println("Fitting...")
	rf.Fit(data, labels)
	log.Println("Done!")

	log.Println("Determining accuracy...")
	correct := 0.0
	for idx, vals := range data {
		predicted := rf.Predict(vals)
		if predicted == labels[idx] {
			correct++
		}
	}

	// Cheating because we're using samples we trained on
	log.Printf("(Cheating) Accuracy of: %.2f", (correct/float64(len(data)))*100)
}

func parseIris() ([][]float64, []string) {

	resp, err := http.Get(iris_url)
	if err != nil {
		log.Fatalf("Failed to Get iris data: %+v", err)
	}
	defer resp.Body.Close()

	r := csv.NewReader(resp.Body)
	data := [][]float64{}
	labels := []string{}
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	for _, record := range records {
		labels = append(labels, record[len(record)-1])
		vals := []float64{}
		for _, val := range record[:len(record)-1] {
			fval, _ := strconv.ParseFloat(val, 64)
			vals = append(vals, fval)
		}
		data = append(data, vals)
	}
	return data, labels
}
