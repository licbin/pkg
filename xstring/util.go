package xstring

// IsCharacterUnique - Determine whether all characters  are different.
// abaa ==>  false
// abcdef ==> true
func IsCharacterUnique(astr string) bool {
	mark := 0
	for _, v := range astr {
		iv := int(v) - int('a')
		if (mark & (1 << iv)) != 0 {
			return false
		}
		mark |= 1 << iv
	}
	return true
}
