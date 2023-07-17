// Code generated by ent, DO NOT EDIT.

// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"entdemo/ent/product"
	"entdemo/ent/schema"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	productMixin := schema.Product{}.Mixin()
	productMixinFields0 := productMixin[0].Fields()
	_ = productMixinFields0
	productMixinFields1 := productMixin[1].Fields()
	_ = productMixinFields1
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescDateCreated is the schema descriptor for dateCreated field.
	productDescDateCreated := productMixinFields0[0].Descriptor()
	// product.DefaultDateCreated holds the default value on creation for the dateCreated field.
	product.DefaultDateCreated = productDescDateCreated.Default.(func() time.Time)
	// productDescDateUpdated is the schema descriptor for dateUpdated field.
	productDescDateUpdated := productMixinFields0[1].Descriptor()
	// product.DefaultDateUpdated holds the default value on creation for the dateUpdated field.
	product.DefaultDateUpdated = productDescDateUpdated.Default.(func() time.Time)
	// product.UpdateDefaultDateUpdated holds the default value on update for the dateUpdated field.
	product.UpdateDefaultDateUpdated = productDescDateUpdated.UpdateDefault.(func() time.Time)
	// productDescName is the schema descriptor for name field.
	productDescName := productFields[0].Descriptor()
	// product.NameValidator is a validator for the "name" field. It is called by the builders before save.
	product.NameValidator = productDescName.Validators[0].(func(string) error)
	// productDescPrice is the schema descriptor for price field.
	productDescPrice := productFields[2].Descriptor()
	// product.PriceValidator is a validator for the "price" field. It is called by the builders before save.
	product.PriceValidator = productDescPrice.Validators[0].(func(float64) error)
	// productDescID is the schema descriptor for id field.
	productDescID := productMixinFields1[0].Descriptor()
	// product.DefaultID holds the default value on creation for the id field.
	product.DefaultID = productDescID.Default.(func() uuid.UUID)
}


const (
	Version = "v0.12.2"                                         // Version of ent codegen.
	Sum     = "h1:Ndl/JvCX76xCtUDlrUfMnOKBRodAtxE5yfGYxjbOxmM=" // Sum of ent codegen.
)
