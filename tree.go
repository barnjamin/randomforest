package randomforest

type Tree struct {
	Forest *RandomForest
	Root   *Node
	Data   [][]float64
	Labels []string
}

func NewTree(rf *RandomForest, data [][]float64, labels []string) *Tree {
	t := &Tree{
		Forest: rf,
		Data:   data,
		Labels: labels,
	}

	indicies := []int{}
	for x := 0; x < len(data); x++ {
		indicies = append(indicies, x)
	}

	t.Root = NewNode(t, indicies, 1)

	return t
}

func (t *Tree) Predict(data []float64) string {
	return t.Root.Traverse(data)
}
