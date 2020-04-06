/*
Package string ...
*/
package string

/*
Given an array of strings, group anagrams together.

Example:

Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
*/

//GroupAnagrams ...
func GroupAnagrams(strs []string) [][]string {
	/*
		memory := make(map[string][]string)
		res := make([][]string, 0, len(strs)) //relocation???
		var chars []byte
		var key string
		for _, str := range strs {
			chars = []byte(str)
			sort.Slice(chars, func(i, j int) bool {
				return chars[i] < chars[j]
			})
			key = string(chars)
			memory[key] = append(memory[key], str)
		}
		for _, v := range memory {
			res = append(res, v)
		}
		return res
	*/

	res := make([][]string, 0, len(strs))
	if len(strs) == 0 {
		return res
	}
	m := make(map[[26]uint8][]string)
	for _, str := range strs {
		var arr [26]uint8
		for _, char := range str {
			arr[char-'a']++
		}
		m[arr] = append(m[arr], str)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
