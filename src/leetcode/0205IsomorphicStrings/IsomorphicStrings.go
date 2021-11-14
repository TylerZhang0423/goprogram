package main

func main() {

}

func isIsomorphic(s string, t string) bool {
	s2tMap := make(map[string]string)
	t2sMap := make(map[string]string)
	for i := 0 ; i < len(s); i++ {
		letterS := string(s[i])
		letterT := string(t[i])
		value1, exist1 := s2tMap[letterS]
		value2, exist2 := t2sMap[letterT]
		if exist1 || exist2 {
			if value1 != letterT || value2 != letterS {
				return false
			}
		} else {
			s2tMap[letterS] = letterT
			t2sMap[letterT] = letterS
		}
	}
	return true
}
