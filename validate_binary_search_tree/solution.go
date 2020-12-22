package validate_binary_search_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isValidBST(root *TreeNode) bool {
    return validate_bst(root, nil, nil)
}


func validate_bst(root *TreeNode, low, high *int) bool {
    // Done here, empty trees ok
    if nil == root {
        return true
    }
    
    // Check with overall tree health using previous values
    if  (nil != low && root.Val <= *low) ||
        (nil != high && root.Val >= *high) {
            
        return false
    }
    
    // Next iteration left subtree, current value then becomes the max
    // vice-versa for right subtree where current value then becomes the min
    return (
        validate_bst(root.Left, low, &root.Val) &&
        validate_bst(root.Right, &root.Val, high))
}
