package repository

type Repository struct {
	DataSource TrieNode
}

type City struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type TrieNode struct {
	children map[rune]*TrieNode
	cities   []City
}

type LoadDataRequest struct {
	DataPath string
}
