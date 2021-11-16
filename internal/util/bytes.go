package util

func moveByte(n int) int {
	if n == 0 {
		return 0
	}
	return n * 8
}

func WriteToBytes(num int, slice []byte) {
	l := len(slice)
	for i := 0; i < l; i++ {
		slice[l-i-1] = byte(num >> moveByte(i))
	}
}

func BytesToInt(data []byte) int {
	l := len(data)
	sum := 0
	for i := 0; i < l; i++ {
		sum += int(data[l-i-1]) << moveByte(i)
	}
	return sum
}
