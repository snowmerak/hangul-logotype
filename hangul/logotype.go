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

func 이건한글인가(r rune) bool {
	if 이건모음인가(r) || 이건자음인가(r) {
		return true
	}
	return false
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

var 초성매핑 = map[rune]int{
	'ㄱ': 0, 'ㄲ': 1, 'ㄴ': 2, 'ㄷ': 3, 'ㄸ': 4,
	'ㄹ': 5, 'ㅁ': 6, 'ㅂ': 7, 'ㅃ': 8, 'ㅅ': 9,
	'ㅆ': 10, 'ㅇ': 11, 'ㅈ': 12, 'ㅉ': 13, 'ㅊ': 14,
	'ㅋ': 15, 'ㅌ': 16, 'ㅍ': 17, 'ㅎ': 18,
}

var 중성매핑 = map[rune]int{
	'ㅏ': 0, 'ㅐ': 1, 'ㅑ': 2, 'ㅒ': 3, 'ㅓ': 4,
	'ㅔ': 5, 'ㅕ': 6, 'ㅖ': 7, 'ㅗ': 8, 'ㅘ': 9,
	'ㅙ': 10, 'ㅚ': 11, 'ㅛ': 12, 'ㅜ': 13, 'ㅝ': 14,
	'ㅞ': 15, 'ㅟ': 16, 'ㅠ': 17, 'ㅡ': 18, 'ㅢ': 19,
	'ㅣ': 20,
}

var 종성매핑 = map[rune]int{
	'ㄱ': 1, 'ㄲ': 2, 'ㄳ': 3, 'ㄴ': 4, 'ㄵ': 5,
	'ㄶ': 6, 'ㄷ': 7, 'ㄹ': 8, 'ㄺ': 9, 'ㄻ': 10,
	'ㄼ': 11, 'ㄽ': 12, 'ㄾ': 13, 'ㄿ': 14, 'ㅀ': 15,
	'ㅁ': 16, 'ㅂ': 17, 'ㅄ': 18, 'ㅅ': 19, 'ㅆ': 20,
	'ㅇ': 21, 'ㅈ': 22, 'ㅊ': 23, 'ㅋ': 24, 'ㅌ': 25,
	'ㅍ': 26, 'ㅎ': 27,
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
	if len(r) == 0 {
		return
	}

	if 이건한글인가(r[0]) {
		switch len(r) {
		case 1:
			builder.WriteRune(r[0])
			return
		case 2:
			조합 := [2]bool{이건자음인가(r[0]), 이건모음인가(r[1])}
			switch 조합 {
			case [2]bool{true, true}:
				초성idx, 초성존재 := 초성매핑[r[0]]
				중성idx, 중성존재 := 중성매핑[r[1]]
				if 초성존재 && 중성존재 {
					합성자 := 0xAC00 + 초성idx*21*28 + 중성idx*28
					builder.WriteRune(rune(합성자))
					return
				}
			}
		case 3:
			조합 := [3]bool{이건자음인가(r[0]), 이건모음인가(r[1]), 이건자음인가(r[2])}
			switch 조합 {
			case [3]bool{true, true, true}:
				초성idx, 초성존재 := 초성매핑[r[0]]
				중성idx, 중성존재 := 중성매핑[r[1]]
				종성idx, 종성존재 := 종성매핑[r[2]]
				if 초성존재 && 중성존재 && 종성존재 {
					합성자 := 0xAC00 + 초성idx*21*28 + 중성idx*28 + 종성idx
					builder.WriteRune(rune(합성자))
					return
				}
			}
		}
	}

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
				writeRuneToBuilder(writer, t)
				t = append(t[:0], r)
				지금은 = 지금은초성
			} else if 이건모음인가(r) {
				writeRuneToBuilder(writer, t)
				t = append(t[:0], r)
				지금은 = 지금은중성
			} else {
				t = append(t, r)
				지금은 = 지금은시작
			}
		case 지금은초성:
			if 이건모음인가(r) {
				지금은 = 지금은중성
				t = append(t, r)
			} else if 이건자음인가(r) {
				writeRuneToBuilder(writer, t)
				t = append(t[:0], r)
				지금은 = 지금은초성
			} else {
				writeRuneToBuilder(writer, t)
				t = append(t[:0], r)
				지금은 = 지금은시작
			}
		case 지금은중성:
			if 이건자음인가(r) {
				지금은 = 지금은종성
				t = append(t, r)
			} else if 이건모음인가(r) {
				writeRuneToBuilder(writer, t)
				t = append(t[:0], r)
				지금은 = 지금은중성
			} else {
				writeRuneToBuilder(writer, t)
				t = append(t[:0], r)
				지금은 = 지금은시작
			}
		case 지금은종성:
			if 이건자음인가(r) {
				writeRuneToBuilder(writer, t)
				t = append(t[:0], r)
				지금은 = 지금은초성
			} else if 이건모음인가(r) {
				// 종성 자음을 빼서 다음 음절의 초성으로 만듦
				writeRuneToBuilder(writer, t[:len(t)-1])
				t = append(t[:0], t[len(t)-1], r)
				지금은 = 지금은중성
			} else {
				writeRuneToBuilder(writer, t)
				t = append(t[:0], r)
				지금은 = 지금은시작
			}
		}
	}
	if len(t) > 0 {
		writeRuneToBuilder(writer, t)
	}
}
