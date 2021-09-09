package common

type StrSlice []string

func (list StrSlice) Has(s string) bool {
	for _, b := range list {
		if b == s {
			return true
		}
	}

	return false
}
