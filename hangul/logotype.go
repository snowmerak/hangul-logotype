package hangul

import (
	"bytes"
)

var engToHan = map[rune]rune{
	'q': 'ㅂ', 'Q': 'ㅂ',
	'w': 'ㅈ', 'W': 'ㅈ',
	'e': 'ㄷ', 'E': 'ㄷ',
	'r': 'ㄱ', 'R': 'ㄱ',
	't': 'ㅅ', 'T': 'ㅅ',
	'y': 'ㅛ', 'Y': 'ㅛ',
	'u': 'ㅕ', 'U': 'ㅕ',
	'i': 'ㅑ', 'I': 'ㅑ',
	'o': 'ㅐ', 'O': 'ㅐ',
	'p': 'ㅔ', 'P': 'ㅔ',
	'a': 'ㅁ', 'A': 'ㅁ',
	's': 'ㄴ', 'S': 'ㄴ',
	'd': 'ㅇ', 'D': 'ㅇ',
	'f': 'ㄹ', 'F': 'ㄹ',
	'g': 'ㅎ', 'G': 'ㅎ',
	'h': 'ㅗ', 'H': 'ㅗ',
	'j': 'ㅓ', 'J': 'ㅓ',
	'k': 'ㅏ', 'K': 'ㅏ',
	'l': 'ㅣ', 'L': 'ㅣ',
	'z': 'ㅋ', 'Z': 'ㅋ',
	'x': 'ㅌ', 'X': 'ㅌ',
	'c': 'ㅊ', 'C': 'ㅊ',
	'v': 'ㅍ', 'V': 'ㅍ',
	'b': 'ㅠ', 'B': 'ㅠ',
	'n': 'ㅜ', 'N': 'ㅜ',
	'm': 'ㅡ', 'M': 'ㅡ',
}

func 이건모음인가(r rune) bool {
	// Check if it's mapped from engToHan
	if han, ok := engToHan[r]; ok {
		switch han {
		case 'ㅑ', 'ㅕ', 'ㅛ', 'ㅠ', 'ㅏ', 'ㅓ', 'ㅗ', 'ㅜ', 'ㅡ', 'ㅣ', 'ㅐ', 'ㅔ':
			return true
		default:
			return false
		}
	}
	// Check direct hangul jamo
	switch r {
	case 'ㅑ', 'ㅕ', 'ㅛ', 'ㅠ', 'ㅏ', 'ㅓ', 'ㅗ', 'ㅜ', 'ㅡ', 'ㅣ', 'ㅐ', 'ㅔ':
		return true
	default:
		return false
	}
}

func 이건자음인가(r rune) bool {
	// Check if it's mapped from engToHan
	if han, ok := engToHan[r]; ok {
		switch han {
		case 'ㄱ', 'ㄴ', 'ㄷ', 'ㄹ', 'ㅁ', 'ㅂ', 'ㅅ', 'ㅇ', 'ㅈ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ', 'ㄲ', 'ㄸ', 'ㅃ', 'ㅆ', 'ㅉ', 'ㄳ', 'ㄵ', 'ㄶ', 'ㄺ', 'ㄻ', 'ㄼ', 'ㄽ', 'ㄾ', 'ㄿ', 'ㅀ', 'ㅄ':
			return true
		default:
			return false
		}
	}
	// Check direct hangul jamo
	switch r {
	case 'ㄱ', 'ㄴ', 'ㄷ', 'ㄹ', 'ㅁ', 'ㅂ', 'ㅅ', 'ㅇ', 'ㅈ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ', 'ㄲ', 'ㄸ', 'ㅃ', 'ㅆ', 'ㅉ', 'ㄳ', 'ㄵ', 'ㄶ', 'ㄺ', 'ㄻ', 'ㄼ', 'ㄽ', 'ㄾ', 'ㄿ', 'ㅀ', 'ㅄ':
		return true
	default:
		return false
	}
}

var breakWords = map[rune]struct{}{
	' ':  {},
	'.':  {},
	',':  {},
	'!':  {},
	'?':  {},
	'\n': {},
	'\t': {},
}

func DefaultBreakWords() map[rune]struct{} {
	bw := make(map[rune]struct{}, len(breakWords))
	for k, v := range breakWords {
		bw[k] = v
	}
	return bw
}

var doubleConsonants = map[[2]rune]rune{
	{'ㄱ', 'ㄱ'}: 'ㄲ',
	{'ㄷ', 'ㄷ'}: 'ㄸ',
	{'ㅂ', 'ㅂ'}: 'ㅃ',
	{'ㅅ', 'ㅅ'}: 'ㅆ',
	{'ㅈ', 'ㅈ'}: 'ㅉ',
	{'ㄱ', 'ㅅ'}: 'ㄳ',
	{'ㄴ', 'ㅈ'}: 'ㄵ',
	{'ㄴ', 'ㅎ'}: 'ㄶ',
	{'ㄹ', 'ㄱ'}: 'ㄺ',
	{'ㄹ', 'ㅁ'}: 'ㄻ',
	{'ㄹ', 'ㅂ'}: 'ㄼ',
	{'ㄹ', 'ㅅ'}: 'ㄽ',
	{'ㄹ', 'ㅌ'}: 'ㄾ',
	{'ㄹ', 'ㅍ'}: 'ㄿ',
	{'ㄹ', 'ㅎ'}: 'ㅀ',
	{'ㅂ', 'ㅅ'}: 'ㅄ',
}

func 겹받침합치기(input []rune) []rune {
	result := make([]rune, 0, len(input))
	for i := 0; i < len(input); i++ {
		if i < len(input)-1 {
			pair := [2]rune{input[i], input[i+1]}
			if dc, ok := doubleConsonants[pair]; ok {
				result = append(result, dc)
				i++
				continue
			}
		}
		result = append(result, input[i])
	}
	return result
}

func writeRuneToBuilder(builder *bytes.Buffer, r []rune) {
	for _, rr := range r {
		builder.WriteRune(rr)
	}
}

func LogoType(writer *bytes.Buffer, input []rune) {
	const (
		지금은시작 = iota
		지금은초성
		지금은중성
		지금은종성
	)

	t := make([]rune, 0, 9)
	지금은 := 지금은시작
	합쳐진문자열 := 겹받침합치기(input)
	for _, r := range 합쳐진문자열 {
		switch 지금은 {
		case 지금은시작:
			if 이건자음인가(r) {
				지금은 = 지금은초성
				t = append(t, r)
			} else if 이건모음인가(r) {
				writeRuneToBuilder(writer, t)
				t = append(t[:0], r)
				지금은 = 지금은중성
			}
		case 지금은초성:
			if 이건모음인가(r) {
				지금은 = 지금은중성
				t = append(t, r)
			} else if 이건자음인가(r) {
				writeRuneToBuilder(writer, t)
				t = append(t[:0], r)
				지금은 = 지금은초성
			}
		case 지금은중성:
			if 이건자음인가(r) {
				지금은 = 지금은종성
				t = append(t, r)
			} else if 이건모음인가(r) {
				writeRuneToBuilder(writer, t)
				t = append(t[:0], r)
				지금은 = 지금은중성
			}
		case 지금은종성:
			if 이건자음인가(r) {
				t = append(t, r)
				writeRuneToBuilder(writer, t)
				t = t[:0]
				지금은 = 지금은시작
			} else if 이건모음인가(r) {
				writeRuneToBuilder(writer, t)
				t = append(t[:0], r)
				지금은 = 지금은중성
			}
		}
	}
	if len(t) > 0 {
		writeRuneToBuilder(writer, t)
	}
}
