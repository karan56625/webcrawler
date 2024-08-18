package models

// Node represents a in the sitemap with its children.
type Node struct {
	URI      string
	Children []*Node
}
