package braille

// BrailleMatrix represents a Braille character as a 2x4 matrix of dots.
// Each element should be either 0 or 1, where 1 indicates a raised dot.
//
// The layout of the matrix is as follows:
// ------------------
// | (0,0) | (1,0) |
// | (0,1) | (1,1) |
// | (0,2) | (1,2) |
// | (0,3) | (1,3) |
// ------------------
type BrailleMatrix [2][4]int

// ToRune converts the BrailleMatrix into its corresponding Unicode rune.
//
// According to the Braille Patterns documentation:
// https://en.wikipedia.org/wiki/Braille_Patterns#Identifying.2C_naming_and_ordering
// the Braille dots are numbered as follows:
//
// ---------
// | 1 | 4 |
// | 2 | 5 |
// | 3 | 6 |
// | 7 | 8 |
// ---------
//
// Each dot has a decimal value used to compute the final Unicode code point:
// (Note: the documentation shows hexadecimal values, here we use decimal)
//
// ------------
// | 1   | 8   |
// | 2   | 16  |
// | 4   | 32  |
// | 64  | 128 |
// ------------
//
// The sum of the values of the raised dots (1) converted to a rune plus the base rune '\u2800'
// produces the Unicode rune representing the Braille character.
func (b BrailleMatrix) ToRune() rune {
	orderedArray := [8]int{
		b[0][0], b[0][1], b[0][2], b[1][0],
		b[1][1], b[1][2], b[0][3], b[1][3],
	}

	mappedDecimalValues := [8]int{
		1, 2, 4, 8,
		16, 32, 64, 128,
	}

	sum := 0

	for i, bit := range orderedArray {
		sum += bit * mappedDecimalValues[i]
	}

	return rune(sum) + '\u2800'
}

func (b BrailleMatrix) ToString() string {
	return string(b.ToRune())
}
