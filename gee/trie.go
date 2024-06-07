package gee

import (
	"strings"
)

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否模糊匹配，part 含有 : 或 * 时为true
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, depth int) {
	// 如果已经处理完所有部分，设置当前节点的 pattern
	if len(parts) == depth {
		n.pattern = pattern
		return
	}

	part := parts[depth]        //获取当前部分
	child := n.matchChild(part) // 查找是否有匹配的节点
	if child == nil {           // 如果没有匹配的节点则创建一个新节点
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, depth+1) // 递归处理下一部分
}

func (n *node) search(parts []string, depth int) *node {
	// 如果已经处理完所有部分或遇到通配符节点（以 * 开头的部分）
	if len(parts) == depth || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil // 未找到匹配的完整路径
		}
		return n // 找到匹配的完整路径
	}

	part := parts[depth]              // 获取当前部分
	children := n.matchChildren(part) // 查找所有匹配的子节点

	// 递归搜索匹配的子节点
	for _, child := range children {
		result := child.search(parts, depth+1)
		if result != nil {
			return result // 找到匹配节点
		}
	}

	return nil // 未找到匹配节点
}
