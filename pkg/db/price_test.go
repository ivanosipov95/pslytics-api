package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB_Price_Get(t *testing.T) {
	setup()
	defer teardown()

	// action
	price, err := DbMgr.GetPriceForProduct("EP4139-CUSA01400_00-MAMA02GP40000002")

	// assert
	assert.Error(t, err)
	assert.Nil(t, price)
}

func TestDB_Price_EnsureExists(t *testing.T) {
	setup()
	defer teardown()

	// action
	err := DbMgr.EnsurePriceExists(&Price{
		ProductID: "EP4139-CUSA01400_00-MAMA02GP40000002",
		Value:     3999,
	})

	// assert
	assert.NoError(t, err)
	price, err := DbMgr.GetPriceForProduct("EP4139-CUSA01400_00-MAMA02GP40000002")
	assert.NoError(t, err)
	assert.Equal(t, int64(3999), price.Value)
	assert.Equal(t, "EP4139-CUSA01400_00-MAMA02GP40000002", price.ProductID)
}
