package korean_romanizer_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUnicodeFinal(t *testing.T) {
	runes := []rune{' ', 'ᆨ', 'ᆩ', 'ᆪ', 'ᆫ', 'ᆬ', 'ᆭ', 'ᆮ', 'ᆯ', 'ᆰ', 'ᆱ', 'ᆲ', 'ᆳ', 'ᆴ', 'ᆵ', 'ᆶ', 'ᆷ', 'ᆸ', 'ᆹ', 'ᆺ', 'ᆻ', 'ᆼ', 'ᆽ', 'ᆾ', 'ᆿ', 'ᇀ', 'ᇁ', 'ᇂ'}
	k := GetUnicodeFinal()
	assert.Equal(t, k, runes)
	k = GetUnicodeFinal()
	assert.Equal(t, k, runes)
}

func TestIsHangul(t *testing.T) {
	hangul := '안'
	assert.Equal(t, IsHangul(hangul), true)
}

func TestSeparateSyllable(t *testing.T) {
	hangul := '안'
	isHangul, initial, medial, final := SeparateSyllable(hangul)
	assert.Equal(t, isHangul, true)
	assert.Equal(t, initial, 11)
	assert.Equal(t, medial, 0)
	assert.Equal(t, final, 4)
}

func TestNewSyllable1(t *testing.T) {
	hangul := '안'
	s := NewSyllable(hangul)
	assert.Equal(t, s.initial, 'ᄋ')
	assert.Equal(t, s.medial, 'ㅏ')
	assert.Equal(t, s.final, 'ᆫ')
}

func TestNewSyllable2(t *testing.T) {
	hangul := '좋'
	s := NewSyllable(hangul)
	assert.Equal(t, s.initial, 'ᄌ')
	assert.Equal(t, s.medial, 'ㅗ')
	assert.Equal(t, s.final, 'ᇂ')
}

func TestSyllable_ConstructSyllable1(t *testing.T) {
	hangul := '안'
	s := NewSyllable(hangul)
	syllabi := s.ConstructSyllable(s.initial, s.medial, s.final)
	assert.Equal(t, syllabi, hangul)
}

func TestSyllable_ConstructSyllable2(t *testing.T) {
	hangul := '좋'
	s := NewSyllable(hangul)
	syllabi := s.ConstructSyllable(s.initial, s.medial, s.final)
	assert.Equal(t, syllabi, hangul)
}
