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

func TestDB_Products_ListWithActiveDiscounts(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	DbMgr.CreateProduct(&Product{
		ID:   "kek",
		Name: "god of war 3",
	})
	DbMgr.EnsureDiscountExists(&Discount{
		ProductID: "kek",
		IsPlus:    true,
		Value:     10,
	})
	DbMgr.EnsureDiscountExists(&Discount{
		ProductID: "kek",
		IsPlus:    false,
		Value:     50,
	})
	DbMgr.EnsureDiscountExists(&Discount{
		ProductID: "another-game",
		IsPlus:    false,
		Value:     50,
	})
	DbMgr.EnsureRateExists(&Rate{
		ProductID: "kek",
	})
	DbMgr.EnsureRateExists(&Rate{
		ProductID: "lolz",
	})
	DbMgr.EnsurePosterExists(&Poster{
		ProductID: "kek",
		URL:       "cdn/url",
	})

	// action
	products, err := DbMgr.GetAllProductsWithActiveDiscounts()

	// assert
	assert.NoError(t, err)
	assert.Equal(t, "cdn/url", products[0].Poster.URL)
	assert.Len(t, products, 1)
	assert.Len(t, products[0].Discounts, 2)
	for _, disc := range products[0].Discounts {
		assert.Equal(t, "kek", disc.ProductID)
	}
}

func TestDB_Products_Get_FilterBy_ID(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	DbMgr.CreateProduct(&Product{
		ID:   "kek",
		Name: "god of war 3",
	})
	DbMgr.EnsureDiscountExists(&Discount{
		ProductID: "kek",
		IsPlus:    true,
		Value:     10,
	})
	DbMgr.EnsureDiscountExists(&Discount{
		ProductID: "kek",
		IsPlus:    false,
		Value:     50,
	})
	DbMgr.EnsureDiscountExists(&Discount{
		ProductID: "another-game",
		IsPlus:    false,
		Value:     50,
	})
	DbMgr.EnsureRateExists(&Rate{
		ProductID: "kek",
	})
	DbMgr.EnsureRateExists(&Rate{
		ProductID: "lolz",
	})
	DbMgr.EnsurePosterExists(&Poster{
		ProductID: "kek",
		URL:       "cdn/url",
	})

	// action
	product, err := DbMgr.GetProductByID("kek")

	// assert
	assert.NoError(t, err)
	assert.Equal(t, "cdn/url", product.Poster.URL)
	assert.Len(t, product.Discounts, 2)
	for _, disc := range product.Discounts {
		assert.Equal(t, "kek", disc.ProductID)
	}
}

func TestDB_Products_SearchProductsByName(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	DbMgr.CreateProduct(&Product{
		ID:   "EP4064-CUSA01857_00-THEESCAPISTS0000",
		Name: "the Escapists",
	})
	DbMgr.CreateProduct(&Product{
		ID:   "EP9000-CUSA00176_00-DEADNATIONGAME01",
		Name: "Dead Nation",
	})
	DbMgr.CreateProduct(&Product{
		ID:   "EP2034-CUSA03952_00-NOMANSSKYHG00001",
		Name: "No Man's Sky",
	})

	// action
	products, err := DbMgr.SearchProductsByName("dead nation")

	// assert
	assert.NoError(t, err)
	assert.Len(t, products, 1)
	assert.Equal(t, "EP9000-CUSA00176_00-DEADNATIONGAME01", products[0].ID)
	assert.Equal(t, "Dead Nation", products[0].Name)
}

func TestDB_Products_SearchProductsByName_Nothing(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	DbMgr.CreateProduct(&Product{
		ID:   "EP4064-CUSA01857_00-THEESCAPISTS0000",
		Name: "the Escapists",
	})
	DbMgr.CreateProduct(&Product{
		ID:   "EP9000-CUSA00176_00-DEADNATIONGAME01",
		Name: "Dead Nation",
	})
	DbMgr.CreateProduct(&Product{
		ID:   "EP2034-CUSA03952_00-NOMANSSKYHG00001",
		Name: "No Man's Sky",
	})

	// action
	products, err := DbMgr.SearchProductsByName("god of war') --")

	// assert
	assert.NoError(t, err)
	assert.Len(t, products, 0)
}

func TestDB_Products_SearchProductsByName_SQL_Injection(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	DbMgr.CreateProduct(&Product{
		ID:   "EP4064-CUSA01857_00-THEESCAPISTS0000",
		Name: "the Escapists",
	})
	DbMgr.CreateProduct(&Product{
		ID:   "EP9000-CUSA00176_00-DEADNATIONGAME01",
		Name: "Dead Nation",
	})
	DbMgr.CreateProduct(&Product{
		ID:   "EP2034-CUSA03952_00-NOMANSSKYHG00001",
		Name: "No Man's Sky",
	})

	// action
	products, err := DbMgr.SearchProductsByName(`god of war'" -- broken sql`)

	// assert
	assert.NoError(t, err)
	assert.Len(t, products, 0)
}
