package two_sum

func twoSum(nums []int, target int) []int {
    // Impossible case
    if len(nums) < 2 {
        return []int{}
    }
    
    // Create map to keep track of numbers and their indexes
    num_idx_map := make(map[int]int)
    
    // Iterate through nums and check for existing diff values
    for i, n := range(nums) {
        diff := target - n
        
        if idx, ok := num_idx_map[diff]; ok {
            return []int{i, idx}
        }
        
        // Haven't passed needed number yet, add current to map
        num_idx_map[n] = i
    }
    
    return []int{}
}
