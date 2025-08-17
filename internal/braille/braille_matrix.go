package braille

import "image/color"

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

// SetFromColor sets the value of the BrailleMatrix at the specified dot position
// based on the provided color. If the color is black, the dot is considered raised (1),
// otherwise it is not raised (0). The inverted parameter allows for inverting the logic.
func (bm *BrailleMatrix) SetFromColor(dotX, dotY int, c color.Color, inverted bool) {
	if c == color.Black {
		if inverted {
			bm[dotX][dotY] = 1
		} else {
			bm[dotX][dotY] = 0
		}
	} else {
		if inverted {
			bm[dotX][dotY] = 0
		} else {
			bm[dotX][dotY] = 1
		}
	}
}

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
func (bm BrailleMatrix) ToRune() rune {
	orderedArray := [8]int{
		bm[0][0], bm[0][1], bm[0][2], bm[1][0],
		bm[1][1], bm[1][2], bm[0][3], bm[1][3],
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

// ToString converts the BrailleMatrix to its string representation,
func (bm BrailleMatrix) ToString() string {
	return string(bm.ToRune())
}
