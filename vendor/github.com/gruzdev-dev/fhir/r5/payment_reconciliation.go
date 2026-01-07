package models

import (
	"encoding/json"
	"fmt"
)

// This resource provides the details including amount of a payment and allocates the payment items being paid.
type PaymentReconciliation struct {
	ResourceType      string                             `json:"resourceType" bson:"resource_type"`                               // Type of resource
	Id                *string                            `json:"id,omitempty" bson:"id,omitempty"`                                // Logical id of this artifact
	Meta              *Meta                              `json:"meta,omitempty" bson:"meta,omitempty"`                            // Metadata about the resource
	ImplicitRules     *string                            `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`         // A set of rules under which this content was created
	Language          *string                            `json:"language,omitempty" bson:"language,omitempty"`                    // Language of the resource content
	Text              *Narrative                         `json:"text,omitempty" bson:"text,omitempty"`                            // Text summary of the resource, for human interpretation
	Contained         []json.RawMessage                  `json:"contained,omitempty" bson:"contained,omitempty"`                  // Contained, inline Resources
	Identifier        []Identifier                       `json:"identifier,omitempty" bson:"identifier,omitempty"`                // Business Identifier for a payment reconciliation
	Type              *CodeableConcept                   `json:"type" bson:"type"`                                                // Category of payment
	Status            string                             `json:"status" bson:"status"`                                            // active | cancelled | draft | entered-in-error
	StatusReason      *string                            `json:"statusReason,omitempty" bson:"status_reason,omitempty"`           // Reason for status change
	Kind              *CodeableConcept                   `json:"kind,omitempty" bson:"kind,omitempty"`                            // Workflow originating payment
	Period            *Period                            `json:"period,omitempty" bson:"period,omitempty"`                        // Period covered
	Created           string                             `json:"created" bson:"created"`                                          // Creation date
	Enterer           *Reference                         `json:"enterer,omitempty" bson:"enterer,omitempty"`                      // Who entered the payment
	IssuerType        *CodeableConcept                   `json:"issuerType,omitempty" bson:"issuer_type,omitempty"`               // Nature of the source
	PaymentIssuer     *Reference                         `json:"paymentIssuer,omitempty" bson:"payment_issuer,omitempty"`         // Party generating payment
	Request           *Reference                         `json:"request,omitempty" bson:"request,omitempty"`                      // Reference to requesting resource
	Requestor         *Reference                         `json:"requestor,omitempty" bson:"requestor,omitempty"`                  // Responsible practitioner
	Outcome           *string                            `json:"outcome,omitempty" bson:"outcome,omitempty"`                      // queued | complete | error | partial
	Disposition       *string                            `json:"disposition,omitempty" bson:"disposition,omitempty"`              // Disposition message
	Date              string                             `json:"date" bson:"date"`                                                // When payment issued
	Location          *Reference                         `json:"location,omitempty" bson:"location,omitempty"`                    // Where payment collected
	Method            *CodeableConcept                   `json:"method,omitempty" bson:"method,omitempty"`                        // Payment instrument
	CardBrand         *string                            `json:"cardBrand,omitempty" bson:"card_brand,omitempty"`                 // Type of card
	AccountNumber     *string                            `json:"accountNumber,omitempty" bson:"account_number,omitempty"`         // Digits for verification
	ExpirationDate    *string                            `json:"expirationDate,omitempty" bson:"expiration_date,omitempty"`       // Expiration year-month
	Processor         *string                            `json:"processor,omitempty" bson:"processor,omitempty"`                  // Processor name
	ReferenceNumber   *string                            `json:"referenceNumber,omitempty" bson:"reference_number,omitempty"`     // Check number or payment reference
	Authorization     *string                            `json:"authorization,omitempty" bson:"authorization,omitempty"`          // Authorization number
	TenderedAmount    *Money                             `json:"tenderedAmount,omitempty" bson:"tendered_amount,omitempty"`       // Amount offered by the issuer
	ReturnedAmount    *Money                             `json:"returnedAmount,omitempty" bson:"returned_amount,omitempty"`       // Amount returned by the receiver
	Amount            *Money                             `json:"amount,omitempty" bson:"amount,omitempty"`                        // Total amount of Payment
	PaymentIdentifier *Identifier                        `json:"paymentIdentifier,omitempty" bson:"payment_identifier,omitempty"` // Business identifier for the payment
	Allocation        []PaymentReconciliationAllocation  `json:"allocation,omitempty" bson:"allocation,omitempty"`                // Settlement particulars
	FormCode          *CodeableConcept                   `json:"formCode,omitempty" bson:"form_code,omitempty"`                   // Printed form identifier
	ProcessNote       []PaymentReconciliationProcessNote `json:"processNote,omitempty" bson:"process_note,omitempty"`             // Note concerning processing
}

func (r *PaymentReconciliation) Validate() error {
	if r.ResourceType != "PaymentReconciliation" {
		return fmt.Errorf("invalid resourceType: expected 'PaymentReconciliation', got '%s'", r.ResourceType)
	}
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
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Kind != nil {
		if err := r.Kind.Validate(); err != nil {
			return fmt.Errorf("Kind: %w", err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	if r.Created == emptyString {
		return fmt.Errorf("field 'Created' is required")
	}
	if r.Enterer != nil {
		if err := r.Enterer.Validate(); err != nil {
			return fmt.Errorf("Enterer: %w", err)
		}
	}
	if r.IssuerType != nil {
		if err := r.IssuerType.Validate(); err != nil {
			return fmt.Errorf("IssuerType: %w", err)
		}
	}
	if r.PaymentIssuer != nil {
		if err := r.PaymentIssuer.Validate(); err != nil {
			return fmt.Errorf("PaymentIssuer: %w", err)
		}
	}
	if r.Request != nil {
		if err := r.Request.Validate(); err != nil {
			return fmt.Errorf("Request: %w", err)
		}
	}
	if r.Requestor != nil {
		if err := r.Requestor.Validate(); err != nil {
			return fmt.Errorf("Requestor: %w", err)
		}
	}
	if r.Date == emptyString {
		return fmt.Errorf("field 'Date' is required")
	}
	if r.Location != nil {
		if err := r.Location.Validate(); err != nil {
			return fmt.Errorf("Location: %w", err)
		}
	}
	if r.Method != nil {
		if err := r.Method.Validate(); err != nil {
			return fmt.Errorf("Method: %w", err)
		}
	}
	if r.TenderedAmount != nil {
		if err := r.TenderedAmount.Validate(); err != nil {
			return fmt.Errorf("TenderedAmount: %w", err)
		}
	}
	if r.ReturnedAmount != nil {
		if err := r.ReturnedAmount.Validate(); err != nil {
			return fmt.Errorf("ReturnedAmount: %w", err)
		}
	}
	if r.Amount != nil {
		if err := r.Amount.Validate(); err != nil {
			return fmt.Errorf("Amount: %w", err)
		}
	}
	if r.PaymentIdentifier != nil {
		if err := r.PaymentIdentifier.Validate(); err != nil {
			return fmt.Errorf("PaymentIdentifier: %w", err)
		}
	}
	for i, item := range r.Allocation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Allocation[%d]: %w", i, err)
		}
	}
	if r.FormCode != nil {
		if err := r.FormCode.Validate(); err != nil {
			return fmt.Errorf("FormCode: %w", err)
		}
	}
	for i, item := range r.ProcessNote {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ProcessNote[%d]: %w", i, err)
		}
	}
	return nil
}

type PaymentReconciliationAllocation struct {
	Id                    *string          `json:"id,omitempty" bson:"id,omitempty"`                                          // Unique id for inter-element referencing
	Identifier            *Identifier      `json:"identifier,omitempty" bson:"identifier,omitempty"`                          // Business identifier of the payment detail
	Predecessor           *Identifier      `json:"predecessor,omitempty" bson:"predecessor,omitempty"`                        // Business identifier of the prior payment detail
	Target                *Reference       `json:"target,omitempty" bson:"target,omitempty"`                                  // Subject of the payment
	TargetItemString      *string          `json:"targetItemString,omitempty" bson:"target_item_string,omitempty"`            // Sub-element of the subject
	TargetItemIdentifier  *Identifier      `json:"targetItemIdentifier,omitempty" bson:"target_item_identifier,omitempty"`    // Sub-element of the subject
	TargetItemPositiveInt *int             `json:"targetItemPositiveInt,omitempty" bson:"target_item_positive_int,omitempty"` // Sub-element of the subject
	Encounter             *Reference       `json:"encounter,omitempty" bson:"encounter,omitempty"`                            // Applied-to encounter
	Account               *Reference       `json:"account,omitempty" bson:"account,omitempty"`                                // Applied-to account
	Type                  *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`                                      // Category of payment
	Submitter             *Reference       `json:"submitter,omitempty" bson:"submitter,omitempty"`                            // Submitter of the request
	Response              *Reference       `json:"response,omitempty" bson:"response,omitempty"`                              // Response committing to a payment
	Date                  *string          `json:"date,omitempty" bson:"date,omitempty"`                                      // Date of commitment to pay
	Responsible           *Reference       `json:"responsible,omitempty" bson:"responsible,omitempty"`                        // Contact for the response
	Payee                 *Reference       `json:"payee,omitempty" bson:"payee,omitempty"`                                    // Recipient of the payment
	Amount                *Money           `json:"amount,omitempty" bson:"amount,omitempty"`                                  // Amount allocated to this payable
	NoteNumber            []int            `json:"noteNumber,omitempty" bson:"note_number,omitempty"`                         // Applicable note numbers
}

func (r *PaymentReconciliationAllocation) Validate() error {
	if r.Identifier != nil {
		if err := r.Identifier.Validate(); err != nil {
			return fmt.Errorf("Identifier: %w", err)
		}
	}
	if r.Predecessor != nil {
		if err := r.Predecessor.Validate(); err != nil {
			return fmt.Errorf("Predecessor: %w", err)
		}
	}
	if r.Target != nil {
		if err := r.Target.Validate(); err != nil {
			return fmt.Errorf("Target: %w", err)
		}
	}
	if r.TargetItemIdentifier != nil {
		if err := r.TargetItemIdentifier.Validate(); err != nil {
			return fmt.Errorf("TargetItemIdentifier: %w", err)
		}
	}
	if r.Encounter != nil {
		if err := r.Encounter.Validate(); err != nil {
			return fmt.Errorf("Encounter: %w", err)
		}
	}
	if r.Account != nil {
		if err := r.Account.Validate(); err != nil {
			return fmt.Errorf("Account: %w", err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Submitter != nil {
		if err := r.Submitter.Validate(); err != nil {
			return fmt.Errorf("Submitter: %w", err)
		}
	}
	if r.Response != nil {
		if err := r.Response.Validate(); err != nil {
			return fmt.Errorf("Response: %w", err)
		}
	}
	if r.Responsible != nil {
		if err := r.Responsible.Validate(); err != nil {
			return fmt.Errorf("Responsible: %w", err)
		}
	}
	if r.Payee != nil {
		if err := r.Payee.Validate(); err != nil {
			return fmt.Errorf("Payee: %w", err)
		}
	}
	if r.Amount != nil {
		if err := r.Amount.Validate(); err != nil {
			return fmt.Errorf("Amount: %w", err)
		}
	}
	return nil
}

type PaymentReconciliationProcessNote struct {
	Id     *string          `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Class  *CodeableConcept `json:"class,omitempty" bson:"class,omitempty"`   // Business kind of note
	Number *int             `json:"number,omitempty" bson:"number,omitempty"` // Note instance identifier
	Type   *string          `json:"type,omitempty" bson:"type,omitempty"`     // display | print | printoper
	Text   *string          `json:"text,omitempty" bson:"text,omitempty"`     // Note explanatory text
}

func (r *PaymentReconciliationProcessNote) Validate() error {
	if r.Class != nil {
		if err := r.Class.Validate(); err != nil {
			return fmt.Errorf("Class: %w", err)
		}
	}
	return nil
}
