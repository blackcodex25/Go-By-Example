package main

import (
	"fmt"

	"github.com/labstack/gommon/log"
)

// โครงสร้างของโหนดในต้นไม้ไบนารี
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// ฟังก์ชันการเดินทาง In-order
func inOrderTraversal(node *TreeNode) {
	if node != nil {
		inOrderTraversal(node.Left)
		fmt.Print(node.Value, " ")
		inOrderTraversal(node.Right)
	}
}

// ฟังก์ชันการค้นหาโหนด
func search(node *TreeNode, target int) (*TreeNode, error) {
	if node == nil {
		return nil, &TreeError{Message: fmt.Sprintf("Node with value %d not found", target)}
	}
	if node.Value == target {
		return node, nil
	}
	if target < node.Value {
		return search(node.Left, target)
	}
	return search(node.Right, target)
}

// Insert adds a value to the binary search tree.
func insert(node *TreeNode, value int) *TreeNode {
	// If the current node is nil, create a new node.
	if node == nil {
		return &TreeNode{Value: value}
	}

	// Traverse the tree to find the correct position for the new value.
	for {
		if value < node.Value {
			if node.Left == nil {
				node.Left = &TreeNode{Value: value}
				break
			}
			node = node.Left
		} else {
			if node.Right == nil {
				node.Right = &TreeNode{Value: value}
				break
			}
			node = node.Right
		}
	}

	return node
}

type TreeError struct {
	Message string
}

func (e *TreeError) Error() string {
	return e.Message
}

// ตรวจสอบการค้นหาโหนด หากค้นหาโหนดไม่พบจะเกิดข้อผิดพลาด แต่ถ้าพบจะแสดงผลลัพธ์// ตรวจสอบการค้นหาโหนด
func handleSearchResult(result *TreeNode, err error, target int) {
	if err != nil {
		log.Errorf("Error: %v for target: %d", target, err)
		return
	}
	fmt.Printf("Found node with value: %d\n", result.Value)
}

// ฟังก์ชันหลัก
func main() {
	// สร้างต้นไม้ไบนารี
	root := &TreeNode{Value: 5}
	root.Left = &TreeNode{Value: 3}
	root.Right = &TreeNode{Value: 8}
	root.Left.Left = &TreeNode{Value: 1}
	root.Left.Right = &TreeNode{Value: 4}
	root.Right.Right = &TreeNode{Value: 9}

	// การเดินทาง In-order
	fmt.Println("In-order Traversal:")
	inOrderTraversal(root) // Output: 1 3 4 5 8 9
	fmt.Println()

	// การค้นหาโหนด
	target := 4
	result, err := search(root, target)
	handleSearchResult(result, err, target)
}

/* TreeNode: โครงสร้างของโหนดที่มีค่าและลูกซ้ายและขวา
inOrderTraversal: ฟังก์ชันสำหรับเดินทางต้นไม้ในลำดับ In-order (ซ้าย, ราก, ขวา)
search: ฟังก์ชันสำหรับค้นหาโหนดที่มีค่าตามที่กำหนด
main: สร้างต้นไม้ตัวอย่างและเรียกใช้ฟังก์ชันเดินทางและค้นหา
*/
