package data

func MakeItUnique(strSlc []string) []string {
	res := []string{}
	mapRes := map[string]int{}

	for i, char := range strSlc {
		mapRes[char] = i
	}

	for key := range mapRes {
		res = append(res, key)
	}

	return res
}
