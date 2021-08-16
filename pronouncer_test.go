package korean_romanizer_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPronouncer(t *testing.T) {
	text := "좋아하고"
	p := NewPronouncer(text)
	assert.Equal(t, p.pronounced, "조아하고")
}

func TestFinalToInitial(t *testing.T) {
	assert.Equal(t, FinalToInitial('ᇂ'), 'ᄒ')
}
