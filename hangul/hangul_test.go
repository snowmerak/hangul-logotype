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
			name:     "double consonant ㄲ",
			input:    []rune{'ㄱ', 'ㄱ'},
			expected: []rune{'ㄲ'},
		},
		{
			name:     "double consonant ㄳ",
			input:    []rune{'ㄱ', 'ㅅ'},
			expected: []rune{'ㄳ'},
		},
		{
			name:     "mixed",
			input:    []rune{'ㄱ', 'ㄱ', 'ㅏ', 'ㄹ', 'ㅁ'},
			expected: []rune{'ㄲ', 'ㅏ', 'ㄻ'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := 겹받침합치기(tt.input)
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
			expected: "간아",
		},
		{
			name:     "double consonant",
			input:    []rune{'ㄱ', 'ㄱ', 'ㅏ'},
			expected: "까",
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

func TestLogoTyper(t *testing.T) {
	lt := NewLogoTyper()
	lt.WriteRune('ㄱ')
	lt.WriteRune('ㅏ')
	lt.WriteRune(' ')
	result := string(lt.Result())
	expected := "가 "
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func Test이건모음인가(t *testing.T) {
	// Test with engToHan mapped runes
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
