package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// TaxReturn represents the structure of a T2 tax return declaration
type TaxReturn struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"` // MongoDB ObjectID
	BusinessNumber string             `json:"business_number" bson:"business_number"`
	TaxYear        int                `json:"tax_year" bson:"tax_year"`
	Revenue        float64            `json:"revenue" bson:"revenue"`
	Expenses       float64            `json:"expenses" bson:"expenses"`
	NetIncome      float64            `json:"net_income" bson:"net_income"`
	TaxableIncome  float64            `json:"taxable_income" bson:"taxable_income"`
	TaxPayable     float64            `json:"tax_payable" bson:"tax_payable"`
	Refund         float64            `json:"refund" bson:"refund"`
	CreatedAt      primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt      primitive.DateTime `json:"updated_at" bson:"updated_at"`
}
