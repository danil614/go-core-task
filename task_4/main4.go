package task_4

type void struct{}

func subSlice(slice1, slice2 []string) []string {
	if len(slice1) == 0 {
		return []string{}
	}

	out := make([]string, 0, len(slice1))

	set := map[string]void{}
	for _, v := range slice2 {
		set[v] = void{}
	}

	for _, v := range slice1 {
		if _, ok := set[v]; !ok {
			out = append(out, v)
		}
	}

	return out
}
