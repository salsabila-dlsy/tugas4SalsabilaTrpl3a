package services

import (
	"context"
	"database/sql"
	"fmt"

	cm "tugas4SalsabilaTrpl3a/Framework/git/order/common"

	_ "github.com/go-sql-driver/mysql"
)

func (PaymentService) OrderHandler(ctx context.Context, req cm.Message) (res cm.Message) {

	var db *sql.DB
	var err error

	defer panicRecovery()

	host := cm.Config.Connection.Host
	port := cm.Config.Connection.Port
	user := cm.Config.Connection.User
	pass := cm.Config.Connection.Password
	data := cm.Config.Connection.Database

	var mySQL = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, host, port, data)

	db, err = sql.Open("mysql", mySQL)

	if err != nil {
		panic(err.Error())
	}

	res.ProductID = req.ProductID

	var product cm.Products	
		

	sql := `SELECT
				ProductID,
				IFNULL(SupplierID,''),
				IFNULL(Category,'') Category,
				IFNULL(QuantityPerUnit,'') QuantityPerUnit,
				IFNULL(UnitePrice,'') UnitePrice,
				IFNULL(UnitsInStock,'') UnitsInStock,
				IFNULL(UnitsOrder,'') UnitsOrder,
				IFNULL(ReorderLevel,'') ReorderLevel ,
				IFNULL(Discontinued,'') Discontinued
			FROM products WHERE ProductID = ?`

   
	result, err := db.Query(sql, req.ProductID)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}
	
	for result.Next() {

		err := result.Scan(&product.ProductID, &product.SupplierID, &product.Category, &product.QuantityPerUnit, &product.UnitePrice, &product.UnitsInStock, &product.UnitsOrder, &product.ReorderLevel, &product.Discontinued)


		if err != nil {
			panic(err.Error())
		}

		products = append(products,product)
			
		}
				
	}
	if &product != nil {
		res.Code = 100
		res.Remark = "Success"
	}

	res.Products = &product	
	
	return
}