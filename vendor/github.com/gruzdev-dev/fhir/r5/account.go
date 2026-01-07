package models

import (
	"encoding/json"
	"fmt"
)

// A financial tool for tracking value accrued for a particular purpose.  In the healthcare field, used to track charges for a patient, cost centers, etc.
type Account struct {
	Id            *string            `json:"id,omitempty" bson:"id,omitempty"`                        // Logical id of this artifact
	Meta          *Meta              `json:"meta,omitempty" bson:"meta,omitempty"`                    // Metadata about the resource
	ImplicitRules *string            `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"` // A set of rules under which this content was created
	Language      *string            `json:"language,omitempty" bson:"language,omitempty"`            // Language of the resource content
	Text          *Narrative         `json:"text,omitempty" bson:"text,omitempty"`                    // Text summary of the resource, for human interpretation
	Contained     []json.RawMessage  `json:"contained,omitempty" bson:"contained,omitempty"`          // Contained, inline Resources
	Identifier    []Identifier       `json:"identifier,omitempty" bson:"identifier,omitempty"`        // Account number
	Status        string             `json:"status" bson:"status"`                                    // active | inactive | entered-in-error | on-hold | unknown
	BillingStatus *CodeableConcept   `json:"billingStatus,omitempty" bson:"billing_status,omitempty"` // Tracks the lifecycle of the account through the billing process
	Type          *CodeableConcept   `json:"type,omitempty" bson:"type,omitempty"`                    // E.g. patient, expense, depreciation
	Name          *string            `json:"name,omitempty" bson:"name,omitempty"`                    // Human-readable label
	Subject       []Reference        `json:"subject,omitempty" bson:"subject,omitempty"`              // The entity that caused the expenses
	ServicePeriod *Period            `json:"servicePeriod,omitempty" bson:"service_period,omitempty"` // Transaction window
	Covers        []Reference        `json:"covers,omitempty" bson:"covers,omitempty"`                // Episodic account covering these encounters/episodes of care
	Coverage      []AccountCoverage  `json:"coverage,omitempty" bson:"coverage,omitempty"`            // The party(s) that are responsible for covering the payment of this account, and what order should they be applied to the account
	Owner         *Reference         `json:"owner,omitempty" bson:"owner,omitempty"`                  // Entity managing the Account
	Description   *string            `json:"description,omitempty" bson:"description,omitempty"`      // Explanation of purpose/use
	Guarantor     []AccountGuarantor `json:"guarantor,omitempty" bson:"guarantor,omitempty"`          // The parties ultimately responsible for balancing the Account
	Diagnosis     []AccountDiagnosis `json:"diagnosis,omitempty" bson:"diagnosis,omitempty"`          // The list of diagnoses relevant to this account
	Procedure     []AccountProcedure `json:"procedure,omitempty" bson:"procedure,omitempty"`          // The list of procedures relevant to this account
	Parent        *Reference         `json:"parent,omitempty" bson:"parent,omitempty"`                // Reference to an associated parent Account
	Currency      *CodeableConcept   `json:"currency,omitempty" bson:"currency,omitempty"`            // The base or default currency
	Balance       []AccountBalance   `json:"balance,omitempty" bson:"balance,omitempty"`              // Calculated account balance(s)
	CalculatedAt  *string            `json:"calculatedAt,omitempty" bson:"calculated_at,omitempty"`   // Time the balance amount was calculated
}

func (r *Account) Validate() error {
	if r.Meta != nil {
		if err := r.Meta.Validate(); err != nil {
			return fmt.Errorf("Meta: %w", err)
		}
	}
	if r.Text != nil {
		if err := r.Text.Validate(); err != nil {
			return fmt.Errorf("Text: %w", err)
		}
	}
	for i, item := range r.Identifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Identifier[%d]: %w", i, err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.BillingStatus != nil {
		if err := r.BillingStatus.Validate(); err != nil {
			return fmt.Errorf("BillingStatus: %w", err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.Subject {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Subject[%d]: %w", i, err)
		}
	}
	if r.ServicePeriod != nil {
		if err := r.ServicePeriod.Validate(); err != nil {
			return fmt.Errorf("ServicePeriod: %w", err)
		}
	}
	for i, item := range r.Covers {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Covers[%d]: %w", i, err)
		}
	}
	for i, item := range r.Coverage {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Coverage[%d]: %w", i, err)
		}
	}
	if r.Owner != nil {
		if err := r.Owner.Validate(); err != nil {
			return fmt.Errorf("Owner: %w", err)
		}
	}
	for i, item := range r.Guarantor {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Guarantor[%d]: %w", i, err)
		}
	}
	for i, item := range r.Diagnosis {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Diagnosis[%d]: %w", i, err)
		}
	}
	for i, item := range r.Procedure {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Procedure[%d]: %w", i, err)
		}
	}
	if r.Parent != nil {
		if err := r.Parent.Validate(); err != nil {
			return fmt.Errorf("Parent: %w", err)
		}
	}
	if r.Currency != nil {
		if err := r.Currency.Validate(); err != nil {
			return fmt.Errorf("Currency: %w", err)
		}
	}
	for i, item := range r.Balance {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Balance[%d]: %w", i, err)
		}
	}
	return nil
}

type AccountGuarantor struct {
	Id             *string    `json:"id,omitempty" bson:"id,omitempty"`                         // Unique id for inter-element referencing
	Party          *Reference `json:"party,omitempty" bson:"party,omitempty"`                   // Responsible entity
	OnHold         bool       `json:"onHold,omitempty" bson:"on_hold,omitempty"`                // Credit or other hold applied
	Period         *Period    `json:"period,omitempty" bson:"period,omitempty"`                 // Guarantee account during
	Account        *Reference `json:"account,omitempty" bson:"account,omitempty"`               // A specific Account for the guarantor
	Responsibility *Quantity  `json:"responsibility,omitempty" bson:"responsibility,omitempty"` // Responsible %'age of charges
	Limit          *Money     `json:"limit,omitempty" bson:"limit,omitempty"`                   // Responsible financial limit
	Rank           *int       `json:"rank,omitempty" bson:"rank,omitempty"`                     // Rank order of guarator
}

func (r *AccountGuarantor) Validate() error {
	if r.Party != nil {
		if err := r.Party.Validate(); err != nil {
			return fmt.Errorf("Party: %w", err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	if r.Account != nil {
		if err := r.Account.Validate(); err != nil {
			return fmt.Errorf("Account: %w", err)
		}
	}
	if r.Responsibility != nil {
		if err := r.Responsibility.Validate(); err != nil {
			return fmt.Errorf("Responsibility: %w", err)
		}
	}
	if r.Limit != nil {
		if err := r.Limit.Validate(); err != nil {
			return fmt.Errorf("Limit: %w", err)
		}
	}
	return nil
}

type AccountDiagnosis struct {
	Id              *string            `json:"id,omitempty" bson:"id,omitempty"`                             // Unique id for inter-element referencing
	Sequence        *int               `json:"sequence,omitempty" bson:"sequence,omitempty"`                 // Ranking of the diagnosis (for each type)
	Condition       *CodeableReference `json:"condition" bson:"condition"`                                   // The diagnosis relevant to the account
	DateOfDiagnosis *string            `json:"dateOfDiagnosis,omitempty" bson:"date_of_diagnosis,omitempty"` // Date of the diagnosis (when coded diagnosis)
	Type            []CodeableConcept  `json:"type,omitempty" bson:"type,omitempty"`                         // Type that this diagnosis has relevant to the account (e.g. admission, billing, discharge …)
	OnAdmission     bool               `json:"onAdmission,omitempty" bson:"on_admission,omitempty"`          // Diagnosis present on Admission
	PackageCode     []CodeableConcept  `json:"packageCode,omitempty" bson:"package_code,omitempty"`          // Package Code specific for billing
}

func (r *AccountDiagnosis) Validate() error {
	if r.Condition == nil {
		return fmt.Errorf("field 'Condition' is required")
	}
	if r.Condition != nil {
		if err := r.Condition.Validate(); err != nil {
			return fmt.Errorf("Condition: %w", err)
		}
	}
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	for i, item := range r.PackageCode {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PackageCode[%d]: %w", i, err)
		}
	}
	return nil
}

type AccountProcedure struct {
	Id            *string            `json:"id,omitempty" bson:"id,omitempty"`                         // Unique id for inter-element referencing
	Sequence      *int               `json:"sequence,omitempty" bson:"sequence,omitempty"`             // Ranking of the procedure (for each type)
	Code          *CodeableReference `json:"code" bson:"code"`                                         // The procedure relevant to the account
	DateOfService *string            `json:"dateOfService,omitempty" bson:"date_of_service,omitempty"` // Date of the procedure (when coded procedure)
	Type          []CodeableConcept  `json:"type,omitempty" bson:"type,omitempty"`                     // How this procedure value should be used in charging the account
	PackageCode   []CodeableConcept  `json:"packageCode,omitempty" bson:"package_code,omitempty"`      // Package Code specific for billing
	Device        []Reference        `json:"device,omitempty" bson:"device,omitempty"`                 // Any devices that were associated with the procedure
}

func (r *AccountProcedure) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	for i, item := range r.PackageCode {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PackageCode[%d]: %w", i, err)
		}
	}
	for i, item := range r.Device {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Device[%d]: %w", i, err)
		}
	}
	return nil
}

type AccountBalance struct {
	Id        *string          `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Aggregate *CodeableConcept `json:"aggregate,omitempty" bson:"aggregate,omitempty"` // Who is expected to pay this part of the balance
	Term      *CodeableConcept `json:"term,omitempty" bson:"term,omitempty"`           // current | 30 | 60 | 90 | 120
	Estimate  bool             `json:"estimate,omitempty" bson:"estimate,omitempty"`   // Estimated balance
	Amount    *Money           `json:"amount" bson:"amount"`                           // Calculated amount
}

func (r *AccountBalance) Validate() error {
	if r.Aggregate != nil {
		if err := r.Aggregate.Validate(); err != nil {
			return fmt.Errorf("Aggregate: %w", err)
		}
	}
	if r.Term != nil {
		if err := r.Term.Validate(); err != nil {
			return fmt.Errorf("Term: %w", err)
		}
	}
	if r.Amount == nil {
		return fmt.Errorf("field 'Amount' is required")
	}
	if r.Amount != nil {
		if err := r.Amount.Validate(); err != nil {
			return fmt.Errorf("Amount: %w", err)
		}
	}
	return nil
}

type AccountCoverage struct {
	Id       *string    `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Coverage *Reference `json:"coverage" bson:"coverage"`                     // The party(s), such as insurances, that may contribute to the payment of this account
	Priority *int       `json:"priority,omitempty" bson:"priority,omitempty"` // The priority of the coverage in the context of this account
}

func (r *AccountCoverage) Validate() error {
	if r.Coverage == nil {
		return fmt.Errorf("field 'Coverage' is required")
	}
	if r.Coverage != nil {
		if err := r.Coverage.Validate(); err != nil {
			return fmt.Errorf("Coverage: %w", err)
		}
	}
	return nil
}
