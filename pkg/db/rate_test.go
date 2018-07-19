package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB_Rate_EnsureExists_NewRate(t *testing.T) {
	setup()
	defer teardown()

	// action
	err := DbMgr.EnsureRateExists(&Rate{
		ProductID: "EP4139-CUSA01400_00-MAMA02GP40000002",
		Total:     100,
		Value:     4.9,
	})

	// assert
	assert.NoError(t, err)
	rate, err := DbMgr.GetRateForProduct("EP4139-CUSA01400_00-MAMA02GP40000002")
	assert.NoError(t, err)
	assert.Equal(t, "EP4139-CUSA01400_00-MAMA02GP40000002", rate.ProductID)
	assert.Equal(t, int64(100), rate.Total)
	assert.Equal(t, 4.9, rate.Value)
}


func TestDB_Rate_EnsureExists_UpdateRate(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	DbMgr.EnsureRateExists(&Rate{
		ProductID: "EP4139-CUSA01400_00-MAMA02GP40000002",
		Total:     50,
		Value:     3.5,
	})

	// action
	err := DbMgr.EnsureRateExists(&Rate{
		ProductID: "EP4139-CUSA01400_00-MAMA02GP40000002",
		Total:     100,
		Value:     4.9,
	})

	// assert
	assert.NoError(t, err)
	rate, err := DbMgr.GetRateForProduct("EP4139-CUSA01400_00-MAMA02GP40000002")
	assert.NoError(t, err)
	assert.Equal(t, "EP4139-CUSA01400_00-MAMA02GP40000002", rate.ProductID)
	assert.Equal(t, int64(100), rate.Total)
	assert.Equal(t, 4.9, rate.Value)
}
