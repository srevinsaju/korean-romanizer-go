package korean_romanizer_go

var vowel = map[rune]string{
	// 단모음 monophthongs
	'ㅏ': "a",
	'ㅓ': "eo",
	'ㅗ': "o",
	'ㅜ': "u",
	'ㅡ': "eu",
	'ㅣ': "i",
	'ㅐ': "ae",
	'ㅔ': "e",
	'ㅚ': "oe",
	'ㅟ': "wi",

	// 이중모음 diphthongs
	'ㅑ': "ya",
	'ㅕ': "yeo",
	'ㅛ': "yo",
	'ㅠ': "yu",
	'ㅒ': "yae",
	'ㅖ': "ye",
	'ㅘ': "wa",
	'ㅙ': "wae",
	'ㅝ': "wo",
	'ㅞ': "we",
	'ㅢ': "ui", //  [붙임 1] ‘ㅢ’는 ‘ㅣ’로 소리 나더라도 ‘ui’로 적는다.
}

/*
### Transcribing consonants ###
Consonants are defined in separate dicts, choseong and jongseong,
for some characters are pronounced differently depending on
its position in the syllable.
e.g. ㄱ, ㄷ, ㅂ, ㄹ are (g, d, b, r) in onset,
                  but (k, t, p, l) in coda.
e.g. ㅇ is a null sound when placed in onset, but becomes [ng] in coda.
*/

var onset = map[rune]string{
	'ᄀ': "g",
	'ᄁ': "kk",
	'ᄏ': "k",
	'ᄃ': "d",
	'ᄄ': "tt",
	'ᄐ': "t",
	'ᄇ': "b",
	'ᄈ': "pp",
	'ᄑ': "p",
	// 파찰음 affricates
	'ᄌ': "j",
	'ᄍ': "jj",
	'ᄎ': "ch",
	// 마찰음 fricatives
	'ᄉ': "s",
	'ᄊ': "ss",
	'ᄒ': "h",
	// 비음 nasals
	'ᄂ': "n",
	'ᄆ': "m",
	// 유음 liquids
	'ᄅ': "r",
	// Null sound
	'ᄋ': "",
}

/*

종성 Jongseong (Syllable Coda)
"The 7 Jongseongs (7종성)"
Only the seven consonants below may appear in coda position
*/

var coda = map[rune]string{
	// 파열음 stops/plosives
	'ᆨ': "k",
	'ᆮ': "t",
	'ᆸ': "p",
	// 비음 nasals
	'ᆫ': "n",
	'ᆼ': "ng",
	'ᆷ': "m",
	// 유음 liquids
	'ᆯ': "l",
}

type Romanizer struct {
	text string
}

func NewRomanizer(text string) Romanizer {
	return Romanizer{text: text}
}

func (r *Romanizer) Romanize() (string, error) {
	pronounced := NewPronouncer(r.text).pronounced
	// hangul := "[가-힣ㄱ-ㅣ]"
	romanized := ""
	for _, runeValue := range pronounced {
		if IsHangul(runeValue) {
			s := NewSyllable(runeValue)
			if (s.medial == ' ' || s.medial == 0) && (s.final == ' ' || s.final == 0) {
				romanized += string(runeValue)
			} else {
				romanized += onset[s.initial] + vowel[s.medial] + coda[s.final]
			}
		} else {
			romanized += string(runeValue)
		}
	}
	return romanized, nil
}
