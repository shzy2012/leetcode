/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 *
 * https://leetcode-cn.com/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (62.31%)
 * Likes:    253
 * Dislikes: 0
 * Total Accepted:    30.7K
 * Total Submissions: 46.4K
 * Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n' +
  '[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
 *
 * 示例:
 *
 * Trie trie = new Trie();
 *
 * trie.insert("apple");
 * trie.search("apple");   // 返回 true
 * trie.search("app");     // 返回 false
 * trie.startsWith("app"); // 返回 true
 * trie.insert("app");
 * trie.search("app");     // 返回 true
 *
 * 说明:
 *
 *
 * 你可以假设所有的输入都是由小写字母 a-z 构成的。
 * 保证所有输入均为非空字符串。
 *
 *
*/

package main

import "fmt"

// @lc code=start
type Trie struct {
	value    string
	children map[rune]*Trie
	isWord   bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{value: "", children: make(map[rune]*Trie), isWord: false}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {

	root := this
	for _, w := range word {

		node := root.match(string(w))
		if node != nil {
			root = node
			continue
		}

		root.children[w] = &Trie{
			value:    string(w),
			children: make(map[rune]*Trie),
			isWord:   false,
		}
		root = root.children[w]
	}

	root.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if n := this.match(word); n != nil && n.isWord {
		return true
	} else {
		return false
	}
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if n := this.match(prefix); n != nil {
		return true
	} else {
		return false
	}
}

func (this *Trie) match(word string) *Trie {
	root := this
	for _, v := range word {
		if _, ok := root.children[v]; ok {
			root = root.children[v]
		} else {
			return nil
		}
	}

	return root
}

func printTree(node *Trie) {

	if len(node.children) <= 0 {
		return
	}

	for _, v := range node.children {
		fmt.Printf("%s-%v\n", v.value, len(v.children))
		printTree(v)
	}
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	println(trie.Search("apple"))   // 返回 true
	println(trie.Search("app"))     // 返回 false
	println(trie.StartsWith("app")) // 返回 true
	trie.Insert("app")
	println(trie.Search("app")) // 返回 true
}
