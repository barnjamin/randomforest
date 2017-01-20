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
	rf.Fit(data, labels)

	for idx, vals := range data {
		prediction := rf.Predict(vals)
		log.Printf("Wanted: %s Got: %s", labels[idx], prediction)
	}
}
