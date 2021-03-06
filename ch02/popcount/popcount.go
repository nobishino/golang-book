package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount2_5(x uint64) (result int) {
	for ; x > 0; x = x & (x - 1) {
		result++
	}
	return
}

func PopCount2_4(x uint64) int {
	var answer uint64
	for i := 0; i < 64; i++ {
		answer += x & 1
		x = x >> 1
	}
	return int(answer)
}

func PopCount2_3(x uint64) int {
	var b byte
	for i := 0; i < 8; i++ {
		b += pc[byte(x>>(i*8))]
	}
	return int(b)
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
