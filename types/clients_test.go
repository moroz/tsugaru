package types_test

import (
	"net/url"
	"oauth-provider/types"
	"testing"

	"github.com/gorilla/schema"
	"github.com/stretchr/testify/assert"
)

func TestCreateClientParams(t *testing.T) {
	decoder := schema.NewDecoder()

	t.Run("decode from urlencoded", func(t *testing.T) {
		values := url.Values{
			"name":         {"Test app"},
			"redirectUrls": {"https://www.example.com/oauth/callback"},
		}

		var params types.CreateClientParams
		err := decoder.Decode(&params, values)
		assert.NoError(t, err)
	})

	t.Run("is valid with valid params", func(t *testing.T) {
		valid := types.CreateClientParams{
			Name: "Test app",
			RedirectURLs: []string{
				"https://www.example.com/oauth/callback",
			},
		}

		ok, errors := valid.Validate()
		assert.True(t, ok)
		assert.Nil(t, errors)
	})

	t.Run("is invalid with blank name", func(t *testing.T) {
		valid := types.CreateClientParams{
			Name: "   \t",
			RedirectURLs: []string{
				"https://www.example.com/oauth/callback",
			},
		}

		ok, errors := valid.Validate()
		assert.False(t, ok)
		assert.Nil(t, errors)
	})
}
