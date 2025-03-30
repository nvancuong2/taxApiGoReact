package models

// TaxReturn represents the structure of a T2 tax return declaration
type TaxReturn struct {
	ID             string  `json:"id" bson:"_id"`                          // Unique identifier
	BusinessNumber string  `json:"business_number" bson:"business_number"` // Business Number (BN)
	TaxYear        int     `json:"tax_year" bson:"tax_year"`               // Tax year
	Revenue        float64 `json:"revenue" bson:"revenue"`                 // Total revenue
	Expenses       float64 `json:"expenses" bson:"expenses"`               // Total expenses
	NetIncome      float64 `json:"net_income" bson:"net_income"`           // Net income
	TaxableIncome  float64 `json:"taxable_income" bson:"taxable_income"`   // Taxable income
	TaxPayable     float64 `json:"tax_payable" bson:"tax_payable"`         // Tax payable
	Refund         float64 `json:"refund" bson:"refund"`                   // Refund amount (if applicable)
	CreatedAt      string  `json:"created_at" bson:"created_at"`           // Date of creation
	UpdatedAt      string  `json:"updated_at" bson:"updated_at"`           // Date of last update
}
