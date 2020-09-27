package deh

type substitution struct {
	FromText string
	ToText   string
}

var hashTable []*substitution = nil
var hashTableEntries = 0
var hashTableLength = 0

func strhash(s string) int {
	p := 0
	h := int(s[p])
	if h != 0 {
		p += 1
		for ; p < len(s); p += 1 {
			h = (h << 5) - h + int(s[p])
		}
	}
	return h
}

func substitutionForString(s string) *substitution {
	if hashTableLength <= 0 {
		return nil
	}

	entry := strhash(s) % hashTableLength

	for ; hashTable[entry] != nil; entry = (entry + 1) % hashTableLength {
		if hashTable[entry].FromText == s {
			return hashTable[entry]
		}
	}
	return nil
}

func String(s string) string {
	subst := substitutionForString(s)
	if subst != nil {
		return subst.ToText
	}
	return s
}
