package resolve_captcha

func checkNumberZero(pixel []Pixel) bool {
	var (
		numberZero []bool
		px         Pixel
	)
	px = pixel[3]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[4]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[5]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[9]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[10]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[11]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[12]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[13]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[14]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[15]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[17]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[18]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[22]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[23]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[24]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[25]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[31]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[32]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[33]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[39]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[40]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[41]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[47]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[48]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[49]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[55]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[56]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[57]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[63]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[64]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[65]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[71]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[73]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[74]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[78]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[79]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[81]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[82]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[83]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[84]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[85]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[86]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[87]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[91]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[92]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[93]
	numberZero = append(numberZero, px.R != 255, px.G != 255, px.B != 255)

	notNumberZero := false
	for _, zero := range numberZero {
		if !zero {
			notNumberZero = true
		}
	}
	return !notNumberZero
}

func checkNumberOne(pixel []Pixel) bool {
	var (
		numberOne []bool
		px        Pixel
	)

	px = pixel[2]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[3]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[8]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[9]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[10]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[11]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[16]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[17]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[18]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[19]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[26]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[27]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[34]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[35]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[42]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[43]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[50]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[51]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[58]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[59]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[66]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[67]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[74]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[75]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[80]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[81]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[82]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[83]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[84]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[85]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[88]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[89]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[90]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[91]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[92]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[93]
	numberOne = append(numberOne, px.R != 255, px.G != 255, px.B != 255)

	notNumberOne := false
	for _, one := range numberOne {
		if !one {
			notNumberOne = true
		}
	}
	return !notNumberOne
}

func checkNumberTwo(pixel []Pixel) bool {
	var (
		numberTwo []bool
		px        Pixel
	)
	px = pixel[1]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[2]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[3]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[4]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[5]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[8]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[9]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[10]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[11]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[12]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[13]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[14]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[16]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[22]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[23]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[30]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[31]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[38]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[39]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[45]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[46]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[51]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[52]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[53]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[58]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[59]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[65]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[66]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[72]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[73]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[80]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[81]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[82]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[83]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[84]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[85]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[86]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[87]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[88]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[89]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[90]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[91]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[92]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[93]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[94]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[95]
	numberTwo = append(numberTwo, px.R != 255, px.G != 255, px.B != 255)

	notNumberTwo := false
	for _, two := range numberTwo {
		if !two {
			notNumberTwo = true
		}
	}
	return !notNumberTwo
}

func checkNumberThree(pixel []Pixel) bool {
	var (
		numberThree []bool
		px          Pixel
	)
	px = pixel[1]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[2]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[3]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[4]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[5]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[8]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[9]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[10]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[11]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[12]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[13]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[14]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[16]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[21]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[22]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[23]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[30]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[31]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[37]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[38]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[41]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[42]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[43]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[44]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[49]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[50]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[51]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[52]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[53]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[54]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[61]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[62]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[63]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[70]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[71]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[72]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[77]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[78]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[79]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[80]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[81]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[82]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[83]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[84]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[85]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[86]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[89]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[90]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[91]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[92]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[93]
	numberThree = append(numberThree, px.R != 255, px.G != 255, px.B != 255)

	notNumberThree := false
	for _, three := range numberThree {
		if !three {
			notNumberThree = true
		}
	}
	return !notNumberThree
}

func checkNumberFour(pixel []Pixel) bool {
	var (
		numberFour []bool
		px         Pixel
	)
	px = pixel[6]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[7]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[13]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[14]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[15]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[20]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[21]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[22]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[23]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[27]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[28]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[30]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[31]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[34]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[35]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[38]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[39]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[41]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[42]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[46]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[47]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[48]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[49]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[54]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[55]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[56]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[57]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[58]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[59]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[60]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[61]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[62]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[63]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[64]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[65]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[66]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[67]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[68]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[69]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[70]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[71]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[78]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[79]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[86]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[87]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[94]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[95]
	numberFour = append(numberFour, px.R != 255, px.G != 255, px.B != 255)

	notNumberFour := false
	for _, four := range numberFour {
		if !four {
			notNumberFour = true
		}
	}
	return !notNumberFour
}

func checkNumberFive(pixel []Pixel) bool {
	var (
		numberFive []bool
		px         Pixel
	)

	px = pixel[0]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[1]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[2]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[3]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[4]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[5]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[6]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[7]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[8]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[9]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[10]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[11]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[12]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[13]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[14]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[15]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[16]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[17]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[24]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[25]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[32]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[33]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[34]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[35]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[36]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[40]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[41]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[42]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[43]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[44]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[45]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[46]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[53]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[54]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[55]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[62]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[63]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[70]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[71]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[72]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[77]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[78]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[79]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[80]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[81]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[82]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[83]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[84]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[85]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[86]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[89]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[90]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[91]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[92]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[93]
	numberFive = append(numberFive, px.R != 255, px.G != 255, px.B != 255)

	notNumberFive := false
	for _, five := range numberFive {
		if !five {
			notNumberFive = true
		}
	}
	return !notNumberFive
}

func checkNumberSix(pixel []Pixel) bool {
	var (
		numberSix []bool
		px        Pixel
	)

	px = pixel[2]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[3]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[4]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[5]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[8]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[9]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[10]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[11]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[12]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[13]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[14]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[16]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[17]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[22]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[24]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[25]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[32]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[34]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[35]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[36]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[37]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[40]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[41]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[42]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[43]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[44]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[45]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[46]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[48]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[49]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[53]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[54]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[55]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[56]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[62]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[63]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[64]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[70]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[71]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[72]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[73]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[77]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[78]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[79]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[80]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[81]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[82]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[83]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[84]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[85]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[86]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[90]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[91]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[92]
	numberSix = append(numberSix, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[93]

	notNumberSix := false
	for _, six := range numberSix {
		if !six {
			notNumberSix = true
		}
	}
	return !notNumberSix
}

func checkNumberSeven(pixel []Pixel) bool {
	var (
		numberSeven []bool
		px          Pixel
	)

	px = pixel[0]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[1]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[2]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[3]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[4]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[5]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[6]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[7]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[8]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[9]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[10]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[11]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[12]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[13]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[14]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[15]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[22]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[23]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[29]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[30]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[36]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[37]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[44]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[45]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[51]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[52]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[59]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[60]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[66]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[67]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[74]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[75]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[81]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[82]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[89]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[90]
	numberSeven = append(numberSeven, px.R != 255, px.G != 255, px.B != 255)

	notNumberSeven := false
	for _, seven := range numberSeven {
		if !seven {
			notNumberSeven = true
		}
	}
	return !notNumberSeven
}

func checkNumberEight(pixel []Pixel) bool {
	var (
		numberEight []bool
		px          Pixel
	)
	px = pixel[1]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[2]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[3]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[4]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[5]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[8]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[9]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[10]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[11]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[12]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[13]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[14]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[16]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[17]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[21]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[22]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[24]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[25]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[29]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[30]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[32]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[33]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[34]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[36]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[37]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[41]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[42]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[43]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[43]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[44]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[45]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[48]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[49]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[51]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[52]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[53]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[54]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[56]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[61]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[62]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[63]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[64]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[70]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[71]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[72]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[73]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[77]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[78]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[79]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[80]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[81]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[82]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[83]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[84]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[85]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[86]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[89]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[90]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[91]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[92]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[93]
	numberEight = append(numberEight, px.R != 255, px.G != 255, px.B != 255)

	notNumberEight := false
	for _, eight := range numberEight {
		if !eight {
			notNumberEight = true
		}
	}
	return !notNumberEight
}

func checkNumberNine(pixel []Pixel) bool {
	var (
		numberNine []bool
		px         Pixel
	)

	px = pixel[1]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[2]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[3]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[4]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[8]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[9]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[10]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[11]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[12]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[13]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[14]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[16]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[17]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[21]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[22]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[24]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[30]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[31]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[32]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[38]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[39]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[40]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[41]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[45]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[46]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[47]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[48]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[49]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[50]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[51]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[52]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[53]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[54]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[55]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[57]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[58]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[59]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[60]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[62]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[63]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[70]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[71]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[72]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[77]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[78]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[80]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[81]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[82]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[83]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[84]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[85]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[86]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[89]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[90]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[91]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)
	px = pixel[92]
	numberNine = append(numberNine, px.R != 255, px.G != 255, px.B != 255)

	notNumberNine := false
	for _, nine := range numberNine {
		if !nine {
			notNumberNine = true
		}
	}
	return !notNumberNine
}
