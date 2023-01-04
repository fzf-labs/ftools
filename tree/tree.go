package tree

import (
	"sort"
)

// Tree 统一定义菜单树的数据结构，也可以自定义添加其他字段
type Tree struct {
	Data     interface{} `json:"data"`     //自定义对象
	Children []Tree      `json:"children"` //子节点
}

// INode 其他的结构体想要生成菜单树，直接实现这个接口
type INode interface {
	// GetId 获取id
	GetId() int
	// GetPid 获取父id
	GetPid() int
	// GetData 获取附加数据
	GetData() interface{}
	// IsRoot 判断当前节点是否是顶层根节点
	IsRoot() bool
}
type INodes []INode

func (nodes INodes) Len() int {
	return len(nodes)
}
func (nodes INodes) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}
func (nodes INodes) Less(i, j int) bool {
	return nodes[i].GetId() < nodes[j].GetId()
}

// GenerateTree 自定义的结构体实现 INode 接口后调用此方法生成树结构
// nodes 需要生成树的节点
// selectedNode 生成树后选中的节点
// menuTrees 生成成功后的树结构对象
func GenerateTree(nodes []INode) (trees []Tree) {
	trees = []Tree{}
	// 定义顶层根和子节点
	var roots, childs []INode
	for _, v := range nodes {
		if v.IsRoot() {
			// 判断顶层根节点
			roots = append(roots, v)
		}
		childs = append(childs, v)
	}

	for _, v := range roots {
		childTree := &Tree{
			Data: v.GetData(),
		}
		// 递归
		recursiveTree(childTree, childs)
		trees = append(trees, *childTree)
	}
	return
}

// recursiveTree 递归生成树结构
// tree 递归的树对象
// nodes 递归的节点
// selectedNodes 选中的节点
func recursiveTree(tree *Tree, nodes []INode) {
	data := tree.Data.(INode)

	for _, v := range nodes {
		if v.IsRoot() {
			// 如果当前节点是顶层根节点就跳过
			continue
		}
		if data.GetId() == v.GetPid() {
			childTree := &Tree{
				Data: v.GetData(),
			}
			recursiveTree(childTree, nodes)
			tree.Children = append(tree.Children, *childTree)
		}
	}
}

// FindRelationSubNode 查询子节点
func FindRelationSubNode(allNodes []INode, node INode) (results []INode) {
	results = make([]INode, 0)
	results = recursiveFindRelationNode(results, allNodes, node, 2)
	sort.Sort(INodes(results))
	return results
}

// recursiveFindRelationNode 递归查询关联父子节点
// nodeMap 查询结果搜集到map中
// allNodes 所有节点
// node 递归节点
// t 递归查找类型：0 查找父、子节点；1 只查找父节点；2 只查找子节点
func recursiveFindRelationNode(results []INode, allNodes []INode, node INode, t int) []INode {
	results = append(results, node)
	for _, v := range allNodes {
		if v.GetId() == node.GetId() {
			continue
		}
		// 查找父节点
		if t == 0 || t == 1 {
			if node.GetPid() == v.GetId() {
				results = append(results, v)
				if v.IsRoot() {
					// 是顶层根节点时，不再进行递归
					continue
				}
				recursiveFindRelationNode(results, allNodes, v, 1)
			}
		}
		// 查找子节点
		if t == 0 || t == 2 {
			if node.GetId() == v.GetPid() {
				results = append(results, v)
				recursiveFindRelationNode(results, allNodes, v, 2)
			}
		}
	}
	return results
}
