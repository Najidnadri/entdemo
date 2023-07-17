// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entdemo/ent/predicate"
	"entdemo/ent/product"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// ProductWhereInput represents a where input for filtering Product queries.
type ProductWhereInput struct {
	Predicates []predicate.Product  `json:"-"`
	Not        *ProductWhereInput   `json:"not,omitempty"`
	Or         []*ProductWhereInput `json:"or,omitempty"`
	And        []*ProductWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *uuid.UUID  `json:"id,omitempty"`
	IDNEQ   *uuid.UUID  `json:"idNEQ,omitempty"`
	IDIn    []uuid.UUID `json:"idIn,omitempty"`
	IDNotIn []uuid.UUID `json:"idNotIn,omitempty"`
	IDGT    *uuid.UUID  `json:"idGT,omitempty"`
	IDGTE   *uuid.UUID  `json:"idGTE,omitempty"`
	IDLT    *uuid.UUID  `json:"idLT,omitempty"`
	IDLTE   *uuid.UUID  `json:"idLTE,omitempty"`

	// "dateCreated" field predicates.
	DateCreated      *time.Time  `json:"datecreated,omitempty"`
	DateCreatedNEQ   *time.Time  `json:"datecreatedNEQ,omitempty"`
	DateCreatedIn    []time.Time `json:"datecreatedIn,omitempty"`
	DateCreatedNotIn []time.Time `json:"datecreatedNotIn,omitempty"`
	DateCreatedGT    *time.Time  `json:"datecreatedGT,omitempty"`
	DateCreatedGTE   *time.Time  `json:"datecreatedGTE,omitempty"`
	DateCreatedLT    *time.Time  `json:"datecreatedLT,omitempty"`
	DateCreatedLTE   *time.Time  `json:"datecreatedLTE,omitempty"`

	// "dateUpdated" field predicates.
	DateUpdated      *time.Time  `json:"dateupdated,omitempty"`
	DateUpdatedNEQ   *time.Time  `json:"dateupdatedNEQ,omitempty"`
	DateUpdatedIn    []time.Time `json:"dateupdatedIn,omitempty"`
	DateUpdatedNotIn []time.Time `json:"dateupdatedNotIn,omitempty"`
	DateUpdatedGT    *time.Time  `json:"dateupdatedGT,omitempty"`
	DateUpdatedGTE   *time.Time  `json:"dateupdatedGTE,omitempty"`
	DateUpdatedLT    *time.Time  `json:"dateupdatedLT,omitempty"`
	DateUpdatedLTE   *time.Time  `json:"dateupdatedLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "description" field predicates.
	Description             *string  `json:"description,omitempty"`
	DescriptionNEQ          *string  `json:"descriptionNEQ,omitempty"`
	DescriptionIn           []string `json:"descriptionIn,omitempty"`
	DescriptionNotIn        []string `json:"descriptionNotIn,omitempty"`
	DescriptionGT           *string  `json:"descriptionGT,omitempty"`
	DescriptionGTE          *string  `json:"descriptionGTE,omitempty"`
	DescriptionLT           *string  `json:"descriptionLT,omitempty"`
	DescriptionLTE          *string  `json:"descriptionLTE,omitempty"`
	DescriptionContains     *string  `json:"descriptionContains,omitempty"`
	DescriptionHasPrefix    *string  `json:"descriptionHasPrefix,omitempty"`
	DescriptionHasSuffix    *string  `json:"descriptionHasSuffix,omitempty"`
	DescriptionEqualFold    *string  `json:"descriptionEqualFold,omitempty"`
	DescriptionContainsFold *string  `json:"descriptionContainsFold,omitempty"`

	// "price" field predicates.
	Price      *float64  `json:"price,omitempty"`
	PriceNEQ   *float64  `json:"priceNEQ,omitempty"`
	PriceIn    []float64 `json:"priceIn,omitempty"`
	PriceNotIn []float64 `json:"priceNotIn,omitempty"`
	PriceGT    *float64  `json:"priceGT,omitempty"`
	PriceGTE   *float64  `json:"priceGTE,omitempty"`
	PriceLT    *float64  `json:"priceLT,omitempty"`
	PriceLTE   *float64  `json:"priceLTE,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *ProductWhereInput) AddPredicates(predicates ...predicate.Product) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the ProductWhereInput filter on the ProductQuery builder.
func (i *ProductWhereInput) Filter(q *ProductQuery) (*ProductQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyProductWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyProductWhereInput is returned in case the ProductWhereInput is empty.
var ErrEmptyProductWhereInput = errors.New("ent: empty predicate ProductWhereInput")

// P returns a predicate for filtering products.
// An error is returned if the input is empty or invalid.
func (i *ProductWhereInput) P() (predicate.Product, error) {
	var predicates []predicate.Product
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, product.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Product, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, product.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Product, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, product.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, product.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, product.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, product.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, product.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, product.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, product.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, product.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, product.IDLTE(*i.IDLTE))
	}
	if i.DateCreated != nil {
		predicates = append(predicates, product.DateCreatedEQ(*i.DateCreated))
	}
	if i.DateCreatedNEQ != nil {
		predicates = append(predicates, product.DateCreatedNEQ(*i.DateCreatedNEQ))
	}
	if len(i.DateCreatedIn) > 0 {
		predicates = append(predicates, product.DateCreatedIn(i.DateCreatedIn...))
	}
	if len(i.DateCreatedNotIn) > 0 {
		predicates = append(predicates, product.DateCreatedNotIn(i.DateCreatedNotIn...))
	}
	if i.DateCreatedGT != nil {
		predicates = append(predicates, product.DateCreatedGT(*i.DateCreatedGT))
	}
	if i.DateCreatedGTE != nil {
		predicates = append(predicates, product.DateCreatedGTE(*i.DateCreatedGTE))
	}
	if i.DateCreatedLT != nil {
		predicates = append(predicates, product.DateCreatedLT(*i.DateCreatedLT))
	}
	if i.DateCreatedLTE != nil {
		predicates = append(predicates, product.DateCreatedLTE(*i.DateCreatedLTE))
	}
	if i.DateUpdated != nil {
		predicates = append(predicates, product.DateUpdatedEQ(*i.DateUpdated))
	}
	if i.DateUpdatedNEQ != nil {
		predicates = append(predicates, product.DateUpdatedNEQ(*i.DateUpdatedNEQ))
	}
	if len(i.DateUpdatedIn) > 0 {
		predicates = append(predicates, product.DateUpdatedIn(i.DateUpdatedIn...))
	}
	if len(i.DateUpdatedNotIn) > 0 {
		predicates = append(predicates, product.DateUpdatedNotIn(i.DateUpdatedNotIn...))
	}
	if i.DateUpdatedGT != nil {
		predicates = append(predicates, product.DateUpdatedGT(*i.DateUpdatedGT))
	}
	if i.DateUpdatedGTE != nil {
		predicates = append(predicates, product.DateUpdatedGTE(*i.DateUpdatedGTE))
	}
	if i.DateUpdatedLT != nil {
		predicates = append(predicates, product.DateUpdatedLT(*i.DateUpdatedLT))
	}
	if i.DateUpdatedLTE != nil {
		predicates = append(predicates, product.DateUpdatedLTE(*i.DateUpdatedLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, product.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, product.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, product.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, product.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, product.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, product.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, product.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, product.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, product.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, product.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, product.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, product.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, product.NameContainsFold(*i.NameContainsFold))
	}
	if i.Description != nil {
		predicates = append(predicates, product.DescriptionEQ(*i.Description))
	}
	if i.DescriptionNEQ != nil {
		predicates = append(predicates, product.DescriptionNEQ(*i.DescriptionNEQ))
	}
	if len(i.DescriptionIn) > 0 {
		predicates = append(predicates, product.DescriptionIn(i.DescriptionIn...))
	}
	if len(i.DescriptionNotIn) > 0 {
		predicates = append(predicates, product.DescriptionNotIn(i.DescriptionNotIn...))
	}
	if i.DescriptionGT != nil {
		predicates = append(predicates, product.DescriptionGT(*i.DescriptionGT))
	}
	if i.DescriptionGTE != nil {
		predicates = append(predicates, product.DescriptionGTE(*i.DescriptionGTE))
	}
	if i.DescriptionLT != nil {
		predicates = append(predicates, product.DescriptionLT(*i.DescriptionLT))
	}
	if i.DescriptionLTE != nil {
		predicates = append(predicates, product.DescriptionLTE(*i.DescriptionLTE))
	}
	if i.DescriptionContains != nil {
		predicates = append(predicates, product.DescriptionContains(*i.DescriptionContains))
	}
	if i.DescriptionHasPrefix != nil {
		predicates = append(predicates, product.DescriptionHasPrefix(*i.DescriptionHasPrefix))
	}
	if i.DescriptionHasSuffix != nil {
		predicates = append(predicates, product.DescriptionHasSuffix(*i.DescriptionHasSuffix))
	}
	if i.DescriptionEqualFold != nil {
		predicates = append(predicates, product.DescriptionEqualFold(*i.DescriptionEqualFold))
	}
	if i.DescriptionContainsFold != nil {
		predicates = append(predicates, product.DescriptionContainsFold(*i.DescriptionContainsFold))
	}
	if i.Price != nil {
		predicates = append(predicates, product.PriceEQ(*i.Price))
	}
	if i.PriceNEQ != nil {
		predicates = append(predicates, product.PriceNEQ(*i.PriceNEQ))
	}
	if len(i.PriceIn) > 0 {
		predicates = append(predicates, product.PriceIn(i.PriceIn...))
	}
	if len(i.PriceNotIn) > 0 {
		predicates = append(predicates, product.PriceNotIn(i.PriceNotIn...))
	}
	if i.PriceGT != nil {
		predicates = append(predicates, product.PriceGT(*i.PriceGT))
	}
	if i.PriceGTE != nil {
		predicates = append(predicates, product.PriceGTE(*i.PriceGTE))
	}
	if i.PriceLT != nil {
		predicates = append(predicates, product.PriceLT(*i.PriceLT))
	}
	if i.PriceLTE != nil {
		predicates = append(predicates, product.PriceLTE(*i.PriceLTE))
	}

	switch len(predicates) {
	case 0:
		return nil, ErrEmptyProductWhereInput
	case 1:
		return predicates[0], nil
	default:
		return product.And(predicates...), nil
	}
}