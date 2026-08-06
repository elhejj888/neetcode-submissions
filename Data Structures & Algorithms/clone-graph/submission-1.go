/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {

	seen := make(map[*Node] *Node)
	var dfs func(node *Node) *Node

	dfs = func(node *Node) *Node {
		if node == nil {
			return node
		}
		if cp, ok := seen[node]; ok {
				return cp
			}
		var copy Node
		copy.Val = node.Val
		seen[node] = &copy
		

		for _,v := range node.Neighbors {


			copy.Neighbors = append(copy.Neighbors,dfs(v))
		}

	return &copy
	}
	
	return dfs(node)
	

}
