package korean_romanizer_go

var doubleConsonantFinal = map[rune]string{
	'ㄳ': "ㄱㅅ",
	'ㄵ': "ᆫㅈ",
	'ᆭ': "ᆫᇂ",
	'ㄺ': "ㄹㄱ",
	'ㄻ': "ㄹㅁ",
	'ㄼ': "ㄹㅂ",
	'ㄽ': "ㄹㅅ",
	'ㄾ': "ㄹㅌ",
	'ㄿ': "ㄹㅍ",
	'ㅀ': "ㄹᇂ",
	'ㅄ': "ㅂㅅ",
	'ㅆ': "ㅅㅅ",
}

const NullConsonant = 'ᄋ'

type Pronouncer struct {
	syllables  []*Syllable
	pronounced string
}

func NewPronouncer(text string) *Pronouncer {
	p := &Pronouncer{}
	var syllables []*Syllable

	for _, runeValue := range text {
		newSyllable := NewSyllable(runeValue)
		syllables = append(syllables, newSyllable)

	}
	p.syllables = syllables

	pronounced := p.FinalSubstitute()
	pronouncedString := ""
	for i := range pronounced {

		pronouncedString += pronounced[i].ToString()
	}
	p.pronounced = pronouncedString
	return p
}

func (p *Pronouncer) FinalSubstitute() []*Syllable {
	var newSyllables []*Syllable
	for idx := range p.syllables {
		syllable := p.syllables[idx]

		var nextSyllable *Syllable
		if idx+1 < len(p.syllables) {
			nextSyllable = p.syllables[idx+1]
		}

		finalIsBeforeC := false
		if nextSyllable != nil {
			finalIsBeforeC = (syllable.final != ' ') && (nextSyllable.initial != ' ' && nextSyllable.initial != NullConsonant)
		}

		finalIsBeforeV := false
		if nextSyllable != nil {
			finalIsBeforeV = syllable.final != ' ' && (nextSyllable.initial == ' ' || nextSyllable.initial == NullConsonant)
		}

		// 1. 받침 ‘ㄲ, ㅋ’, ‘ㅅ, ㅆ, ㅈ, ㅊ, ㅌ’, ‘ㅍ’은 어말 또는 자음 앞에서 각각 대표음 [ㄱ, ㄷ, ㅂ]으로 발음한다.
		// 2. 겹받침 ‘ㄳ’, ‘ㄵ’, ‘ㄼ, ㄽ, ㄾ’, ‘ㅄ’은 어말 또는 자음 앞에서 각각 [ㄱ, ㄴ, ㄹ, ㅂ]으로 발음한다.
		// 3. 겹받침 ‘ㄺ, ㄻ, ㄿ’은 어말 또는 자음 앞에서 각각 [ㄱ, ㅁ, ㅂ]으로 발음한다.
		// <-> 단, 국어의 로마자 표기법 규정에 의해 된소리되기는 표기에 반영하지 않으므로 제외.

		if syllable.final != ' ' || finalIsBeforeC {
			if IndexRune([]rune{'ᆩ', 'ᆿ', 'ᆪ', 'ᆰ'}, syllable.final) != -1 {
				syllable.final = 'ᆨ'
			} else if IndexRune([]rune{'ᆺ', 'ᆻ', 'ᆽ', 'ᆾ', 'ᇀ'}, syllable.final) != -1 {
				syllable.final = 'ᆮ'
			} else if IndexRune([]rune{'ᇁ', 'ᆹ', 'ᆵ'}, syllable.final) != -1 {
				syllable.final = 'ᆸ'
			} else if syllable.final == 'ᆬ' {
				syllable.final = 'ᆫ'
			} else if IndexRune([]rune{'ᆲ', 'ᆳ', 'ᆴ'}, syllable.final) != -1 {
				syllable.final = 'ᆯ'
			} else if syllable.final == 'ᆱ' {
				syllable.final = 'ᆷ'
			}
		}

		//  4. 받침 ‘ㅎ’의 발음은 다음과 같다.
		if IndexRune([]rune{'ᇂ', 'ᆭ', 'ᆶ'}, syllable.final) != -1 {
			if nextSyllable != nil {
				// ‘ㅎ(ㄶ, ㅀ)’ 뒤에 ‘ㄱ, ㄷ, ㅈ’이 결합되는 경우에는, 뒤 음절 첫소리와 합쳐서 [ㅋ, ㅌ, ㅊ]으로 발음한다.
				// ‘ㅎ(ㄶ, ㅀ)’ 뒤에 ‘ㅅ’이 결합되는 경우에는, ‘ㅅ’을 [ㅆ]으로 발음한다.
				if IndexRune([]rune{'ᄀ', 'ᄃ', 'ᄌ', 'ᄉ'}, nextSyllable.final) != -1 {
					changeTo := map[rune]rune{'ᄀ': 'ᄏ', 'ᄃ': 'ᄐ', 'ᄌ': 'ᄎ', 'ᄉ': 'ᄊ'}
					syllable.final = ' '
					nextSyllable.initial = changeTo[nextSyllable.initial]
				} else if nextSyllable.initial == 'ᄂ' {
					// 3. ‘ㅎ’ 뒤에 ‘ㄴ’이 결합되는 경우에는, [ㄴ]으로 발음한다.
					// TODO: [붙임] ‘ㄶ, ㅀ’ 뒤에 ‘ㄴ’이 결합되는 경우에는, ‘ㅎ’을 발음하지 않는다.
					if IndexRune([]rune{'ᆭ', 'ᆶ'}, syllable.final) != -1 {
						if syllable.final == 'ᆭ' {
							syllable.final = 'ᆫ'
						} else if syllable.final == 'ᆶ' {
							syllable.final = 'ᆯ'
						} else {
							syllable.final = 'ᆫ'
						}
					}
				} else if nextSyllable.initial == NullConsonant {
					// 4. ‘ㅎ(ㄶ, ㅀ)’ 뒤에 모음으로 시작된 어미나 접미사가 결합되는 경우에는,
					// ‘ㅎ’을 발음하지 않는다
					if IndexRune([]rune{'ᆭ', 'ᆶ'}, syllable.final) != -1 {
						if syllable.final == 'ᆭ' {
							syllable.final = 'ᆫ'
						} else if syllable.final == 'ᆶ' {
							syllable.final = 'ᆯ'
						}
					} else {
						syllable.final = ' '
					}
				} else {
					if syllable.final == 'ᇂ' {
						syllable.final = ' '
					}
				}
			} else {
				if syllable.final == 'ᇂ' {
					syllable.final = ' '
				}
			}
		}
		// 5. 홑받침이나 쌍받침이 모음으로 시작된 조사나 어미, 접미사와 결합되는 경우에는,
		// 제 음가대로 뒤 음절 첫소리로 옮겨 발음한다.
		if nextSyllable != nil && finalIsBeforeV {
			if nextSyllable.initial == NullConsonant && (syllable.final != 'ᆼ' && syllable.final != ' ') {
				nextSyllable.initial = FinalToInitial(syllable.final)
				syllable.final = ' '
			}
		}
		// fmt.Println("F", syllable.ToString(), nextSyllable.ToString())

		// 6. 겹받침이 모음으로 시작된 조사나 어미, 접미사와 결합되는 경우에는,
		// 뒤엣것만을 뒤 음절 첫소리로 옮겨 발음한다.(이 경우, ‘ㅅ’은 된소리로 발음함.)
		if val, ok := doubleConsonantFinal[syllable.final]; ok {
			syllable.final = rune(val[0])
			nextSyllable.initial = FinalToInitial(rune(val[1]))
		}
		if nextSyllable != nil {
			p.syllables[idx+1] = nextSyllable
		}
		newSyllables = append(newSyllables, syllable)
	}
	return newSyllables
}
