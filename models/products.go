package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Product struct {
	ProdID   string
	ProdName string
	ProdType string
}

func ReturnAllProducts() ([]*Product, map[string]float64) {
	products := []*Product{}
	recommendations := make(map[string]float64)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	i := 0
	for i < 10 {
		id := fmt.Sprintf("g_%v", i)
		name := fmt.Sprintf("gold_%v", i)
		products = append(products, &Product{id, name, "gold"})
		recommendations[id] = r1.Float64()
		i++
	}
	i = 0
	for i < 10 {
		id := fmt.Sprintf("s_%v", i)
		name := fmt.Sprintf("silver_%v", i)
		products = append(products, &Product{id, name, "silver"})
		recommendations[id] = r1.Float64()
		i++
	}
	i = 0
	for i < 10 {
		id := fmt.Sprintf("b_%v", i)
		name := fmt.Sprintf("bronze_%v", i)
		products = append(products, &Product{id, name, "bronze"})
		recommendations[id] = r1.Float64()
		i++
	}
	i = 0
	for i < 10 {
		id := fmt.Sprintf("d_%v", i)
		name := fmt.Sprintf("diamond_%v", i)
		products = append(products, &Product{id, name, "diamond"})
		recommendations[id] = r1.Float64()
		i++
	}
	return products, recommendations
}
