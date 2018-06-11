package datacontroller

import (
	"practise/models"
	"practise/priorityqueue"
)

type RulesEngine interface {
	Filter(products []*models.Product, recos map[string]float64) ([]*models.Product, error)
}

type rulesEngine struct {
	numberOfProducts   int
	rulesEngineEnabled []bool
}

func NewRulesEngine(numberOfProducts int, rulesEngineEnabled []bool) RulesEngine {
	return &rulesEngine{numberOfProducts, rulesEngineEnabled}
}

type productsWithRating struct {
	product *models.Product
	rating  float64
}

func (p productsWithRating) CompareTo(q priorityqueue.PriorityQueue) bool {
	return p.rating > q.(productsWithRating).rating
}

func (r *rulesEngine) ruleZero(products []*models.Product, recos map[string]float64) (map[string]*priorityqueue.Queue, *priorityqueue.Queue) {
	remaining := priorityqueue.NewPriorityQueue()
	productByType := make(map[string]*priorityqueue.Queue)
	for _, product := range products {
		t := product.ProdType
		if _, ok := productByType[t]; !ok {
			productByType[t] = priorityqueue.NewPriorityQueue()
		}
		prodWithRating := productsWithRating{product, recos[product.ProdID]}
		productByType[t].Push(prodWithRating)
		remaining.Push(prodWithRating)
	}
	return productByType, remaining
}

func (r *rulesEngine) ruleOne(productByType map[string]*priorityqueue.Queue) ([]*models.Product, *priorityqueue.Queue) {
	filteredProducts := []*models.Product{}
	remaining := priorityqueue.NewPriorityQueue()
	for _, product := range productByType {
		if topProduct, err := product.Pop(); err == nil {
			prod := topProduct.(productsWithRating).product
			filteredProducts = append(filteredProducts, prod)
			for product.Length() > 0 {
				if topProduct, err := product.Pop(); err == nil {
					remaining.Push(topProduct)
				}
			}
		}
	}
	return filteredProducts, remaining
}

func (r *rulesEngine) ruleTwo(products *priorityqueue.Queue, toChoose int) ([]*models.Product, *priorityqueue.Queue) {
	filteredProducts := []*models.Product{}
	for products.Length() > 0 && toChoose > 0 {
		if topProduct, err := products.Pop(); err == nil {
			prod := topProduct.(productsWithRating).product
			filteredProducts = append(filteredProducts, prod)
			toChoose--
		}
	}
	return filteredProducts, products
}

func (r *rulesEngine) Filter(products []*models.Product, recos map[string]float64) ([]*models.Product, error) {
	filteredProducts := []*models.Product{}
	ruleZeroProducts, remaining := r.ruleZero(products, recos)
	ruleOneProducts := []*models.Product{}
	ruleTwoProducts := []*models.Product{}
	if r.rulesEngineEnabled[0] {
		ruleOneProducts, remaining = r.ruleOne(ruleZeroProducts)
	}
	if r.rulesEngineEnabled[1] {
		ruleTwoProducts, remaining = r.ruleTwo(remaining, r.numberOfProducts-len(ruleOneProducts))
	}
	for _, p := range ruleOneProducts {
		filteredProducts = append(filteredProducts, p)
	}
	for _, p := range ruleTwoProducts {
		filteredProducts = append(filteredProducts, p)
	}
	return filteredProducts, nil
}
