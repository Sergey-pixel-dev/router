package router

func IsEqualPath(p1 string, p2 string) bool {
	lenp1 := len(p1)
	lenp2 := len(p2)
	if p1[len(p1)-1] == '/' {
		lenp1--
	}
	if p2[len(p2)-1] == '/' {
		lenp2--
	}
	return p1[:lenp1] == p2[:lenp2]
}

func Contains[T comparable](arr []T, el T) bool {
	for _, it := range arr {
		if it == el {
			return true
		}
	}
	return false
}
