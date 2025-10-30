package hangul

import (
	"bytes"
)

type KeyboardLayout map[rune]rune

var (
	DubeolsikLayout = KeyboardLayout{
		'q': 'ㅂ', 'Q': 'ㅃ',
		'w': 'ㅈ', 'W': 'ㅉ',
		'e': 'ㄷ', 'E': 'ㄸ',
		'r': 'ㄱ', 'R': 'ㄲ',
		't': 'ㅅ', 'T': 'ㅆ',
		'y': 'ㅛ', 'Y': 'ㅛ',
		'u': 'ㅕ', 'U': 'ㅕ',
		'i': 'ㅑ', 'I': 'ㅑ',
		'o': 'ㅐ', 'O': 'ㅒ',
		'p': 'ㅔ', 'P': 'ㅖ',
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
	SebulshikFinalLayout = KeyboardLayout{
		'!':  'ㄲ',
		'#':  'ㅈ',
		'$':  'ㄿ',
		'%':  'ㄾ',
		'\'': 'ㅌ',
		'(':  '\'',
		')':  '~',
		'+':  '+',
		',':  ',',
		'-':  ')',
		'.':  '.',
		'/':  'ㅗ',
		'0':  'ㅋ',
		'1':  'ㅎ',
		'2':  'ㅆ',
		'3':  'ㅂ',
		'4':  'ㅛ',
		'5':  'ㅠ',
		'6':  'ㅑ',
		'7':  'ㅖ',
		'8':  'ㅢ',
		'9':  'ㅜ',
		':':  '4',
		';':  'ㅂ',
		'<':  ',',
		'=':  '>',
		'>':  '.',
		'?':  '!',
		'@':  'ㄺ',
		'A':  'ㄷ',
		'B':  '?',
		'C':  'ㅋ',
		'D':  'ㄼ',
		'E':  'ㄵ',
		'F':  'ㄻ',
		'G':  'ㅒ',
		'H':  '0',
		'I':  '7',
		'J':  '1',
		'K':  '2',
		'L':  '3',
		'M':  '"',
		'N':  '-',
		'O':  '8',
		'P':  '9',
		'Q':  'ㅍ',
		'R':  'ㅀ',
		'S':  'ㄶ',
		'T':  'ㄽ',
		'U':  '6',
		'V':  'ㄳ',
		'W':  'ㅌ',
		'X':  'ㅄ',
		'Y':  '5',
		'Z':  'ㅊ',
		'[':  '(',
		'\\': ':',
		']':  '<',
		'^':  '=',
		'_':  ';',
		'`':  '*',
		'a':  'ㅇ',
		'b':  'ㅜ',
		'c':  'ㅔ',
		'd':  'ㅣ',
		'e':  'ㅕ',
		'f':  'ㅏ',
		'g':  'ㅡ',
		'h':  'ㄴ',
		'i':  'ㅁ',
		'j':  'ㅇ',
		'k':  'ㄱ',
		'l':  'ㅈ',
		'm':  'ㅎ',
		'n':  'ㅅ',
		'o':  'ㅊ',
		'p':  'ㅍ',
		'q':  'ㅅ',
		'r':  'ㅐ',
		's':  'ㄴ',
		't':  'ㅓ',
		'u':  'ㄷ',
		'v':  'ㅗ',
		'w':  'ㄹ',
		'x':  'ㄱ',
		'y':  'ㄹ',
		'z':  'ㅁ',
		'{':  '%',
		'|':  '\\',
		'}':  '/',
	}
)

var defaultKeyboardLayout = DubeolsikLayout

func 이건영어인가(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		return true
	}
	return false
}

func layoutOrDefault(layout KeyboardLayout) KeyboardLayout {
	if layout == nil {
		return defaultKeyboardLayout
	}
	return layout
}

func layoutLookup(layout KeyboardLayout, r rune) (rune, bool) {
	layout = layoutOrDefault(layout)
	if layout == nil {
		return 0, false
	}
	han, ok := layout[r]
	return han, ok
}

func 이건모음인가(r rune) bool {
	return 이건모음인가WithLayout(r, nil)
}

func 이건모음인가WithLayout(r rune, layout KeyboardLayout) bool {
	if han, ok := layoutLookup(layout, r); ok {
		switch han {
		case 'ㅑ', 'ㅕ', 'ㅛ', 'ㅠ', 'ㅏ', 'ㅓ', 'ㅗ', 'ㅜ', 'ㅡ', 'ㅣ', 'ㅐ', 'ㅔ', 'ㅒ', 'ㅖ',
			'ㅘ', 'ㅙ', 'ㅚ', 'ㅝ', 'ㅞ', 'ㅟ', 'ㅢ':
			return true
		default:
			return false
		}
	}
	switch r {
	case 'ㅑ', 'ㅕ', 'ㅛ', 'ㅠ', 'ㅏ', 'ㅓ', 'ㅗ', 'ㅜ', 'ㅡ', 'ㅣ', 'ㅐ', 'ㅔ', 'ㅒ', 'ㅖ',
		'ㅘ', 'ㅙ', 'ㅚ', 'ㅝ', 'ㅞ', 'ㅟ', 'ㅢ':
		return true
	default:
		return false
	}
}

func 이건자음인가(r rune) bool {
	return 이건자음인가WithLayout(r, nil)
}

func 이건자음인가WithLayout(r rune, layout KeyboardLayout) bool {
	if han, ok := layoutLookup(layout, r); ok {
		switch han {
		case 'ㄱ', 'ㄴ', 'ㄷ', 'ㄹ', 'ㅁ', 'ㅂ', 'ㅅ', 'ㅇ', 'ㅈ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ', 'ㄲ', 'ㄸ', 'ㅃ', 'ㅆ', 'ㅉ', 'ㄳ', 'ㄵ', 'ㄶ', 'ㄺ', 'ㄻ', 'ㄼ', 'ㄽ', 'ㄾ', 'ㄿ', 'ㅀ', 'ㅄ':
			return true
		default:
			return false
		}
	}
	switch r {
	case 'ㄱ', 'ㄴ', 'ㄷ', 'ㄹ', 'ㅁ', 'ㅂ', 'ㅅ', 'ㅇ', 'ㅈ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ', 'ㄲ', 'ㄸ', 'ㅃ', 'ㅆ', 'ㅉ', 'ㄳ', 'ㄵ', 'ㄶ', 'ㄺ', 'ㄻ', 'ㄼ', 'ㄽ', 'ㄾ', 'ㄿ', 'ㅀ', 'ㅄ':
		return true
	default:
		return false
	}
}

func 이건한글인가(r rune) bool {
	return 이건한글인가WithLayout(r, nil)
}

func 이건한글인가WithLayout(r rune, layout KeyboardLayout) bool {
	if 이건모음인가WithLayout(r, layout) || 이건자음인가WithLayout(r, layout) {
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

var doubleVowels = map[[2]rune]rune{
	{'ㅗ', 'ㅏ'}: 'ㅘ',
	{'ㅗ', 'ㅐ'}: 'ㅙ',
	{'ㅗ', 'ㅣ'}: 'ㅚ',
	{'ㅜ', 'ㅓ'}: 'ㅝ',
	{'ㅜ', 'ㅔ'}: 'ㅞ',
	{'ㅜ', 'ㅣ'}: 'ㅟ',
	{'ㅡ', 'ㅣ'}: 'ㅢ',
}

var 겹받침분리 = map[rune][2]rune{
	'ㄳ': {'ㄱ', 'ㅅ'},
	'ㄵ': {'ㄴ', 'ㅈ'},
	'ㄶ': {'ㄴ', 'ㅎ'},
	'ㄺ': {'ㄹ', 'ㄱ'},
	'ㄻ': {'ㄹ', 'ㅁ'},
	'ㄼ': {'ㄹ', 'ㅂ'},
	'ㄽ': {'ㄹ', 'ㅅ'},
	'ㄾ': {'ㄹ', 'ㅌ'},
	'ㄿ': {'ㄹ', 'ㅍ'},
	'ㅀ': {'ㄹ', 'ㅎ'},
	'ㅄ': {'ㅂ', 'ㅅ'},
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

func 겹자합치기(input []rune) []rune {
	result := make([]rune, 0, len(input))
	for i := 0; i < len(input); i++ {
		if i < len(input)-1 {
			pair := [2]rune{input[i], input[i+1]}
			if dc, ok := doubleConsonants[pair]; ok {
				result = append(result, dc)
				i++
				continue
			}
			if dv, ok := doubleVowels[pair]; ok {
				result = append(result, dv)
				i++
				continue
			}
		}
		result = append(result, input[i])
	}
	return result
}

func writeRuneToBuilder(builder *bytes.Buffer, layout KeyboardLayout, r []rune) {
	if len(r) == 0 {
		return
	}

	layout = layoutOrDefault(layout)

	if 이건한글인가WithLayout(r[0], layout) {
		switch len(r) {
		case 1:
			builder.WriteRune(r[0])
			return
		case 2:
			조합 := [2]bool{이건자음인가WithLayout(r[0], layout), 이건모음인가WithLayout(r[1], layout)}
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
			조합 := [3]bool{
				이건자음인가WithLayout(r[0], layout),
				이건모음인가WithLayout(r[1], layout),
				이건자음인가WithLayout(r[2], layout),
			}
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
	logoTypeWithLayout(writer, input, nil)
}

func LogoTypeWithLayout(writer *bytes.Buffer, input []rune, layout KeyboardLayout) {
	logoTypeWithLayout(writer, input, layout)
}

func logoTypeWithLayout(writer *bytes.Buffer, input []rune, layout KeyboardLayout) {
	const (
		지금은시작 = iota
		지금은초성
		지금은중성
		지금은종성
	)

	layout = layoutOrDefault(layout)

	t := make([]rune, 0, 9)
	지금은 := 지금은시작
	합쳐진문자열 := 겹자합치기(input)
	for _, r := range 합쳐진문자열 {
		switch 지금은 {
		case 지금은시작:
			if 이건자음인가WithLayout(r, layout) {
				writeRuneToBuilder(writer, layout, t)
				t = append(t[:0], r)
				지금은 = 지금은초성
			} else if 이건모음인가WithLayout(r, layout) {
				writeRuneToBuilder(writer, layout, t)
				t = append(t[:0], r)
				지금은 = 지금은중성
			} else {
				t = append(t, r)
				지금은 = 지금은시작
			}
		case 지금은초성:
			if 이건모음인가WithLayout(r, layout) {
				지금은 = 지금은중성
				t = append(t, r)
			} else if 이건자음인가WithLayout(r, layout) {
				writeRuneToBuilder(writer, layout, t)
				t = append(t[:0], r)
				지금은 = 지금은초성
			} else {
				writeRuneToBuilder(writer, layout, t)
				t = append(t[:0], r)
				지금은 = 지금은시작
			}
		case 지금은중성:
			if 이건자음인가WithLayout(r, layout) {
				지금은 = 지금은종성
				t = append(t, r)
			} else if 이건모음인가WithLayout(r, layout) {
				writeRuneToBuilder(writer, layout, t)
				t = append(t[:0], r)
				지금은 = 지금은중성
			} else {
				writeRuneToBuilder(writer, layout, t)
				t = append(t[:0], r)
				지금은 = 지금은시작
			}
		case 지금은종성:
			if 이건자음인가WithLayout(r, layout) {
				writeRuneToBuilder(writer, layout, t)
				t = append(t[:0], r)
				지금은 = 지금은초성
			} else if 이건모음인가WithLayout(r, layout) {
				종성 := t[len(t)-1]
				if 분리됨, ok := 겹받침분리[종성]; ok {
					t[len(t)-1] = 분리됨[0]
					writeRuneToBuilder(writer, layout, t)
					t = append(t[:0], 분리됨[1], r)
					지금은 = 지금은중성
				} else {
					writeRuneToBuilder(writer, layout, t[:len(t)-1])
					t = append(t[:0], 종성, r)
					지금은 = 지금은중성
				}
			} else {
				writeRuneToBuilder(writer, layout, t)
				t = append(t[:0], r)
				지금은 = 지금은시작
			}
		}
	}
	if len(t) > 0 {
		writeRuneToBuilder(writer, layout, t)
	}
}
