package main

import (
	"fmt"
	"log"
)

// Node 二元搜尋樹的節點
type Node struct {
	parent      *Node
	left, right *Node
	bst         *BST //所屬的BST,為了更新BST結構中的變數num
	depth       int  //該節點在BST的哪一層
	value       int
}

// Left 查詢該node的左節點
func (node Node) Left() *Node { return node.left }

// Right 查詢該node的右節點
func (node Node) Right() *Node { return node.right }

// Depth 查詢該node的Depth
func (node Node) Depth() int { return node.depth }

// Value 查詢該node的value
func (node Node) Value() int { return node.value }

// BST 二元搜尋樹結構
type BST struct {
	root  *Node
	num   int
	hight int
}

// New 產生一個新的二元搜尋樹
func New() *BST { return new(BST) }

// Len 回傳搜尋樹中的節點個數
func (b *BST) Len() int { return b.num }

// Hight 回傳搜尋樹的高度
func (b *BST) Hight() int { return b.hight }

// Push 將資料插入二元搜尋樹的結構
func (b *BST) Push(v int) {
	b.root = push(b.root, &Node{value: v, bst: b, depth: 1})
}

func push(root, node *Node) *Node {
	if root == nil {
		root = node
		root.bst.num++

		if node.depth > root.bst.hight {
			root.bst.hight++
		}
		return root
	}

	v := node.value
	switch {
	case root.value > v:
		node.depth++
		node.parent = root
		root.left = push(root.left, node)
	case root.value < v:
		node.depth++
		node.parent = root
		root.right = push(root.right, node)
	default:
		log.Fatalf("BST內有同樣數值%d,違反定義無法使用二元搜尋樹 \n", v)

	}
	return root
}

// Remove 移除BST內的特定節點
func (b *BST) Remove(want int) {
	node, ok := b.Search(want)
	if node != nil && !ok {
		log.Printf("數值%d,在BST 無法找到\n", want)
		return 
	}

	if node == nil && !ok {
		log.Println("BST內無元素 無法移除元素")
		return 
	}

	// fmt.Printf("%p %#v \n",node,node)
	// fmt.Printf("%p %#v \n",node.parent,node.parent)

	// 根據刪除的節點狀態,有四種情況
	switch  {
	case (node.Left() == nil) && (node.Right() == nil) :
		if node.parent.left == node {
			node.parent.left = nil
		} else {
			node.parent.right = nil
		}
	case node.Left() == nil:
		if node.parent.left == node {
			node.parent.left = node.right
		} else {
			node.parent.right = node.right
		}
		node.right.parent = node.parent
		node.right = nil
	case node.Right() == nil:
		if node.parent.left == node {
			node.parent.left = node.left
		} else {
			node.parent.right = node.left
		}
		node.left.parent = node.parent
		node.left = nil
	default:
		maxNode:=max(node)

		if node.parent.left == node {
			node.parent.left = maxNode
		} else {
			node.parent.right = maxNode
		}

		node.left.parent = maxNode.left
		minNode:=min(maxNode)
		minNode.left = node.left
		maxNode.parent = node.parent
		node.right = nil
		node.left = nil
	}
	node.bst.num--
	node.parent = nil
	node.bst = nil
	return 
}

// min 以該節點為子樹,尋找子樹中的最小值節點
func min(tree *Node) *Node{
	for {
        if tree.left == nil {
            return tree
        }
        tree = tree.left
    }
}

// max 以該節點為子樹,尋找子樹中的最大值節點
func max(tree *Node) *Node{
	for {
        if tree.right == nil {
            return tree
        }
        tree = tree.right
    }
}

// DataOfBST 回傳二元搜尋樹內的數值,以中序走訪的方式排列
func (b *BST) DataOfBST() []int {
	nodeSlice := make([]Node, 0, b.num)
	nodeSlice = inOrder(b.root, nodeSlice)

	valueSlice := make([]int, 0, b.num)
	for _, n := range nodeSlice {
		valueSlice = append(valueSlice, n.value)
	}
	return valueSlice
}

func inOrder(node *Node, slice []Node) []Node {
	if node != nil {
		slice = inOrder(node.left, slice)
		slice = append(slice, *node)
		slice = inOrder(node.right, slice)
	}
	return slice
}

// Search 尋找搜尋樹中是否存在特定元素(數值)
// receiver使用value type,避免迴圈運作時,更改實體物件
func (b BST) Search(want int) (v *Node, ok bool) {
	if b.root == nil {
		//log.Println("BST內無元素,無法查詢")
		return b.root, false
	}

	//用迴圈替換遞迴
	for b.root != nil {

		//成立條件的code要放在迴圈最一開始,才可以判斷是否繼續
		if b.root.value == want {
			//log.Printf("數值%d,在BST 可以找到\n", want)
			return b.root, true
		}

		if b.root.value > want {
			b.root = b.root.left
		} else if b.root.value < want {
			b.root = b.root.right
		}
	}

	//log.Printf("數值%d,在BST 無法找到\n", want)
	return b.root, false
}

// Sort 輸入未排序數列,可回傳排序後的數列
func Sort(values []int) []int {
	bst := &BST{}
	for _, v := range values {
		bst.Push(v)
	}
	values = bst.DataOfBST()
	return values
}

func main() {
	bst := New()

	bst.Push(25)
	bst.Push(15)
	bst.Push(30)
	bst.Push(12)
	bst.Push(17)
	bst.Push(27)
	bst.Push(16)
	bst.Push(39)

	// fmt.Println("搜尋樹的節點個數=", bst.Len())
	// fmt.Println("搜尋樹的高度=", bst.Hight())
	fmt.Println("新增後,搜尋樹內的元素列表", bst.DataOfBST())

	// unsortData := []int{14, 34, 87, 35, 23, 1, 82}
	// fmt.Println("未排序", unsortData)
	// fmt.Println("已排序", Sort(unsortData))

	// if node, ok := bst.Search(17); ok {
	// 	fmt.Printf("尋找到該節點 Value=%d Depth=%d Pointer=%p Parent=%p\n", node.Value(), node.Depth(), node,node.parent)
	// } else {
	// 	fmt.Println("找不到該節點")
	// }
	// fmt.Printf("%p %p\n",bst.root.left.right,bst.root.left)

	bst.Remove(30)
	fmt.Println("移除後,搜尋樹內的元素列表", bst.DataOfBST())
	// if node, ok := bst.Search(30); ok {
	// 	fmt.Printf("尋找到該節點 Value=%d Depth=%d Pointer=%p Parent=%p\n", node.Value(), node.Depth(), node, node.parent)
	// } else {
	// 	fmt.Println("找不到該節點")
	// }
	// fmt.Printf("%v %p %p\n", bst.root, bst.root, bst.root.parent)

	for _,v:=range bst.DataOfBST(){
		node,_:=bst.Search(v)
		fmt.Printf("%p %#v \n",node,node)
		fmt.Printf("%p %#v \n",node.parent,node.parent)
		fmt.Println()
	}
}
