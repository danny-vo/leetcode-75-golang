package valid_parentheses


func isValid(s string) bool {
    stack := make(Stack, 0)
    
    // Go per rune in string
    for _, r := range(s) {
        // Add to stack left hand side of parentheses
        if '(' == r || '[' == r || '{' == r {
            stack.Push(r)
        // If right hand side of parentheses, latest in stack should pair correctly
        } else {
            left, ok := stack.Pop()
            if !ok || !is_valid_pair(left, r) {
                return false
            }
        }
    }
    
    // All good if nothing left in stack
    return true && stack.Is_Empty()
}


func is_valid_pair(a, b rune) bool {
    if '(' == a && ')' != b {
        return false
    } else if '[' == a && ']' != b {
        return false
    } else if '{' == a && '}' != b {
        return false
    }
    
    return true
}


type Stack []rune


func (s *Stack) Is_Empty() bool {
    return 0 == len(*s)
}


func (s *Stack) Push(r rune) {
    *s = append(*s, r)
}


func (s *Stack) Pop() (rune, bool) {
    if s.Is_Empty() {
        return rune(0), false
    }
    
    idx := len(*s) - 1
    r := (*s)[idx]
    *s = (*s)[:idx]
    
    return r, true
}
