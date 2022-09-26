package main

import (
	"fmt"
	"sort"
)

// ❌自己的办法
func isAnagram(s string, t string) bool {
	// 两个字符相等的情况
	if s == t {
		return true
	}

	// 比较两个字符串每个字母 ASCII 加起来的总值
	s_arr := []rune(s)
	s_num := 0
	for _, v := range s_arr {
		i1 := int(v)
		s_num = s_num + i1
	}

	t_arr := []rune(t)
	t_num := 0
	for _, v := range t_arr {
		i1 := int(v)
		t_num = t_num + i1
	}

	return s_num == t_num
}

// ☑️ 看了评论后 改进的办法
func isAnagramFix(s string, t string) bool {
	var funCompare = func(s1 string) [26]int {
		var arr26 [26]int
		r := []rune(s1)

		for _, v := range r {
			i1 := int(v) - 97
			arr26[i1]++
		}
		return arr26
	}

	return funCompare(s) == funCompare(t)
}

// ☑️ 官方的办法：排序
func isAnagramSort(s string, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}

// ☑️ 官方的办法：哈希表
func isAnagramHash(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}

// ☑️ 进阶解答
func isAnagramUtf(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}

func main() {

	// 自己的办法 这个用例通不过，a = 97 c=99 b = 98
	s1 := "ac"
	t1 := "bb"

	fmt.Println(isAnagram(s1, t1)) // Fail
	fmt.Println(isAnagramFix(s1, t1))
	fmt.Println(isAnagramSort(s1, t1))
	fmt.Println(isAnagramHash(s1, t1))
	fmt.Println(isAnagramUtf(s1, t1))
}
