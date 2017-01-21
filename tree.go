package randomforest

type Tree struct {
	Forest       *RandomForest
	Root         *Node
	Data         [][]float64
	Labels       []int
	UniqueCounts []float64

	MaxDepth int

	labelToIdx map[string]int
	idxToLabel map[int]string
}

func NewTree(rf *RandomForest, data [][]float64, labels []string) *Tree {
	t := &Tree{
		Forest: rf,
		Data:   data,
	}

	li := map[string]int{}
	for _, label := range labels {
		li[label] = 0
	}

	var newIdx int
	il := map[int]string{}
	for label, _ := range li {
		li[label] = newIdx
		il[newIdx] = label
		newIdx++
	}

	l := []int{}
	for _, label := range labels {
		l = append(l, li[label])
	}

	t.labelToIdx = li
	t.idxToLabel = il
	t.Labels = l

	uniques := make([]float64, len(il))
	for _, label := range l {
		uniques[label] += 1.0
	}
	t.UniqueCounts = uniques

	indicies := []int{}
	for x := 0; x < len(data); x++ {
		indicies = append(indicies, x)
	}

	t.Root = NewNode(t, indicies, 1)

	return t
}

func (t *Tree) Predict(data []float64) string {
	return t.idxToLabel[t.Root.Traverse(data)]
}
