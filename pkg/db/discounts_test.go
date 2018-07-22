package db

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func timePointer(time time.Time) *time.Time {
	return &time
}

func TestDB_Discounts_Get(t *testing.T) {
	setup()
	defer teardown()

	// action
	discount, err := DbMgr.GetDiscountForProduct("EP4139-CUSA01400_00-MAMA02GP40000002", false)

	// assert
	assert.Error(t, err)
	assert.Nil(t, discount)
}

func TestDB_Discounts_EnsureExists(t *testing.T) {
	setup()
	defer teardown()

	// action
	err := DbMgr.EnsureDiscountExists(&Discount{
		ProductID:  "EP4139-CUSA01400_00-MAMA02GP40000002",
		IsPlus:     false,
		Value:      1000,
		Percentage: 50,
		Since:      timePointer(time.Now().UTC()),
		Till:       timePointer(time.Now().UTC().Add(time.Hour * 24)),
	})

	// assert
	assert.NoError(t, err)
	poster, err := DbMgr.GetDiscountForProduct("EP4139-CUSA01400_00-MAMA02GP40000002", false)
	assert.NoError(t, err)
	assert.Equal(t, "EP4139-CUSA01400_00-MAMA02GP40000002", poster.ProductID)
	assert.Equal(t, int64(1000), poster.Value)
	assert.Equal(t, int64(50), poster.Percentage)
}

func TestDB_Discounts_EnsureNotExists(t *testing.T) {
	setup()
	defer teardown()

	// action
	err := DbMgr.EnsureDiscountExists(&Discount{
		ProductID:  "EP4139-CUSA01400_00-MAMA02GP40000002",
		IsPlus:     false,
		Value:      1000,
		Percentage: 50,
		Since:      timePointer(time.Now().UTC()),
		Till:       timePointer(time.Now().UTC().Add(time.Hour * 24)),
	})

	// assert
	assert.NoError(t, err)
	_, err = DbMgr.GetDiscountForProduct("EP4139-CUSA01400_00-MAMA02GP40000002", true)
	assert.Error(t, err)
}
