package gobmh

import ()

const (
	ASIZE = 256
)

func IndexHorspool(haystack, needle []byte) int {
	var scan, last, offset, maxoffset, hlen, nlen int
	var badCharSkip [ASIZE]int

	hlen = len(haystack)
	nlen = len(needle)

	if (haystack == nil || needle == nil) || hlen < nlen {
		return -1
	}

	if nlen == 0 {
		return 0
	}

	// --- Preprocessing
	last = nlen - 1
	maxoffset = hlen - nlen

	// Initialize to default value
	// Skip to length of needle, when character doesn't occur in the needle
	for scan = 0; scan < ASIZE; scan++ {
		badCharSkip[scan] = nlen
	}

	// Analyse the needle
	for scan = 0; scan < last; scan++ {
		badCharSkip[needle[scan]] = last - scan
	}

	// --- Searching
	// Find a way to do it without using offset vars, 
	for offset <= maxoffset {
		for scan = last; haystack[scan+offset] == needle[scan]; scan-- {
			if scan == 0 {
				return offset
			}
		}

		offset += badCharSkip[haystack[offset+last]]
	}

	return -1
}
