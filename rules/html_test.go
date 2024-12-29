package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTML(t *testing.T) {
	t.Run("Valid HTML with tags", func(t *testing.T) {
		err := HTML("<p>This is a paragraph</p>")
		assert.NoError(t, err)
	})

	t.Run("Valid HTML with self-closing tag", func(t *testing.T) {
		err := HTML("<img src='image.jpg' />")
		assert.NoError(t, err)
	})

	t.Run("Valid HTML with entities", func(t *testing.T) {
		err := HTML("5 &lt; 10 &amp; 20")
		assert.NoError(t, err)
	})

	t.Run("Invalid HTML - random string", func(t *testing.T) {
		err := HTML("Just some random text")
		assert.Error(t, err)
		assert.Equal(t, "invalid html: Just some random text", err.Error())
	})

	t.Run("Invalid HTML - empty string", func(t *testing.T) {
		err := HTML("")
		assert.Error(t, err)
		assert.Equal(t, "invalid html: ", err.Error())
	})

	t.Run("Invalid input type - integer", func(t *testing.T) {
		err := HTML(12345)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid html: ")
	})

	t.Run("Invalid input type - array", func(t *testing.T) {
		err := HTML([]string{"<p>", "text", "</p>"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported type for input field")
	})
}
