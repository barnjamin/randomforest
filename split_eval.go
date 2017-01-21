package randomforest

type SplitEvaluator func(*Tree, []int, []int) float64

func Gini(t *Tree, left, right []int) float64 {
	var score, proportion float64

	for label, cnt := range t.UniqueCounts {
		leftcnt, rightcnt := 0, 0

		for _, idx := range left {
			if t.Labels[idx] == label {
				leftcnt++
			}
		}
		proportion = float64(leftcnt) / float64(len(left))
		score += (proportion * (1 - proportion)) * cnt

		for _, idx := range right {
			if t.Labels[idx] == label {
				rightcnt++
			}
		}
		proportion = float64(rightcnt) / float64(len(right))
		score += (proportion * (1 - proportion)) * cnt

	}

	return score
}
