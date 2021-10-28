package list

type item string

func newItem(name string) item {
	return item(name)
}
