package _90_单词规律

import "strings"

/*

	给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

	这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

	示例1:

	输入: pattern = "abba", str = "dog cat cat dog"
	输出: true
	示例 2:

	输入:pattern = "abba", str = "dog cat cat fish"
	输出: false
	示例 3:

	输入: pattern = "aaaa", str = "dog cat cat dog"
	输出: false
	示例 4:

	输入: pattern = "abba", str = "dog dog dog dog"
	输出: false

	说明:
	你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。

*/

// 解法1 官方题解
/*
	思路及解法
	需要判断字符与字符串之间是否恰好一一对应。即任意一个字符都对应着唯一的字符串，任意一个字符串也只被唯一的一个字符对应。在集合论中，这种关系被称为「双射」。

	想要解决本题，我们可以利用哈希表记录每一个字符对应的字符串，以及每一个字符串对应的字符。然后我们枚举每一对字符与字符串的配对过程，不断更新哈希表，如果发生了冲突，则说明给定的输入不满足双射关系。

	在实际代码中，我们枚举 pattern 中的每一个字符，利用双指针来均摊线性地找到该字符在 str 中对应的字符串。每次确定一个字符与字符串的组合，我们就检查是否出现冲突，最后我们再检查两字符串是否比较完毕即可。
*/
func wordPattern1(pattern string, s string) bool {
	word2ch := map[string]byte{}   // "dog" : 'a'
	ch2word := map[byte]string{}   // 'a' : "dog"
	words := strings.Split(s, " ") // "dog" "dog"
	if len(pattern) != len(words) {
		return false
	}
	for i, word := range words {
		ch := pattern[i]
		if word2ch[word] > 0 && word2ch[word] != ch || ch2word[ch] != "" && ch2word[ch] != word {
			return false
		}
		word2ch[word] = ch
		ch2word[ch] = word
	}
	return true
}
