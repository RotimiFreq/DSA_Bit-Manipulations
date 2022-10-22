package Bit_Manipulations


func SwapBit(bit uint) uint {
	var even_bits uint
	var odd_bits uint

	even_bits = bit & 0xAAAAAAAA
	odd_bits = bit & 0x55555555

	even_bits >>= 1
	odd_bits <<= 1

	return even_bits | odd_bits

}


