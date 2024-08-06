package main

type model struct {
	choises  []string
	cursor   int
	selected map[int]struct{}
}

func initModel() model {

	return model{
		choises: []string{"one", "two", "three", "four", "five"},
		selected: make(map[int]struct{}),
	}

}
