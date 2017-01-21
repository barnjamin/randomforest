package randomforest

import "math/rand"

type Node struct {
	Tree *Tree

	FeatureIndex int
	Value        float64

	IdxCnt int

	Left  *Node
	Right *Node

	Label int
}

func NewNode(tree *Tree, indicies []int, level int) *Node {
	node := &Node{
		Tree:   tree,
		Label:  getMaxLabel(tree, indicies),
		IdxCnt: len(indicies),
	}

	if tree.MaxDepth < level {
		tree.MaxDepth = level
	}

	if level >= tree.Forest.MaxDepth || len(indicies) <= tree.Forest.MinSize {
		return node
	}

	var (
		maxScore              float64
		best_left, best_right []int
	)

	featureSlice := randomSlice(len(tree.Data[0]), tree.Forest.FeatureCount)
	for _, idx := range featureSlice {
		for _, row := range tree.Data {

			left, right := split(idx, row[idx], indicies, tree.Data, tree.Labels)

			score := tree.Forest.Evaluator(tree, left, right)

			if score > maxScore {
				maxScore = score
				best_left, best_right = left, right

				node.FeatureIndex = idx
				node.Value = row[idx]
			}
		}
	}

	if len(best_left) > 0 {
		node.Left = NewNode(tree, best_left, level+1)
	}

	if len(best_right) > 0 {
		node.Right = NewNode(tree, best_right, level+1)
	}

	return node

}

func (n *Node) Traverse(data []float64) int {
	if data[n.FeatureIndex] < n.Value {
		if n.Left != nil {
			return n.Left.Traverse(data)
		}
		return n.Label
	} else if n.Right != nil {
		if n.Right != nil {
			return n.Right.Traverse(data)
		}
		return n.Label
	}

	return n.Label
}

func getMaxLabel(tree *Tree, indicies []int) int {

	labels := map[int]int{}
	for _, idx := range indicies {
		labels[tree.Labels[idx]] += 1
	}

	var maxCnt int
	var maxLabel int
	for label, cnt := range labels {
		if cnt > maxCnt {
			maxCnt = cnt
			maxLabel = label
		}
	}

	return maxLabel
}

func split(index int, splitVal float64, indicies []int, data [][]float64, labels []int) ([]int, []int) {
	left, right := []int{}, []int{}
	for _, idx := range indicies {
		if data[idx][index] < splitVal {
			left = append(left, labels[idx])
		} else {
			right = append(right, labels[idx])
		}
	}
	return left, right
}

func randomSlice(maxVal, size int) []int {
	idxMap := map[int]bool{}
	indicies := []int{}
	for len(indicies) < size {
		idx := rand.Intn(maxVal)
		if _, ok := idxMap[idx]; !ok {
			idxMap[idx] = true
			indicies = append(indicies, idx)
		}
	}
	return indicies
}
