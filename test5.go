package main

//Permutations is a struc holding methods for string permutations
type Permutations struct {
	resultingPermutations []string
}

//GetPermutations returns array of strings for a given string
//returned strings are permutations (all possible combinations of chars from given string)
func (perm *Permutations) GetPermutations(stringToPermute string) []string {

	if len(stringToPermute) > 5 {
		notSoPoliteReturnOfAngryMachine := []string{"we can do this, but we won't"}
		return notSoPoliteReturnOfAngryMachine
	}

	toPermute := []byte(stringToPermute)
	perm.permute(toPermute, 0, len(stringToPermute)-1)
	return perm.resultingPermutations
}

func (perm *Permutations) permute(toPermute []byte, startingIndex int, endIndex int) {
	if startingIndex == endIndex {
		perm.resultingPermutations = append(perm.resultingPermutations, string(toPermute))
	} else {

		for i := startingIndex; i <= endIndex; i++ {
			//golang can swap like this
			toPermute[startingIndex], toPermute[i] = toPermute[i], toPermute[startingIndex]
			//recursion
			perm.permute(toPermute, startingIndex+1, endIndex)
			//back
			toPermute[startingIndex], toPermute[i] = toPermute[i], toPermute[startingIndex]
		}
	}
}
