package item

type Item byte

type Priority int64
type priorities map[Item]Priority

var itemPriorities priorities

func init() {
	itemPriorities = generatePriorities()
}

func generatePriorities() priorities {
	p := priorities{}
	for i, j := 'a', 1; i <= 'z'; i, j = i+1, j+1 {
		p[Item(i)] = Priority(j)
	}

	for i, j := 'A', 27; i <= 'Z'; i, j = i+1, j+1 {
		p[Item(i)] = Priority(j)
	}

	return p
}

func (i Item) Priority() Priority {
	return itemPriorities[i]
}
