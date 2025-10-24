package hangul

import "bytes"

type LogoTyper struct {
	buffer     []rune
	result     *bytes.Buffer
	breakWords map[rune]struct{}
}

func NewLogoTyper() *LogoTyper {
	return &LogoTyper{
		buffer:     make([]rune, 0, 9),
		result:     bytes.NewBuffer(nil),
		breakWords: DefaultBreakWords(),
	}
}

func (lt *LogoTyper) WithBreakWords(breakWords map[rune]struct{}) *LogoTyper {
	lt.breakWords = breakWords
	return lt
}

func (lt *LogoTyper) Reset() {
	lt.result.Reset()
	lt.buffer = make([]rune, 0, 9)
}

func (lt *LogoTyper) WriteRune(r rune) {
	lt.buffer = append(lt.buffer, r)
	if _, ok := lt.breakWords[r]; ok {
	}
}
