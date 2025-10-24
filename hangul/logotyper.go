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
		LogoType(lt.result, 겹받침합치기(lt.buffer[:len(lt.buffer)-1])) // Process without the break word
		lt.result.WriteRune(r)                                    // Write the break word
		lt.buffer = lt.buffer[:0]                                 // Reset buffer
	}
}

func (lt *LogoTyper) Result() []byte {
	LogoType(lt.result, 겹받침합치기(lt.buffer))
	return lt.result.Bytes()
}
