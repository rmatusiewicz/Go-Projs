package MatchingStrings

func MatchingStrings(strings []string, queries []string) []int32 {
	stringMap := genStrCountMap(strings)
	var res []int32
	for _, qry := range queries {
		res = append(res, int32(stringMap[qry]))
	}
	return res
}

func genStrCountMap(strs []string) map[string]int {
	m := make(map[string]int)
	for _, str := range strs {
		if elmCnt, ok := m[str]; ok {
			m[str] = elmCnt + 1
		} else {
			m[str] = 1
		}
	}
	return m
}
