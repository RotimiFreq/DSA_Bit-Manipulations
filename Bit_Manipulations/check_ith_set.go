package Bit_Manipulations

func Check_ith_set(Num, ind int) bool {

	check := 1
	check = check << ind
	result := Num & check

	if result == 0 {
		return false
	} else {
		return true
	}

}
