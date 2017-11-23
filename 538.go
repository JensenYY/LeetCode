package main

import (
	"fmt"
	//"github.com/golang-collections/collections/stack"
)

/**
 * Definition for a binary tree node.*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
//This can get the right answer,but can not AC
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	ts := stack.New()
	coms := stack.New()
	r := root
	//Inorder Traversal the tree,and put the node to stack
	for r != nil || ts.Len() > 0 {
		for r != nil {
			ts.Push(r)
			r = r.Left
		}
		for ts.Len() > 0 {
			r = ts.Pop().(*TreeNode)
			coms.Push(r)
			r = r.Right
		}
	}
	count := 0
	for coms.Len() > 0 {
		if count == 0 {
			count = coms.Pop().(*TreeNode).Val
		}
		n := coms.Pop().(*TreeNode)
		n.Val += count
		count = n.Val
	}
	return root
}*/

/*for AC this code, I have to use c++... XD
class Solution {
private:
    int sum = 0;
public:
    TreeNode* convertBST(TreeNode* root) {
        if(NULL == root){
            return NULL;
        }
        convertBST(root->right);
        root->val += sum;
        sum = root->val;
        convertBST(root->left);
        return root;
    }
};
*/

//The code learn from a c++ code
func computeSum(root *TreeNode, currsum *int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right = computeSum(root.Right, currsum)
	root.Val += *currsum
	*currsum = root.Val
	root.Left = computeSum(root.Left, currsum)
	return root
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	currsum := 0
	return computeSum(root, &currsum)
}

/*The golang code fasted in leetcode
func convertBSTHelper(n *TreeNode, sum *int) {
    if n == nil {
        return
    }
    
    convertBSTHelper(n.Right, sum)
    *sum += n.Val
    n.Val = *sum
    convertBSTHelper(n.Left, sum)
}

func convertBST(root *TreeNode) *TreeNode {
    var sum int
    convertBSTHelper(root, &sum)
    return root
}
*/

func main() {
	var n TreeNode
	var n1 TreeNode
	var n2 TreeNode
	n.Val = 5
	n1.Val = 2
	n2.Val = 13
	n.Left = &n1
	n.Right = &n2
	convertBST(&n)
	fmt.Println(n.Val, n1.Val, n2.Val)
}
