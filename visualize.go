package randomforest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"

	svg "github.com/ajstarks/svgo"
)

var (
	width         int
	height        int
	radius        = 200
	level_spacing = radius * 2
	filename      = "tree.svg"

	circleStyle = "fill:white;stroke:black"
	textStyle   = "text-anchor:middle;font-size:50px"
	lineStyle   = "stroke-width:2;stroke:black"
)

func Visualize(tree *Tree) error {

	var buff bytes.Buffer

	maxLeaves := int(math.Pow(2, float64(tree.MaxDepth)))

	width = (radius + 5) * maxLeaves
	height = ((radius * 3) * tree.MaxDepth)

	canvas := svg.New(&buff)
	canvas.Start(width, height)

	//Add Root
	right := width / 2
	top := (level_spacing / 2)

	canvas.Circle(right, top, radius, circleStyle)

	canvas.Text(right, top, nodeText(tree.Root), textStyle)

	//Recurse through left and right branches
	if tree.Root.Left != nil {
		addToCanvas(canvas, tree.Root.Left, 1, right, -1)
	}

	if tree.Root.Right != nil {
		addToCanvas(canvas, tree.Root.Right, 1, right, 1)
	}

	canvas.End()

	ioutil.WriteFile(filename, buff.Bytes(), 0666)

	return nil
}

func addToCanvas(canvas *svg.SVG, node *Node, level, right, direction int) {

	blockSize := (width / int(math.Pow(2, float64(level))))
	newRight := right + (direction * (blockSize / 2))
	top := (level_spacing * level)

	canvas.Line(right, level_spacing*(level-1), newRight, top, lineStyle)

	canvas.Circle(newRight, top, radius, circleStyle)

	canvas.Text(newRight, top, nodeText(node), textStyle)

	if node.Left != nil {
		addToCanvas(canvas, node.Left, level+1, newRight, -1)
	}
	if node.Right != nil {
		addToCanvas(canvas, node.Right, level+1, newRight, 1)
	}
}

func nodeText(n *Node) string {
	return fmt.Sprintf("L:%d S:%d F:%d V:%.2f", n.Label, n.IdxCnt, n.FeatureIndex, n.Value)
}
