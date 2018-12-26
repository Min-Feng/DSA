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
	value       int
}

// Left 查詢該node的左節點
func (node Node) Left() *Node { return node.left }

// Right 查詢該node的右節點
func (node Node) Right() *Node { return node.right }

// Value 查詢該node的value
func (node Node) Value() int { return node.value }

// BST 二元搜尋樹結構
type BST struct {
	root *Node
	num  int
}

// New 產生一個新的二元搜尋樹
func New() *BST { return new(BST) }

// Len 回傳搜尋樹中的節點個數
func (b *BST) Len() int { return b.num }

// Push 將資料插入二元搜尋樹的結構
func (b *BST) Push(v int) {
	b.root = push(b.root, &Node{value: v, bst: b})
}

func push(root, node *Node) *Node {
	if root == nil {
		root = node
		root.bst.num++
		return root
	}

	v := node.value
	switch {
	case root.value > v:
		node.parent = root
		root.left = push(root.left, node)
	case root.value < v:
		node.parent = root
		root.right = push(root.right, node)
	default:
		log.Fatalf("BST內有同樣數值%d,違反定義無法使用二元搜尋樹 \n", v)

	}
	return root
}

// Remove 移除BST內的特定節點
func (b *BST) Remove(want int) {
	oldNode, ok := b.Search(want)
	if oldNode != nil && !ok {
		log.Printf("數值%d,在BST 無法找到\n", want)
		return
	}

	if oldNode == nil && !ok {
		log.Println("BST內無元素 無法移除元素")
		return
	}

	// fmt.Printf("%p %#v \n",oldNode,oldNode)
	// fmt.Printf("%p %#v \n",oldNode.parent,oldNode.parent)

	// 刪除的節點,依據其子節點狀態,有四種情況
	var newNode *Node
	switch {
	case (oldNode.Left() == nil) && (oldNode.Right() == nil):
		newNode = nil
		updateChildOfParentOfOldNode(oldNode, newNode)
	case oldNode.Left() == nil:
		newNode = oldNode.right
		updateChildOfParentOfOldNode(oldNode, newNode)
		newNode.parent = oldNode.parent
		oldNode.right = nil // avoid memory leaks
	case oldNode.Right() == nil:
		newNode = oldNode.left
		updateChildOfParentOfOldNode(oldNode, newNode)
		newNode.parent = oldNode.parent
		oldNode.left = nil // avoid memory leaks
	default:
		// 新節點的選擇為 左子樹中的最大節點or右子樹中的最小節點 在此選前者
		newNode = max(oldNode.left)

		// 更新 舊節點其父節點訊息
		updateChildOfParentOfOldNode(oldNode, newNode)

		// 更新 新節點 訊息
		if newNode != oldNode.left {
			// 若新節點不為舊節點的左節點,則肯定新節點存在於(舊節點其左節點)的右子樹中
			min(newNode).left = oldNode.left

			// 更新 新節點其父節點訊息
			updateChildOfParentOfOldNode(newNode, nil)
		}
		newNode.parent = oldNode.parent
		newNode.right = oldNode.right

		// 更新 舊節點其子節點訊息
		oldNode.right.parent = newNode
		if min(newNode) != oldNode.left {
			oldNode.left.parent = min(newNode)
		}

		// 釋放舊節點
		oldNode.right = nil // avoid memory leaks
		oldNode.left = nil  // avoid memory leaks
	}
	oldNode.bst.num--
	oldNode.parent = nil // avoid memory leaks
	oldNode.bst = nil    // avoid memory leaks
	return
}

// updateChildOfParentOfOldNode 移動或刪除二元樹的節點(oldNode)
// 更新其父節點(parent)的子節點資訊,使其連接到新節點(newNode)
func updateChildOfParentOfOldNode(oldNode, newNode *Node) {
	// if old node is BST root node !
	if oldNode.parent == nil {
		newNode.bst.root = newNode
		return
	}

	if oldNode.parent.left == oldNode {
		oldNode.parent.left = newNode
	} else {
		oldNode.parent.right = newNode
	}
}

// min 以該節點為子樹,尋找子樹中的最小值節點
func min(tree *Node) *Node {
	for {
		if tree.left == nil {
			return tree
		}
		tree = tree.left
	}
}

// max 以該節點為子樹,尋找子樹中的最大值節點
func max(tree *Node) *Node {
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
	// fmt.Printf("%p %#v \n", node, node)
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
	bst.Push(45)
	bst.Push(40)
	bst.Push(65)
	bst.Push(32)
	bst.Push(43)
	bst.Push(41)
	bst.Push(42)
	bst.Push(50)
	fmt.Printf("新增後,節點個數 = %d \t 搜尋樹內的元素列表 = %v\n", bst.Len(), bst.DataOfBST())

	// unsortData := []int{14, 34, 87, 35, 23, 1, 82}
	// fmt.Println("未排序", unsortData)
	// fmt.Println("已排序", Sort(unsortData))

	// bst.Remove(15)
	// bst.Remove(17)
	// bst.Remove(30)
	// bst.Remove(39)
	bst.Remove(45)
	bst.Remove(41)
	fmt.Printf("新增後,節點個數 = %d \t 搜尋樹內的元素列表 = %v\n", bst.Len(), bst.DataOfBST())

	// if node, ok := bst.Search(16); ok {
	// 	fmt.Printf("尋找到該節點 %p %#v \n", node, node)
	// } else {
	// 	fmt.Println("找不到該節點")
	// }

	// for _, v := range bst.DataOfBST() {
	// 	node, _ := bst.Search(v)
	// 	fmt.Printf("%p %#v \n", node, node)
	// 	fmt.Printf("Pointer of parent = %p \n", node.parent)
	// 	fmt.Println()
	// }
}
