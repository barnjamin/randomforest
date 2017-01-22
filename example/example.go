package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"strconv"

	"github.com/barnjamin/bayes"
	"github.com/barnjamin/randomforest"
)

var iris_url = "https://archive.ics.uci.edu/ml/machine-learning-databases/iris/iris.data"

func main() {

	data, labels := parseIris()
	trainData, trainLabels, testData, testLabels := bayes.Split(data, labels, 0.3)

	log.Printf("%d %d %d %d", len(trainData), len(trainLabels), len(testData), len(testLabels))
	rf := randomforest.New(100)

	log.Println("Fitting...")
	rf.Fit(trainData, trainLabels)
	log.Println("Done!")

	log.Println("Determining accuracy...")
	correct := 0.0
	for idx, vals := range testData {
		predicted := rf.Predict(vals)
		if predicted == testLabels[idx] {
			correct++
		}
	}

	// Cheating because we're using samples we trained on
	log.Printf("Accuracy of: %.2f", (correct/float64(len(testData)))*100)
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
