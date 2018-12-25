package main

import (
	"fmt"
	"log"
)

// Node 二元搜尋樹的節點
type Node struct {
	left, right *Node
	bst         *BST //所屬的BST,為了更新BST結構中的變數num
	level       int  //該節點在BST的哪一層
	value       int
}

// Left 查詢該node的左節點
func (node Node) Left() *Node { return node.left }

// Right 查詢該node的右節點
func (node Node) Right() *Node { return node.right }

// Level 查詢該node的level
func (node Node) Level() int { return node.level }

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
	b.root = push(b.root, &Node{value: v, bst: b, level: 1})
}

func push(root, node *Node) *Node {
	if root == nil {
		root = node
		root.bst.num++

		if node.level > root.bst.hight {
			root.bst.hight++
		}
		return root
	}

	v := node.value
	if root.value > v {
		node.level++
		root.left = push(root.left, node)
	} else if root.value < v {
		node.level++
		root.right = push(root.right, node)
	} else {
		log.Fatalf("BST內有同樣數值%d,違反定義無法使用二元搜尋樹 \n", v)
	}
	return root
}

// Remove 移除BST內的特定節點
func (b *BST) Remove(want int) (ok bool) {
	node, ok := b.Search(want)
	if node != nil && !ok {
		log.Printf("數值%d,在BST 無法找到\n", want)
		return ok
	}

	if node = nil && !ok {
		log.Println("BST內無元素 無法移除元素")
		return ok
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
	fmt.Println("搜尋樹內的元素列表", bst.DataOfBST())

	// unsortData := []int{14, 34, 87, 35, 23, 1, 82}
	// fmt.Println("未排序", unsortData)
	// fmt.Println("已排序", Sort(unsortData))

	if node, ok := bst.Search(30); ok {
		fmt.Printf("尋找到該節點 Value=%d Level=%d Pointer=%p\n", node.Value(), node.Level(), node)
	} else {
		fmt.Println("找不到該節點")
	}
	fmt.Println(bst.root)
}
