package korean_romanizer_go

import "testing"
import "github.com/stretchr/testify/assert"

func TestNewRomanizer(t *testing.T) {
	text := "안녕하세요"
	r := NewRomanizer(text)
	if r.text != text {
		t.Errorf("%s doesn't match %s", r.text, text)
	}
}

func TestRomanizer_Romanize(t *testing.T) {
	text := "안녕하세요"
	r := NewRomanizer(text)
	romanized, err := r.Romanize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, romanized, "annyeonghaseyo")
}

func TestRomanizer_Romanize_SpacedText(t *testing.T) {
	text := "아이유 방탄소년단"
	r := NewRomanizer(text)
	romanized, err := r.Romanize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, romanized, "aiyu bangtansonyeondan")
}

func TestRomanizer_Romanize_OnsetGDB1(t *testing.T) {
	text := "구미"
	r := NewRomanizer(text)
	romanized, err := r.Romanize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, romanized, "gumi")
}

func TestRomanizer_Romanize_OnsetGDB2(t *testing.T) {
	text := "영동"
	r := NewRomanizer(text)
	romanized, err := r.Romanize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, romanized, "yeongdong")
}

func TestRomanizer_Romanize_OnsetGDB3(t *testing.T) {
	text := "한밭"
	r := NewRomanizer(text)
	romanized, err := r.Romanize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, romanized, "hanbat")
}

func TestRomanizer_Romanize_CodaGDB1(t *testing.T) {
	text := "밝다"
	r := NewRomanizer(text)
	romanized, err := r.Romanize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, romanized, "bakda")
}

func TestRomanizer_Romanize_CodaGDB2(t *testing.T) {
	text := "바닷가"
	r := NewRomanizer(text)
	romanized, err := r.Romanize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, romanized, "badatga")
}

func TestRomanizer_Romanize_CodaGDB3(t *testing.T) {
	text := "없다"
	r := NewRomanizer(text)
	romanized, err := r.Romanize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, romanized, "eopda")
}

func TestRomanizer_Romanize_CodaGDB4(t *testing.T) {
	text := "앞만"
	r := NewRomanizer(text)
	romanized, err := r.Romanize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, romanized, "apman")
}

func TestRomanizer_Romanize_RL1(t *testing.T) {
	text := "구리"
	r := NewRomanizer(text)
	romanized, err := r.Romanize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, romanized, "guri")
}

func TestRomanizer_Romanize_RL2(t *testing.T) {
	text := "설악"
	r := NewRomanizer(text)
	romanized, err := r.Romanize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, romanized, "seorak")
}

func TestRomanizer_Romanize_NextSyllableNullInitial1(t *testing.T) {
	text := "강약"
	r := NewRomanizer(text)
	romanized, err := r.Romanize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, romanized, "gangyak")
}

func TestRomanizer_Romanize_NextSyllableNullInitial2(t *testing.T) {
	text := "강원"
	r := NewRomanizer(text)
	romanized, err := r.Romanize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, romanized, "gangwon")
}

func TestRomanizer_Romanize_NextSyllableNullInitial3(t *testing.T) {
	text := "좋아하고"
	r := NewRomanizer(text)
	romanized, err := r.Romanize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, romanized, "joahago")
}

func TestRomanizer_Romanize_NextSyllableNullInitial4(t *testing.T) {
	text := "좋은"
	r := NewRomanizer(text)
	romanized, err := r.Romanize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, romanized, "joeun")
}
func TestRomanizer_Romanize_NonSyllables1(t *testing.T) {
	text := "ㅠㄴㅁㄱ"
	r := NewRomanizer(text)
	romanized, err := r.Romanize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, romanized, "ㅠㄴㅁㄱ")
}

func TestRomanizer_Romanize_NonSyllables2(t *testing.T) {
	text := "ㅠ동"
	r := NewRomanizer(text)
	romanized, err := r.Romanize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, romanized, "ㅠdong")
}

func TestRomanizer_Romanize_NonSyllables3(t *testing.T) {
	text := "apple"
	r := NewRomanizer(text)
	romanized, err := r.Romanize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, romanized, "apple")
}
