package handlers

import (
	"log"
	"net/http"
	"product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l: l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	if h.Method == http.MethodGet {
		p.getProducts(rw, h)
		return
	} else if h.Method == http.MethodPost {
		p.postProduct(rw, h)
		return
	}

	// catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) postProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Products")

	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	data.AddProduct(prod)

}

func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()

	// d, err := json.Marshal(lp)

	err := lp.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	// rw.Write(d)

	// Using Json encoding
}
