package structs

type Student struct {
	Name  string
	Score int
}

type StuSlice []Student

func (ss StuSlice) Len() int {
	return len(ss)
}
func (ss StuSlice) Less(i, j int) bool {
	return ss[i].Score < ss[j].Score
}
func (ss StuSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
