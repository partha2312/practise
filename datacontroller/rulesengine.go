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

type filteredProducts struct {
	products      []*models.Product
	productByType map[string]*priorityqueue.Queue
	remaining     *priorityqueue.Queue
}

func (p productsWithRating) CompareTo(q priorityqueue.PriorityQueue) bool {
	return p.rating > q.(productsWithRating).rating
}

func (r *rulesEngine) Filter(products []*models.Product, recos map[string]float64) ([]*models.Product, error) {
	filteredProducts := r.preprocess(products, recos)
	if r.rulesEngineEnabled[0] {
		r.productsWithHighestRatingByType(filteredProducts, r.numberOfProducts-len(filteredProducts.products))
	}
	if r.rulesEngineEnabled[1] {
		r.productsWithHighestRating(filteredProducts, r.numberOfProducts-len(filteredProducts.products))
	}
	return filteredProducts.products, nil
}

func (r *rulesEngine) preprocess(products []*models.Product, recos map[string]float64) *filteredProducts {
	filteredProducts := &filteredProducts{}
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
	filteredProducts.productByType = productByType
	filteredProducts.remaining = remaining
	return filteredProducts
}

func (r *rulesEngine) productsWithHighestRatingByType(filteredProducts *filteredProducts, toChoose int) {
	remaining := priorityqueue.NewPriorityQueue()
	for _, product := range filteredProducts.productByType {
		if topProduct, err := product.Pop(); err == nil {
			prod := topProduct.(productsWithRating).product
			filteredProducts.products = append(filteredProducts.products, prod)
			for product.Length() > 0 {
				if topProduct, err := product.Pop(); err == nil {
					remaining.Push(topProduct)
				}
			}
		}
		if toChoose--; toChoose <= 0 {
			break
		}
	}
	filteredProducts.remaining = remaining
}

func (r *rulesEngine) productsWithHighestRating(filteredProducts *filteredProducts, toChoose int) {
	for filteredProducts.remaining.Length() > 0 && toChoose > 0 {
		if topProduct, err := filteredProducts.remaining.Pop(); err == nil {
			prod := topProduct.(productsWithRating).product
			filteredProducts.products = append(filteredProducts.products, prod)
			toChoose--
		}
	}
}
