package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `jaon:"price validate:"gt=0"`
	Sku         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	// sku is of format abc-ab-df
	re := regexp.MustCompile(`[a-z]+-[0-9]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) == 1 {
		return true
	}

	return false
}

/*
 * to/from Json serializes the contents of the collection
 * Json encoder provides better performance than json.Un/Marshal as it does not
 * Have to buffer the output into a n in memory slice of bytes
 * this reduces allocations and the overheads of the service
 */
func (p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func AddProduct(p *Product) {
	p.ID = getNextId()
	productList = append(productList, p)
}

func getNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func GetProducts() Products {
	return productList
}

// dummy data
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		Sku:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Expresso",
		Description: "dark coffee without milk and sugar",
		Price:       1.99,
		Sku:         "figi123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          3,
		Name:        "Tea",
		Description: "Java tea, empowers my android development",
		Price:       0.99,
		Sku:         "paris69",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
