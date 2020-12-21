package container_with_most_water


func maxArea(height []int) int {
    max_area, current_area := 0, 0
    
    for left, right := 0, len(height)-1; left < right; {
        current_area = (right - left) * min(height[left], height[right])
        max_area = max(max_area, current_area)
        
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
