package repository

import "strings"

func (t *TrieNode) Insert(city City) {
	node := t
	for _, char := range strings.ToLower(city.Name) {
		if node.children == nil {
			node.children = make(map[rune]*TrieNode)
		}
		if _, exists := node.children[char]; !exists {
			node.children[char] = &TrieNode{}
		}
		node = node.children[char]
	}
	node.cities = append(node.cities, city)
}

func (t *TrieNode) Search(prefix string) []City {
	node := t
	for _, char := range strings.ToLower(prefix) {
		if node.children[char] == nil {
			return nil
		}
		node = node.children[char]
	}
	return node.collectCities()
}

func (t *TrieNode) collectCities() []City {
	var results []City
	queue := []*TrieNode{t}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		results = append(results, curr.cities...)
		for _, child := range curr.children {
			queue = append(queue, child)
		}
	}
	return results
}
