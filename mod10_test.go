package mod10_test

import (
	"testing"

	"github.com/bycrates/mod10"
	"github.com/stretchr/testify/assert"
)

// Checks against a known valid list of results
func TestStaticDigitsCheck(t *testing.T) {
	pairs := map[string]string{
		"299812919":       "2998129197",
		"497871337":       "4978713370",
		"495751699":       "4957516992",
		"1054169822":      "10541698220",
		"36332802471":     "363328024718",
		"142290000136281": "1422900001362812",
	}

	for withoutControlbit, withControlbit := range pairs {
		added, err := mod10.AddControlBitString(withoutControlbit)
		assert.Nil(t, err)

		checked, _ := mod10.CheckString(added)
		assert.Nil(t, err)

		assert.True(t, checked)
		assert.Equal(t, withControlbit, added)
	}
}

// Check that the generated string with added
// control bit is valid against itself
func TestGenerateCheckItself(t *testing.T) {
	added := mod10.AddControlBit(100515)
	checked, err := mod10.CheckString(added)

	assert.NotEmpty(t, added)
	assert.Nil(t, err)
	assert.True(t, checked)
}
