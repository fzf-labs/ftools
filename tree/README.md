## 树形结构数据封装

```go
package tree

type Tree struct {
	Id    int         `json:"id"`
	Pid   int         `json:"pid"`
	Data  interface{} `json:"data"`
	Child []Tree      `json:"child"`
}

type Trees []Tree

func (t Trees) Len() int {
	return len(t)
}

func (t Trees) Less(i, j int) bool {
	return t[i].Id < t[j].Id
}

func (t Trees) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func GenerateTree(list []Tree) []Tree {
	var trees []Tree
	// Define the top-level root and child nodes
	var roots, childs []Tree
	for _, v := range list {
		if v.Pid == 0 {
			// Determine the top-level root node
			roots = append(roots, v)
		}
		childs = append(childs, v)
	}

	for _, v := range roots {
		childTree := &Tree{
			Id:    v.Id,
			Pid:   v.Pid,
			Data:  v.Data,
			Child: make([]Tree, 0),
		}
		// recursive
		recursiveTree(childTree, childs)

		trees = append(trees, *childTree)
	}
	return trees
}

func GenerateRootTree(list []Tree, id int) Tree {
	var tree Tree
	trees := GenerateTree(list)
	for _, v := range trees {
		if v.Id == id {
			tree = v
			break
		}
	}
	return tree
}

func recursiveTree(tree *Tree, allNodes []Tree) {
	for _, v := range allNodes {
		if v.Pid == 0 {
			// If the current node is the top-level root node, skip
			continue
		}
		if tree.Id == v.Pid {
			childTree := &Tree{
				Id:    v.Id,
				Pid:   v.Pid,
				Data:  v.Data,
				Child: make([]Tree, 0),
			}
			recursiveTree(childTree, allNodes)
			tree.Child = append(tree.Child, *childTree)
		}
	}
}

// FindSubNode 查询子级
func FindSubNode(node *Tree, allNodes []Tree) {
	for _, v := range allNodes {
		if node.Id == v.Pid {
			FindSubNode(&v, allNodes)
			node.Child = append(node.Child, v)
		}
	}
}

// FindParentNode 查询父级
func FindParentNode(node *Tree, allNodes []Tree) {
	for _, v := range allNodes {
		temp := v
		if node.Pid == temp.Id {
			temp.Child = append(temp.Child, *node)
			*node = temp
			FindParentNode(node, allNodes)
		}
	}
}

func TreeRange(trees []Tree) {
	for _, v := range trees {
		if len(v.Child) > 0 {
			TreeRange(v.Child)

		}
	}
}

```