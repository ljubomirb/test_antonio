package main

//IsAnagram returns true if strings are anagrams (mirror-likes of each other)
func IsAnagram(testStringA string, testStringB string) bool {

	/*
		strings must be the same size
	*/
	if len(testStringA) != len(testStringB) {
		return false
	}

	strLen := len(testStringA)

	for i := 0; i < strLen; i++ {
		if testStringA[i] != testStringB[strLen-1-i] {
			return false
		}
	}

	return true
}
