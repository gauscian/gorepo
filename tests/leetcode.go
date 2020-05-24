/**
 * Definition for a binary tree node.

 */
package some

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type hybrid struct {
	c *TreeNode
	p *TreeNode
	g *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	queue := [][]TreeNode{{}}

	sum_ := 0

	for len(queue) > 0 {
		count := len(queue)
		for count > 0 {
			top, queue := queue[0], queue[1:]
			c_ := top.c
			p_ := top.p
			g_ := top.g

			count -= 1
			if c_.left != nil {
				append(queue, h{
					c: &c_.left,
					p: &c_,
					g: &p_,
				})
			}
			if c_.right != nil {
				append(queue, h{
					c: &c_.right,
					p: &c_,
					g: &p_,
				})
			}

			if g != nil && (g.val % 2 == 0) {
				sum_ += c_.val
			}
		}
	}
	return sum_
}
