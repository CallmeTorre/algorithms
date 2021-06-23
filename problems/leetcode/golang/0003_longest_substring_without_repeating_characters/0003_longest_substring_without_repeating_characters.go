//TODO Rethink the approach
/*
Given a string s, find the length of the longest substring without repeating characters.

Example 1:
	Input: s = "abcabcbb"
	Output: 3
	Explanation: The answer is "abc", with the length of 3.

Example 2:
	Input: s = "bbbbb"
	Output: 1
	Explanation: The answer is "b", with the length of 1.

Example 3:
	Input: s = "pwwkew"
	Output: 3
	Explanation: The answer is "wke", with the length of 3.
	Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Example 4:
	Input: s = ""
	Output: 0
*/

/*
"abcabcbb"
"bbbbb"
"pwwkew"
""
" "
"au"
"dvdf"
"qrsvbspk"
*/

func lengthOfLongestSubstring(s string) int {
    longest := 0
    memory := make(map[byte]int)
    windowFrom, windowTo := 0, 0
    for windowTo < len(s){
        currentChar := s[windowTo]
        if _, found := memory[currentChar]; found{
            if len(memory) > longest{
                longest = len(memory)
            }
            for windowFrom <= memory[currentChar]{
                delete(memory, s[windowFrom])
                 windowFrom++
            }
        }
        memory[currentChar] = windowTo
        windowTo++ 
    }
    if len(memory) > longest{
        longest = len(memory)
    }
    return longest
}