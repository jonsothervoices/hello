package bits

func getBit(n, b int) bool { return (1<<b)&n != 0 }

func setBit(n, b int) int { return (1 << b) | n }

func clearBit(n, b int) int { return n & (^(1 << b)) }

func updateBit(n, b int, v bool) int {
	if v {
		return setBit(n, b)
	}
	return clearBit(n, b)
}
