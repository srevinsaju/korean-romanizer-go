package korean_romanizer_go

const UnicodeInitialStartIndex = 4352
const UnicodeInitialStopIndex = 4371

const UnicodeFinalStartIndex = 0x11a8
const UnicodeFinalStopIndex = 0x11c3

const UnicodeOffset = 44032
const UnicodeInitialOffset = 588
const UnicodeMedialOffset = 28

var unicodeInitial []rune
var unicodeFinal []rune
var unicodeMedial = []rune{'ㅏ', 'ㅐ', 'ㅑ', 'ㅒ', 'ㅓ', 'ㅔ', 'ㅕ', 'ㅖ', 'ㅗ', 'ㅘ', 'ㅙ', 'ㅚ', 'ㅛ', 'ㅜ', 'ㅝ', 'ㅞ', 'ㅟ', 'ㅠ', 'ㅡ', 'ㅢ', 'ㅣ'}

var unicodeCompatibleInitials = []rune{'ᄀ', 'ᄁ', 'ᄂ', 'ᄃ', 'ᄄ', 'ᄅ', 'ᄆ', 'ᄇ', 'ᄈ', 'ᄉ', 'ᄊ', 'ᄋ', 'ᄌ', 'ᄍ', 'ᄎ', 'ᄏ', 'ᄐ', 'ᄑ', 'ᄒ'}
var unicodeCompatibleConsonants = []rune{'ㄱ', 'ㄲ', 'ㄴ', 'ㄷ', 'ㄸ', 'ㄹ', 'ㅁ', 'ㅂ', 'ㅃ', 'ㅅ', 'ㅆ', 'ㅇ', 'ㅈ', 'ㅉ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ'}
var unicodeCompatibleFinals = []rune{'ᆨ', 'ᆩ', 'ᆫ', 'ᆮ', '_', 'ᆯ', 'ᆷ', 'ᆸ', '_', 'ᆺ', 'ᆻ', 'ᆼ', 'ᆽ', '_', 'ᆾ', 'ᆿ', 'ᇀ', 'ᇁ', 'ᇂ'}

func GetUnicodeInitial() []rune {
	if len(unicodeInitial) == 0 {
		for i := UnicodeInitialStartIndex; i < UnicodeInitialStopIndex; i++ {
			unicodeInitial = append(unicodeInitial, rune(i))
		}
	}
	return unicodeInitial
}

func GetUnicodeFinal() []rune {
	if len(unicodeFinal) == 0 {
		unicodeFinal = append(unicodeFinal, ' ')
		for i := UnicodeFinalStartIndex; i < UnicodeFinalStopIndex; i++ {
			unicodeFinal = append(unicodeFinal, rune(i))
		}
	}
	return unicodeFinal
}

type Syllable struct {
	char    rune
	initial rune
	medial  rune
	final   rune
}

func (s *Syllable) ConstructSyllable(initial rune, medial rune, final rune) rune {
	var constructed rune
	if IsHangul(s.char) {
		i := int(initial) - 4352
		m := IndexRune(unicodeMedial, medial)
		f := 0
		if final == ' ' {
			f = 0
		} else {
			f = IndexRune(unicodeFinal, final)
		}
		constructed = rune(((i * UnicodeInitialOffset) + (m * UnicodeMedialOffset) + f) + UnicodeOffset)
	} else {
		constructed = s.char
	}
	return constructed
}

func (s *Syllable) ToString() string {
	char := s.ConstructSyllable(s.initial, s.medial, s.final)
	return string(rune(char))
}

func NewSyllable(char rune) *Syllable {
	s := &Syllable{char: char}
	isHangul, initial, medial, final := SeparateSyllable(char)
	if isHangul {
		s.initial = rune(GetUnicodeInitial()[initial])
		s.medial = rune(unicodeMedial[medial])
		s.final = rune(GetUnicodeFinal()[final])
	} else {
		s.initial = rune(initial)
		s.medial = ' '
		s.final = ' '
	}
	return s
}

func SeparateSyllable(char rune) (bool, int, int, int) {
	var (
		initial int
		medial  int
		final   int
	)
	if IsHangul(char) {
		initial = (int(char) - UnicodeOffset) / UnicodeInitialOffset
		medial = ((int(char) - UnicodeOffset) - UnicodeInitialOffset*initial) / UnicodeMedialOffset
		final = ((int(char) - UnicodeOffset) - UnicodeInitialOffset*initial) - UnicodeMedialOffset*medial
	} else {
		initial = int(char)
	}
	return IsHangul(char), initial, medial, final
}

func IsHangul(char rune) bool {
	return 0xAC00 <= int(char) && int(char) <= 0xD7A3
}

func FinalToInitial(char rune) rune {
	idx := IndexRune(unicodeCompatibleFinals, char)
	if idx == -1 {
		idx = IndexRune(unicodeCompatibleConsonants, char)
	}
	if idx == -1 {
		idx = IndexRune(unicodeCompatibleInitials, char)
	}
	if idx != -1 {
		return GetUnicodeInitial()[idx]
	} else {
		return char
	}
}
