package bitset

//bit parity, 1 if there is an odd number of 1's, 0 if there is an even
//equivalent of an exclusive or on all the bits in the set
//Code taken from Hacker's Delight, Second Edition
//by Henry S. Warren, Jr  pgs. 97-98

func parity(x uint64) uint64 {
	x = x ^ (x >> 1)
	x = (x ^ (x >>2)) & 0x1111111111111111
	x *=  0x1111111111111111
	return (x >> 60) & 1
}

func paritySliceGo(s []uint64) (par uint64) {
	for _, x := range s {
		par ^= parity(x)
	}
	return
}