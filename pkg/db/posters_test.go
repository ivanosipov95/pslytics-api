package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB_Posters_Get(t *testing.T) {
	setup()
	defer teardown()

	// action
	poster, err := DbMgr.GetPosterForProduct("EP4139-CUSA01400_00-MAMA02GP40000002")

	// assert
	assert.Error(t, err)
	assert.Nil(t, poster)
}

func TestDB_Posters_EnsureExists(t *testing.T) {
	setup()
	defer teardown()

	// action
	err := DbMgr.EnsurePosterExists(&Poster{
		ProductID: "EP4139-CUSA01400_00-MAMA02GP40000002",
		URL:       "http://cdn/pic.jpg",
	})

	// assert
	assert.NoError(t, err)
	poster, err := DbMgr.GetPosterForProduct("EP4139-CUSA01400_00-MAMA02GP40000002")
	assert.NoError(t, err)
	assert.Equal(t, "http://cdn/pic.jpg", poster.URL)
	assert.Equal(t, "EP4139-CUSA01400_00-MAMA02GP40000002", poster.ProductID)
}
