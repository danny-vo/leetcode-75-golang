package longest_palindromic_substring


func longestPalindrome(s string) string {
    longest_palindrome := ""
    
    for i := 0; i < len(s); i++ {
        // Determine longest string between existing, expanse originating from
        // current rune and current rune with next rune
        longest_palindrome = longer_string(
            longest_palindrome,
            longer_string(
                expand(i, i, s),
                expand(i, i+1, s),
            ),
        )
    }
    
    return longest_palindrome
}


func expand(left, right int, s string) string {
    // Out of index conditions
    if right >= len(s) || left < 0 {
        return ""
    }
    
    // Iterate outwards from both left and right,
    // until left/rightmost runes do not match or out of bounds
    for ; left >= 0 && right < len(s) &&
        s[left] == s[right];
    left, right = left-1, right+1 {}
    
    // Note left++ is needed due to the nature of the for loop decrement
    // right++ is not due to the slice bounding off at right
    return s[left+1 : right]
}


func longer_string(s0, s1 string) string {
    if len(s0) > len(s1) {
        return s0
    }
    
    return s1
}
