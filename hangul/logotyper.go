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
	if 이건영어인가(r) {
		mapped, ok := engToHan[r]
		if ok {
			r = mapped
		}
	}
	lt.buffer = append(lt.buffer, r)
	if _, ok := lt.breakWords[r]; ok {
		LogoType(lt.result, 겹자합치기(lt.buffer[:len(lt.buffer)-1]))
		lt.result.WriteRune(r)
		lt.buffer = lt.buffer[:0]
	}
}

func (lt *LogoTyper) WriteRunes(runes []rune) {
	for _, r := range runes {
		lt.WriteRune(r)
	}
}

func (lt *LogoTyper) Result() []byte {
	LogoType(lt.result, 겹자합치기(lt.buffer))
	return lt.result.Bytes()
}
