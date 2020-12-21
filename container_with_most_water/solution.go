package container_with_most_water


func maxArea(height []int) int {
    max_area, current_area := 0, 0
    
    // Naviagte the array in O(N) by traversing with two pointers
    for left, right := 0, len(height)-1; left < right; {
        // Area for rectangle is L * H (H is constarained by the shorter "bar")
        current_area = (right - left) * min(height[left], height[right])
        max_area = max(max_area, current_area)
        
        // The pointer to the shorter element is moved on fro
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    
    return max_area
}


func max(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}


func min(a, b int) int {
    if a < b {
        return a
    }
    
    return b
}
