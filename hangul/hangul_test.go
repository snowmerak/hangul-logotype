package hangul

import (
	"bytes"
	"testing"
)

func Test겹받침합치기(t *testing.T) {
	tests := []struct {
		name     string
		input    []rune
		expected []rune
	}{
		{
			name:     "no double consonants",
			input:    []rune{'ㄱ', 'ㅏ'},
			expected: []rune{'ㄱ', 'ㅏ'},
		},
		{
			name:     "double consonant separate letters",
			input:    []rune{'ㄱ', 'ㄱ'},
			expected: []rune{'ㄱ', 'ㄱ'},
		},
		{
			name:     "double consonant ㄳ",
			input:    []rune{'ㄱ', 'ㅅ'},
			expected: []rune{'ㄳ'},
		},
		{
			name:     "mixed with actual double final",
			input:    []rune{'ㄱ', 'ㄱ', 'ㅏ', 'ㄹ', 'ㅁ'},
			expected: []rune{'ㄱ', 'ㄱ', 'ㅏ', 'ㄻ'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := 겹자합치기(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			for i, r := range result {
				if r != tt.expected[i] {
					t.Errorf("at index %d, expected %c, got %c", i, tt.expected[i], r)
				}
			}
		})
	}
}

func TestLogoType(t *testing.T) {
	tests := []struct {
		name     string
		input    []rune
		expected string
	}{
		{
			name:     "simple syllable",
			input:    []rune{'ㄱ', 'ㅏ'},
			expected: "가",
		},
		{
			name:     "syllable with final",
			input:    []rune{'ㄱ', 'ㅏ', 'ㄴ'},
			expected: "간",
		},
		{
			name:     "multiple syllables",
			input:    []rune{'ㄱ', 'ㅏ', 'ㄴ', 'ㅏ'},
			expected: "가나",
		},
		{
			name:     "explicit double consonant rune",
			input:    []rune{'ㄲ', 'ㅏ'},
			expected: "까",
		},
		{
			name:     "repeated consonant without shift",
			input:    []rune{'ㄱ', 'ㄱ', 'ㅏ'},
			expected: "ㄱ가",
		},
		{
			name:     "word: 한글 (hangul)",
			input:    []rune{'ㅎ', 'ㅏ', 'ㄴ', 'ㄱ', 'ㅡ', 'ㄹ'},
			expected: "한글",
		},
		{
			name:     "word: 사랑 (love)",
			input:    []rune{'ㅅ', 'ㅏ', 'ㄹ', 'ㅏ', 'ㅇ'},
			expected: "사랑",
		},
		{
			name:     "word: 컴퓨터 (computer)",
			input:    []rune{'ㅋ', 'ㅓ', 'ㅁ', 'ㅍ', 'ㅠ', 'ㅌ', 'ㅓ'},
			expected: "컴퓨터",
		},
		{
			name:     "word with double final: 닭 (chicken)",
			input:    []rune{'ㄷ', 'ㅏ', 'ㄹ', 'ㄱ'},
			expected: "닭",
		},
		{
			name:     "word: 각궁 (repeated consonants)",
			input:    []rune{'ㄱ', 'ㅏ', 'ㄱ', 'ㄱ', 'ㅜ', 'ㅇ'},
			expected: "각궁",
		},
		{
			name:     "word with space: 안녕 하세요",
			input:    []rune{'ㅇ', 'ㅏ', 'ㄴ', 'ㄴ', 'ㅕ', 'ㅇ', ' ', 'ㅎ', 'ㅏ', 'ㅅ', 'ㅔ', 'ㅇ', 'ㅛ'},
			expected: "안녕 하세요",
		},
		{
			name:     "consecutive consonants after final",
			input:    []rune{'ㄱ', 'ㅏ', 'ㄴ', 'ㄱ', 'ㅏ'},
			expected: "간가",
		},
		{
			name:     "consonant only start",
			input:    []rune{'ㄱ', 'ㄴ', 'ㄷ'},
			expected: "ㄱㄴㄷ",
		},
		{
			name:     "vowel only start",
			input:    []rune{'ㅏ', 'ㅓ', 'ㅗ'},
			expected: "ㅏㅓㅗ",
		},
		{
			name:     "mixed with special chars",
			input:    []rune{'ㅎ', 'ㅏ', 'ㄴ', '.', 'ㄱ', 'ㅡ', 'ㄹ', '!'},
			expected: "한.글!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			LogoType(&buf, tt.input)
			result := buf.String()
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestLogoType_Complex(t *testing.T) {
	tests := []struct {
		name     string
		input    []rune
		expected string
	}{
		{
			name:     "double final moves to initial (달가)",
			input:    []rune{'ㄷ', 'ㅏ', 'ㄹ', 'ㄱ', 'ㅏ'},
			expected: "달가",
		},
		{
			name:     "double final preserved (닭ㄱ)",
			input:    []rune{'ㄷ', 'ㅏ', 'ㄹ', 'ㄱ', 'ㄱ'},
			expected: "닭ㄱ",
		},
		{
			name:     "complex vowel combination (ㅗ + ㅏ)",
			input:    []rune{'ㄱ', 'ㅗ', 'ㅏ'},
			expected: "과",
		},
		{
			name:     "complex vowel combination (ㅗ + ㅐ)",
			input:    []rune{'ㄱ', 'ㅗ', 'ㅐ'},
			expected: "괘",
		},
		{
			name:     "complex vowel combination (ㅗ + ㅣ)",
			input:    []rune{'ㄱ', 'ㅗ', 'ㅣ'},
			expected: "괴",
		},
		{
			name:     "complex vowel + final (완)",
			input:    []rune{'ㅇ', 'ㅗ', 'ㅏ', 'ㄴ'},
			expected: "완",
		},
		{
			name:     "complex vowel (ㅜ + ㅓ)",
			input:    []rune{'ㄷ', 'ㅜ', 'ㅓ'},
			expected: "둬",
		},
		{
			name:     "complex vowel + final (ㅜ + ㅓ + ㄴ)",
			input:    []rune{'ㄷ', 'ㅜ', 'ㅓ', 'ㄴ'},
			expected: "둰",
		},
		{
			name:     "vowel + vowel separation (과ㅣ)",
			input:    []rune{'ㄱ', 'ㅗ', 'ㅏ', 'ㅣ'},
			expected: "과ㅣ",
		},
		{
			name:     "invalid double final (각ㄷ)",
			input:    []rune{'ㄱ', 'ㅏ', 'ㄱ', 'ㄷ'},
			expected: "각ㄷ",
		},
		{
			name:     "vowel starting without initial (ㅏㄴ)",
			input:    []rune{'ㅏ', 'ㄴ'},
			expected: "ㅏㄴ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			LogoType(&buf, tt.input)
			result := buf.String()
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestLogoTyper(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic hangul input",
			input:    "ㄱㅏ ",
			expected: "가 ",
		},
		{
			name:     "english to hangul conversion",
			input:    "rk ",
			expected: "가 ",
		},
		{
			name:     "english word: gksrmf (한글)",
			input:    "gksrmf",
			expected: "한글",
		},
		{
			name:     "shift double consonant",
			input:    "Rk",
			expected: "까",
		},
		{
			name:     "lowercase repeated consonant",
			input:    "rrk",
			expected: "ㄱ가",
		},
		{
			name:     "double s without shift",
			input:    "tt",
			expected: "ㅅㅅ",
		},
		{
			name:     "english sentence: dkssud gksrmf",
			input:    "dkssud gksrmf",
			expected: "안녕 한글",
		},
		{
			name:     "mixed with special chars",
			input:    "gksrmf!",
			expected: "한글!",
		},
		{
			name:     "english: tkfkdgo (사랑해)",
			input:    "tkfkdgo",
			expected: "사랑해",
		},
		{
			name:     "long sentence: 안녕하세요 반갑습니다",
			input:    "dkssudgktpdy qksrkqtmqslek",
			expected: "안녕하세요 반갑습니다",
		},
		{
			name:     "long sentence with punctuation",
			input:    "gksrmfdms dnleogks gksrmfdl ehlqslek.",
			expected: "한글은 위대한 한글이 됩니다.",
		},
		{
			name:     "sentence: 프로그래밍은 재미있어요!",
			input:    "vmfhrmfoalddms woaldlTdjdy!",
			expected: "프로그래밍은 재미있어요!",
		},
		{
			name:     "sentence with newline: 저는 한글을 사랑합니다!",
			input:    "wjsms gksrmfdmf tkfkdgkqslek!",
			expected: "저는 한글을 사랑합니다!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lt := NewLogoTyper()
			for _, r := range tt.input {
				lt.WriteRune(r)
			}
			result := string(lt.Result())
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestLogoTyper_Complex_Dubeolshik(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "final to initial movement (닮)",
			input:    "ekfa",
			expected: "닮",
		},
		{
			name:     "double final preserved (닭)",
			input:    "ekfr",
			expected: "닭",
		},
		{
			name:     "complex vowel ㅗ + ㅏ (와)",
			input:    "dhk",
			expected: "와",
		},
		{
			name:     "complex vowel ㅗ + ㅐ (왜)",
			input:    "dho",
			expected: "왜",
		},
		{
			name:     "complex vowel ㅗ + ㅣ (외)",
			input:    "dhl",
			expected: "외",
		},
		{
			name:     "complex vowel + final (완)",
			input:    "dhks",
			expected: "완",
		},
		{
			name:     "complex vowel ㅜ + ㅓ (둬)",
			input:    "enj",
			expected: "둬",
		},
		{
			name:     "complex vowel ㅜ + ㅓ + final (둰)",
			input:    "enjs",
			expected: "둰",
		},
		{
			name:     "shift + consonant (까)",
			input:    "Rk",
			expected: "까",
		},
		{
			name:     "shift + vowel (ㅖ)",
			input:    "P",
			expected: "ㅖ",
		},
		{
			name:     "shift + vowel + final (뎨)",
			input:    "eP",
			expected: "뎨",
		},
		{
			name:     "number breaks syllable (가1)",
			input:    "rk1",
			expected: "가1",
		},
		{
			name:     "special char breaks syllable (한ㄱ!)",
			input:    "gksr!",
			expected: "한ㄱ!",
		},
		{
			name:     "repeated consonants (가가가)",
			input:    "rkrkrk",
			expected: "가가가",
		},
		{
			name:     "repeated vowels (ㅏ닫)",
			input:    "keke",
			expected: "ㅏ닫",
		},
		{
			name:     "vowel start (안녕)",
			input:    "dkssud",
			expected: "안녕",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lt := NewLogoTyper()
			lt.WriteString(tt.input)
			result := string(lt.Result())
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestLogoTyperSebulshik(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "word: 한글",
			input:    "mfskgw",
			expected: "한글",
		},
		{
			name:     "word: 세벌식",
			input:    "nc;twndx",
			expected: "세벌식",
		},
		{
			name:     "final consonant with digit",
			input:    "hv1",
			expected: "놓",
		},
		{
			name:     "double final consonant",
			input:    "kfX",
			expected: "값",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lt := NewLogoTyperWithLayout(SebulshikFinalLayout)
			for _, r := range tt.input {
				lt.WriteRune(r)
			}
			result := string(lt.Result())
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestLogoTyperSebulshik_Complex(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "initial-medial-final (학)",
			input:    "mfk",
			expected: "학",
		},
		{
			name:     "initial-medial-final + initial (학ㄴ)",
			input:    "mfks",
			expected: "학ㄴ",
		},
		{
			name:     "initial-medial-final + medial (하가)",
			input:    "mfkf",
			expected: "하가",
		},
		{
			name:     "initial-medial-final + final (학ㅈ)",
			input:    "mfkl",
			expected: "학ㅈ",
		},
		{
			name:     "initial-medial-final + initial-medial (학느)",
			input:    "mfksg",
			expected: "학느",
		},
		{
			name:     "double final + medial (갑사)",
			input:    "kfXf",
			expected: "갑사",
		},
		{
			name:     "double final + initial (값ㄴ)",
			input:    "kfXs",
			expected: "값ㄴ",
		},
		{
			name:     "complex medial (ㄴㅇㅂ)",
			input:    "sj;",
			expected: "ㄴㅇㅂ",
		},
		{
			name:     "complex medial + final (ㄴㅇㅂㄱ)",
			input:    "sj;k",
			expected: "ㄴㅇㅂㄱ",
		},
		{
			name:     "shift + initial (ㄶㅏ)",
			input:    "Sf",
			expected: "ㄶㅏ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lt := NewLogoTyperWithLayout(SebulshikFinalLayout)
			lt.WriteString(tt.input)
			result := string(lt.Result())
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func Test이건모음인가(t *testing.T) {
	// Test with default keyboard layout mapping
	if 이건모음인가('a') { // ㅁ is consonant
		t.Error("expected false for 'a' -> ㅁ")
	}
	if !이건모음인가('o') { // ㅐ is vowel
		t.Error("expected true for 'o' -> ㅐ")
	}
	// Test with direct hangul
	if !이건모음인가('ㅏ') {
		t.Error("expected true for direct hangul vowel ㅏ")
	}
	if 이건모음인가('ㄱ') {
		t.Error("expected false for direct hangul consonant ㄱ")
	}
}

func Test이건자음인가(t *testing.T) {
	if !이건자음인가('q') { // ㅂ is consonant
		t.Error("expected true for 'q' -> ㅂ")
	}
	if 이건자음인가('o') { // ㅐ is vowel
		t.Error("expected false for 'o' -> ㅐ")
	}
}
