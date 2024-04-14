package dlx

type Column struct {
	*Node
	ID   string
	Size int
}

func (c *Column) Cover() {
	c.Left.Right = c.Right
	c.Right.Left = c.Left

	for node := c.Down; node != c.Node; node = node.Down {
		for rowNode := node.Right; rowNode != node; rowNode = rowNode.Right {
			rowNode.Up.Down = rowNode.Down
			rowNode.Down.Up = rowNode.Up
			rowNode.Column.Size--
		}
	}
}

func (c *Column) Uncover() {
	for node := c.Up; node != c.Node; node = node.Up {
		for rowNode := node.Left; rowNode != node; rowNode = rowNode.Left {
			rowNode.Up.Down = rowNode
			rowNode.Down.Up = rowNode
			rowNode.Column.Size++
		}
	}

	c.Left.Right = c.Node
	c.Right.Left = c.Node
}
