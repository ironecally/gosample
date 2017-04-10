package product

import (
	"database/sql"
	"fmt"

	"github.com/tokopedia/gosample/db"
)

type ProductInfo struct {
	ProductID   int64         `db:"product_id"`
	ProductName string        `db:"product_name"`
	Discount    sql.NullInt64 `db:"discount"`
}

func GetProduct(productID int64) (ProductInfo, error) {
	query := `
		SELECT 
			product_id,
			product_name
		FROM
			ws_product
		WHERE
			product_id = $1`

	// var dbResult []ProductInfo
	// db.DBPools.DB1.Select(&dbResult, query)

	var dbResult ProductInfo
	err := db.DBPools.DB1.Get(&dbResult, query, productID)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("error get product data", err.Error())
		return dbResult, err
	}

	// row, _ := db.DBPools.DB1.Query(query, productID)
	// for row.Next() {
	// 	var temp1 int64
	// 	var temp2 string
	// 	row.Scan(&temp1, &temp2)
	// 	dbResult = append(dbResult, ProductInfo{ProductID: temp1, ProductName: temp2})
	// }
	// err := db.DBPools.DB1.QueryRow(query).Scan(&dbResult.ProductID, &dbResult.ProductName)
	// if err != nil {
	// 	fmt.Println("error get product data", err.Error())
	// }

	return dbResult, nil
}
