package day_seven

type Dir struct {
	name     string
	parent   *Dir
	children map[string]*Dir
	contents []*File
	isRoot   bool
}

type File struct {
	size int
	name string
	ext  string
}
