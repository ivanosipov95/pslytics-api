package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB_Products_Create(t *testing.T) {
	setup()
	defer teardown()

	err := DbMgr.CreateProduct(&Product{
		ID:   "god-of-war",
		Name: "god of war 3",
	})

	assert.NoError(t, err)
}

func TestDB_Products_List(t *testing.T) {
	setup()
	defer teardown()

	DbMgr.CreateProduct(&Product{
		ID:   "god-of-war",
		Name: "god of war 3",
	})

	products, err := DbMgr.GetAllProducts()

	assert.NoError(t, err)
	assert.Len(t, products, 1)
	assert.Equal(t, "god-of-war", products[0].ID)
	assert.Equal(t, "god of war 3", products[0].Name)
}
