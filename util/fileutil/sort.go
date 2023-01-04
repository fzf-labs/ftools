package fileutil

type FileName []string

func (f FileName) Len() int {
	return len(f)
}

func (f FileName) Less(i, j int) bool {
	return fsort(fsplit(f[i]), fsplit(f[j]))
}

func (f FileName) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func fsplit(str string) []int32 {
	ints := make([]int32, 0)
	var tmp int32
	for _, v := range str {
		if v >= 48 && v <= 57 {
			tmp = tmp*100 + v
			continue
		} else {
			if tmp != 0 {
				ints = append(ints, tmp)
			}
			tmp = v
		}
		ints = append(ints, tmp)
		tmp = 0
	}
	return ints
}

func fsort(i []int32, j []int32) bool {
	lenI := len(i)
	lenJ := len(j)
	total := lenI
	if total < lenJ {
		total = lenJ
	}
	for p := 0; p < total; p++ {
		if p > lenI {
			return true
		}
		if p > lenJ {
			return false
		}
		if i[p] > j[p] {
			return false
		}
		if i[p] < j[p] {
			return true
		}
	}
	return false
}
