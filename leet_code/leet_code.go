package leet_code

import (
	"sort"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	var first rune
	rs := make([][]rune, 0, len(strs))
	for i := range strs {
		if len(strs[i]) == 0 {
			return ""
		}
		r := []rune(strs[i])
		if i == 0 {
			first = r[0]
		} else {
			if r[0] != first {
				return ""
			}
		}
		rs = append(rs, r)
	}
	firstRunes := rs[0]
	for i := 1; i < len(firstRunes); i++ {
		for j := 1; j < len(rs); j++ {
			if len(rs[j]) == i {
				return string(firstRunes[:i])
			}
			if firstRunes[i] != rs[j][i] {
				return string(firstRunes[:i])
			}
		}
	}
	return strs[0]
}

func isValid(s string) bool {
	var m = make(map[rune]int, 6)
	m['('] = 1
	m[')'] = -1
	m['{'] = 2
	m['}'] = -2
	m['['] = 3
	m[']'] = -3
	rs := []rune(s)
	res := make([]int, 0, len(rs))
	for i := range rs {
		if num, ok := m[rs[i]]; ok {
			res = append(res, num)
		}
	}
	if len(res)%2 != 0 {
		return false
	}
	return check(res, []int{}, 0)
}

func check(nums, mark []int, index int) bool {
	if index > len(nums)-1 {
		if len(mark) != 0 {
			return false
		}
		return true
	}
	if nums[index] > 0 {
		if len(nums) >= index+2 {
			if nums[index+1] > 0 {
				mark = append(mark, nums[index], nums[index+1])
				index += 2
				return check(nums, mark, index)
			} else {
				if nums[index]+nums[index+1] != 0 {
					return false
				}
				index += 2
				return check(nums, mark, index)
			}
		} else {
			return false
		}
	} else {
		if len(mark) == 0 {
			return false
		}
		if nums[index]+mark[len(mark)-1] != 0 {
			return false
		}

		temp := make([]int, len(mark)-1)
		copy(temp, mark[:len(mark)-1])
		index++
		return check(nums, temp, index)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	l := &ListNode{}
	if list1 == nil {
		l.Val = list2.Val
		l.Next = mergeTwoLists(nil, list2.Next)
	} else if list2 == nil {
		l.Val = list1.Val
		l.Next = mergeTwoLists(nil, list1.Next)
	} else {
		val1 := list1.Val
		val2 := list2.Val
		if val1 <= val2 {
			l.Val = val1
			l.Next = mergeTwoLists(list1.Next, list2)
		} else {
			l.Val = val2
			l.Next = mergeTwoLists(list1, list2.Next)
		}
	}
	return l
}

func removeDuplicates(nums []int) int {
	var (
		num    int
		length int
	)
	if len(nums) == 1 {
		return 1
	}
	num = nums[0]
	length = 1
	for i := 1; i < len(nums); i++ {
		nums[length] = nums[i]
		if num == nums[i] {
			continue
		} else {
			num = nums[i]
			length++
		}
	}
	return length
}

func removeElement(nums []int, val int) int {
	var l int
	length := len(nums)
	if length == 0 {
		return 0
	}
	for i := 0; i < length; i++ {
		if i+l >= length {
			break
		}

		if nums[i] == val {
			l++
			for {
				if i+l >= length {
					break
				}
				if nums[length-l] == val {
					l++
				} else {
					nums[i], nums[length-l] = nums[length-l], nums[i]
					break
				}
			}

		}
	}
	return length - l
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l = &ListNode{}
	if l1 == nil && l2 == nil {
		return l
	}
	var over bool
	if l1 != nil {
		l.Val += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		l.Val += l2.Val
		l2 = l2.Next
	}
	if l.Val > 9 {
		l.Val = l.Val % 10
		over = true
	}
	var next *ListNode
	if l1 != nil || l2 != nil || over {
		next = &ListNode{}
		l.Next = next
	}
	for l1 != nil || l2 != nil || over {
		if over {
			next.Val = 1
		}
		if l1 != nil {
			next.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			next.Val += l2.Val
			l2 = l2.Next
		}
		if next.Val > 9 {
			next.Val = next.Val % 10
			over = true
		} else {
			over = false
		}
		if l1 != nil || l2 != nil || over {
			next_ := &ListNode{}
			next.Next = next_
			next = next_
		}
	}
	return l
}

func maxCount(m int, n int, ops [][]int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	if len(ops) == 0 {
		return m * n
	}
	var (
		x, y int
	)
	for i := range ops {
		if len(ops[i]) != 2 {
			continue
		}
		if ops[i][0] == 0 || ops[i][1] == 0 {
			continue
		}
		if 0 < ops[i][0] && ops[i][0] <= m {
			if x == 0 || ops[i][0] < x {
				x = ops[i][0]
			}
		}

		if 0 < ops[i][1] && ops[i][1] <= n {
			if y == 0 || ops[i][1] < y {
				y = ops[i][1]
			}
		}
	}
	return x * y
}

/*
给你一个下标从 0 开始的整数数组 nums 。在一步操作中，你可以执行以下步骤：

从 nums 选出 两个 相等的 整数
从 nums 中移除这两个整数，形成一个 数对
请你在 nums 上多次执行此操作直到无法继续执行。

返回一个下标从 0 开始、长度为 2 的整数数组 answer 作为答案，
其中 answer[0] 是形成的数对数目，answer[1] 是对 nums 尽可能执行上述操作后剩下的整数数目。
*/
func numberOfPairs(nums []int) []int {
	var (
		res  = make([]int, 2)
		mark = make(map[int]int, len(nums))
	)
	for i := range nums {
		mark[nums[i]]++
	}
	for _, v := range mark {
		div := v / 2
		mod := v % 2
		res[0] += div
		if mod != 0 {
			res[1]++
		}
	}
	return res
}

/*
给你一个下标从 0 开始的数组 nums ，数组中的元素都是 正 整数。请你选出两个下标 i 和 j（i != j），
且 nums[i] 的数位和 与  nums[j] 的数位和相等。

请你找出所有满足条件的下标 i 和 j ，找出并返回 nums[i] + nums[j] 可以得到的 最大值
*/
func maximumSum(nums []int) int {
	var (
		res   = -1
		total int
		max   int
		t     int
		ok    bool
		marks = make(map[int]int, len(nums))
	)
	for i := range nums {
		total = parseNum(nums[i])
		if max, ok = marks[total]; ok {
			t = max + nums[i]
			if t > res {
				res = t
			}
			if max < nums[i] {
				marks[total] = nums[i]
			}
		} else {
			marks[total] = nums[i]
		}
	}
	return res
}

func parseNum(num int) int {
	var (
		mod int
		res int
	)
	for {
		if num == 0 {
			break
		}
		mod = num % 10
		num = num / 10
		res += mod
	}
	return res
}

/**
给你两个正整数数组 nums 和 numsDivide 。你可以从 nums 中删除任意数目的元素。

请你返回使 nums 中 最小 元素可以整除 numsDivide 中所有元素的 最少 删除次数。如果无法得到这样的元素，返回 -1 。

如果 y % x == 0 ，那么我们说整数 x 整除 y
*/
func minOperations(nums []int, numsDivide []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var (
		i   int
		has = false
	)
	for i = range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		var pass = true
		for j := range numsDivide {
			if numsDivide[j]%nums[i] != 0 {
				pass = false
				break
			}
		}
		if pass {
			has = true
			break
		}
	}
	if has {
		return i
	}
	return -1
}

/**
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度
*/
func longestValidParentheses(s string) int {
	bs := []byte(s)
	for {
		var has bool
		var prefix = -1
		for i := 0; i < len(bs); i++ {
			if bs[i] == 0 {
				continue
			}
			if prefix == -1 {
				prefix = i
				continue
			}
			if bs[i] == 41 && bs[prefix] == 40 {
				bs[i] = 0
				bs[prefix] = 0
				has = true
				prefix = -1
			} else {
				prefix = i
			}
		}
		if !has {
			break
		}
	}
	var (
		count int
		max   int
	)
	for i := range bs {
		if bs[i] == 0 {
			count++
		} else {
			if max < count {
				max = count
			}
			count = 0
		}
	}
	if max < count {
		max = count
	}
	return max
}
