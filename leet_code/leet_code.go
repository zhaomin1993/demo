package leet_code

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
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

/*
*
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

/*
*
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

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/
func trap(height []int) int {
	var (
		ans = 0
		h1  = 0
		h2  = 0
		r   = 0
	)
	for i, v := range height {
		if v > h1 {
			h1 = v
		}
		r = height[len(height)-i-1]
		if r > h2 {
			h2 = r
		}
		ans += h1 + h2 - v
	}
	return ans - len(height)*h1
}

/*
给定一个单词数组 words 和一个长度 maxWidth ，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。

你应该使用 “贪心算法” 来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。

要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。

文本的最后一行应为左对齐，且单词之间不插入额外的空格。
*/
func fullJustify(words []string, maxWidth int) []string {
	var (
		length      = 0
		res         = make([]string, 0, len(words)/2)
		temp        = make([]string, 0, 2)
		maxWidthStr = strconv.Itoa(maxWidth)
	)
	for _, v := range words {
		newL := length + len(v)
		if newL < maxWidth {
			length += len(v) + 1
			temp = append(temp, v)
		} else if newL == maxWidth {
			temp = append(temp, v)
			res = append(res, strings.Join(temp, " "))
			length = 0
			temp = make([]string, 0, 2)
		} else {
			l := len(temp)
			gap := l + maxWidth - length
			if l == 1 {
				res = append(res, fmt.Sprintf("%-"+maxWidthStr+"s", temp[0]))
			} else {
				div := gap / (l - 1)
				mod := gap % (l - 1)
				if l == 2 {
					mod = 0
				}
				var s string
				for i := range temp {
					space := div
					if i < mod {
						space++
					}
					s += temp[i]
					if i != len(temp)-1 {
						s += fmt.Sprintf("%"+strconv.Itoa(space)+"s", "")
					}
				}
				res = append(res, s)
			}
			length = len(v) + 1
			temp = make([]string, 0, 2)
			temp = append(temp, v)
		}
	}
	if len(temp) > 0 {
		res = append(res, fmt.Sprintf("%-"+maxWidthStr+"s", strings.Join(temp, " ")))
	}
	return res
}

/*
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
*/
func isMatch(s string, p string) bool {
	sn := len(s)
	pn := len(p)
	i := 0
	j := 0
	start := -1
	match := 0
	for i < sn {
		if j < pn && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < pn && p[j] == '*' {
			start = j
			match = i
			j++
		} else if start != -1 {
			j = start + 1
			match++
			i = match
		} else {
			return false
		}
	}
	for j < pn {
		if p[j] != '*' {
			return false
		}
		j++
	}
	return true
}

/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。
*/
func mergeKLists(lists []*ListNode) *ListNode {
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			lists = append(lists[:i], lists[i+1:]...)
			i--
		}
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 0 {
		return nil
	}
	var node, f = &ListNode{}, &ListNode{}
	f.Next = node

	for {
		var (
			index int
			min   = lists[0].Val
		)
		for i := range lists {
			if lists[i].Val < min {
				min = lists[i].Val
				index = i
			}
		}
		node.Val = min
		lists[index] = lists[index].Next
		if lists[index] == nil {
			lists = append(lists[:index], lists[index+1:]...)
		}
		if len(lists) == 0 {
			break
		}
		node.Next = &ListNode{}
		node = node.Next
	}
	return f.Next
}

/*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；
如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
*/
type LRUCache struct {
	cache    map[int]int
	queue    []int
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]int, capacity),
		queue:    make([]int, 0, capacity),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	value, ok := this.cache[key]
	if !ok { // 不存在返回-1
		value = -1
	} else { // 存在
		// 查找index
		var index = -1
		for i, v := range this.queue {
			if v == key {
				index = i
				break
			}
		}
		// 根据index删除这个key
		this.queue = append(this.queue[:index], this.queue[index+1:]...)
		// 再将key放到最后
		this.queue = append(this.queue, key)
	}
	return value
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.cache[key]
	if ok { // 存在则更新
		// 查找index
		var index = -1
		for i, v := range this.queue {
			if v == key {
				index = i
				break
			}
		}
		// 根据index删除这个key
		this.queue = append(this.queue[:index], this.queue[index+1:]...)
	} else if len(this.cache) == this.capacity { // 不存在并且容量已满
		// 如果缓存已满，则删除queue中第一个key
		delKey := this.queue[0]
		this.queue = this.queue[1:]
		delete(this.cache, delKey)
	}
	// 再将key放到最后
	this.queue = append(this.queue, key)
	// 赋值
	this.cache[key] = value
}

// reverseList 反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var node = &ListNode{}
	for head != nil {
		node.Val = head.Val
		if head.Next != nil {
			n := &ListNode{}
			n.Next = node
			node = n
		}
		head = head.Next
	}
	return node
}

// 链表排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := &ListNode{Next: head}
	quickSort(prev, nil)
	return prev.Next
}

func isOrdered(l, r *ListNode) bool {
	node := l.Next
	temp := node.Val
	for node != r {
		if temp > node.Val {
			return false
		}
		temp = node.Val
		node = node.Next
	}
	return true
}

func partition(l, r *ListNode) *ListNode {
	p := l.Next.Next
	lessHead := &ListNode{}
	lessPointer := lessHead
	moreHead := &ListNode{}
	morePointer := moreHead
	m := l.Next
	for p != r {
		if p.Val <= m.Val {
			lessPointer.Next = p
			lessPointer = lessPointer.Next
		} else {
			morePointer.Next = p
			morePointer = morePointer.Next
		}
		p = p.Next
	}
	// 形成 l->lessHead->lessPointer->m->moreHead->morePointer->r的链表
	if lessHead.Next != nil {
		l.Next = lessHead.Next
		lessPointer.Next = m
	} else {
		l.Next = m
	}
	if moreHead.Next != nil {
		m.Next = moreHead.Next
		morePointer.Next = r
	} else {
		m.Next = r
	}
	return m
}

func quickSort(l, r *ListNode) {
	if l.Next == r || l.Next.Next == r {
		return
	}
	m := partition(l, r)
	// 优化点1 如果有序就返回 最好效率提升到O(n)
	if isOrdered(l, r) {
		return
	}
	quickSort(l, m)
	quickSort(m, r)
}
