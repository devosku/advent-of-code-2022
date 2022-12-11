package elftools

// filetype 0 for dir 1 for file
type Node struct {
	FileType int
	Name     string
	Parent   *Node
	Children map[string]*Node
	Size     int
}
