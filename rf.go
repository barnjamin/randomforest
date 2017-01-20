package randomforest

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type RandomForest struct {
	TreeCount    int            // Number of trees to include in forest
	MaxDepth     int            //
	MinSize      int            //
	FeatureCount int            // Number of features to include in each tree
	Trees        []*Tree        // List of trees created in `Fit`
	Evaluator    SplitEvaluator // Function to evaluate each split
	SampleSize   float64        // Ratio of rows to consider in each tree

	Data   [][]float64
	Labels []string
}

func New(treeCnt int) *RandomForest {
	return &RandomForest{
		TreeCount:  treeCnt,
		Evaluator:  Gini,
		MaxDepth:   10,
		MinSize:    1,
		SampleSize: 1.0,
	}
}

func (r *RandomForest) Fit(data [][]float64, labels []string) error {
	if len(data) == 0 {
		return errors.New("No data provided")
	}

	if len(data) != len(labels) {
		return fmt.Errorf("Data and labels dont match %d vs %d", len(data), len(labels))
	}

	r.Data = data
	r.Labels = labels

	//If not set, set to default
	if r.FeatureCount == 0 {
		r.FeatureCount = int(math.Sqrt(float64(len(data[0]))))
	}

	for x := 0; x < r.TreeCount; x++ {
		d, l := r.Sample()
		tree := NewTree(r, d, l)
		r.Trees = append(r.Trees, tree)
	}

	return nil
}

func (r *RandomForest) Predict(data []float64) string {
	votes := map[string]int{}
	for _, tree := range r.Trees {
		votes[tree.Predict(data)] += 1
	}

	var prediction string
	var maxCount int
	for vote, cnt := range votes {
		if cnt > maxCount {
			maxCount = cnt
			prediction = vote
		}
	}
	return prediction
}

func (r *RandomForest) Sample() ([][]float64, []string) {
	sampleData := [][]float64{}
	sampleLabels := []string{}

	sampleCnt := int(float64(len(r.Data)) * r.SampleSize)
	for x := 0; x < sampleCnt; x++ {
		idx := rand.Intn(len(r.Data))

		sampleData = append(sampleData, r.Data[idx])
		sampleLabels = append(sampleLabels, r.Labels[idx])
	}

	return sampleData, sampleLabels
}
