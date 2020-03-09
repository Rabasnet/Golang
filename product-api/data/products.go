package data

import (
	"encoding/json"
	"io"
)

//product defines the structure of the api product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json: "description"`
	Price       float64 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json: "-"`
	DeletedOn   string  `json: "-"`
}

//Products is an array collection of Product
type Products []*Product

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

//work on progress
