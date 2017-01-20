package randomforest

import "log"

func Example() {

	data := [][]float64{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
		{4, 4, 4},
	}
	labels := []string{"ben", "cassie", "floki", "luna"}

	rf := New(100)

	// Define your own split evaluator
	// Return a score of how good this split is
	// Default is Gini
	//
	// rf.SplitEvaluator = func(left, right, all []string) float64 {
	// 	return rand.Float64()
	// }

	rf.Fit(data, labels)

	for idx, vals := range data {
		prediction := rf.Predict(vals)
		log.Printf("Wanted: %s Got: %s", labels[idx], prediction)
	}
}
