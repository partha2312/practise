package main

import (
	"fmt"
	"practise/datacontroller"
	"practise/models"
	"time"
)

func main() {
	fmt.Println("hello, world")
	products, recommendations := models.ReturnAllProducts()
	rulesEngineEnable := []bool{true, true}
	rulesEngine := datacontroller.NewRulesEngine(6, rulesEngineEnable)
	start := time.Now()
	filteredProducts, err := rulesEngine.Filter(products, recommendations)
	since := time.Since(start)
	fmt.Println(since.Seconds())
	if err == nil {
		for _, product := range filteredProducts {
			fmt.Println(product, recommendations[product.ProdID])
		}
	} else {
		fmt.Println(err)
	}
}
