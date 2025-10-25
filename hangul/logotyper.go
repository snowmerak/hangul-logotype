package hangul

import "bytes"

type LogoTyper struct {
	buffer     []rune
	result     *bytes.Buffer
	breakWords map[rune]struct{}
	layout     KeyboardLayout
}

func NewLogoTyper() *LogoTyper {
	return NewLogoTyperWithLayout(nil)
}

func NewLogoTyperWithLayout(layout KeyboardLayout) *LogoTyper {
	return &LogoTyper{
		buffer:     make([]rune, 0, 9),
		result:     bytes.NewBuffer(nil),
		breakWords: DefaultBreakWords(),
		layout:     layoutOrDefault(layout),
	}
}

func (lt *LogoTyper) WithBreakWords(breakWords map[rune]struct{}) *LogoTyper {
	lt.breakWords = breakWords
	return lt
}

func (lt *LogoTyper) WithLayout(layout KeyboardLayout) *LogoTyper {
	lt.layout = layoutOrDefault(layout)
	return lt
}

func (lt *LogoTyper) Reset() {
	lt.result.Reset()
	lt.buffer = make([]rune, 0, 9)
}

func (lt *LogoTyper) WriteRune(r rune) {
	if mapped, ok := layoutLookup(lt.layout, r); ok {
		r = mapped
	}
	lt.buffer = append(lt.buffer, r)
	if _, ok := lt.breakWords[r]; ok {
		logoTypeWithLayout(lt.result, 겹자합치기(lt.buffer[:len(lt.buffer)-1]), lt.layout)
		lt.result.WriteRune(r)
		lt.buffer = lt.buffer[:0]
	}
}

func (lt *LogoTyper) WriteRunes(runes []rune) {
	for _, r := range runes {
		lt.WriteRune(r)
	}
}

func (lt *LogoTyper) WriteString(s string) {
	for _, r := range s {
		lt.WriteRune(r)
	}
}

func (lt *LogoTyper) Result() []byte {
	logoTypeWithLayout(lt.result, 겹자합치기(lt.buffer), lt.layout)
	return lt.result.Bytes()
}
