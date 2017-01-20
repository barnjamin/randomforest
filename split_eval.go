package randomforest

type SplitEvaluator func([]string, []string, []string) float64

func Gini(left, right, all []string) float64 {
	var score, proportion float64

	for _, label := range all {
		leftcnt, rightcnt := 0, 0
		for _, left_label := range left {
			if label == left_label {
				leftcnt++
			}
		}
		proportion = float64(leftcnt) / float64(len(left))
		score += proportion * (1 - proportion)

		for _, right_label := range right {
			if label == right_label {
				rightcnt++
			}
		}
		proportion = float64(rightcnt) / float64(len(right))
		score += proportion * (1 - proportion)

	}
	return score
}
