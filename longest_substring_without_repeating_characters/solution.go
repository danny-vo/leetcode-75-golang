package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
    set := make(map[rune]int)
    caboose, leader := 0, 0
    highest_len := 0
    
    for leader < len(s) {
        // Check if rune has already existed - if so the new "start point" is that rune's location + 1
        if idx, ok := set[rune(s[leader])]; ok {
            highest_len = max(highest_len, leader - caboose)
            for caboose < idx + 1 {
                delete(set, rune(s[caboose]))
                caboose++
            }
        }
        
        // Add rune and index to set
        set[rune(s[leader])] = leader
        leader++
    }
    
    return max(highest_len, leader - caboose)
}


func max(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}
