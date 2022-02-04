package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	binaryArray := []int{10, 5, 15, 3, 7, 10, 18}
	treeArray := createBinaryTreeFunction(binaryArray)
	for _, treeArray := range treeArray {
		fmt.Print(fmt.Sprintf("%v ", treeArray.Val))
	}

}

func createBinaryTreeFunction(binaryArray []int) []*TreeNode {
	// left side 	2*i +1 left
	// right side   right 2(i+1)
	// parent floor ((i-1)/2)
	//test

	if len(binaryArray) == 0 || binaryArray == nil {
		return nil
	}
	rootTree := &TreeNode{Val: binaryArray[0]}
	treeArray := make([]*TreeNode, 0)
	treeArray = append(treeArray, rootTree)
	for i := 1; i < len(binaryArray); i++ {
		parentIndex := math.Floor(float64((i - 1) / 2))
		if i%2 == 1 {
			currentTree := &TreeNode{Val: binaryArray[i]}
			treeArray[int(parentIndex)].Left = currentTree
			treeArray = append(treeArray, currentTree)
		} else {
			currentTree := &TreeNode{Val: binaryArray[i]}
			treeArray[int(parentIndex)].Right = currentTree
			treeArray = append(treeArray, currentTree)
		}
	}
	return treeArray
}
