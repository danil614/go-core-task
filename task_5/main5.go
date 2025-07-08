package task_5

type void struct{}

func intersectSlices(slice1, slice2 []int) (bool, []int) {
	intersection := make([]int, 0, len(slice1))

	if len(slice1) == 0 || len(slice2) == 0 {
		return false, intersection
	}

	// Создаем set для первого слайса
	set := make(map[int]void)
	for _, v := range slice2 {
		set[v] = void{}
	}

	seen := make(map[int]void)
	isIntersect := false

	// Проверяем пересечения и собираем результат
	for _, v := range slice1 {
		if _, exists := set[v]; exists {
			if _, ok := seen[v]; !ok {
				seen[v] = void{}
				intersection = append(intersection, v)
				isIntersect = true
			}
		}
	}

	return isIntersect, intersection
}
