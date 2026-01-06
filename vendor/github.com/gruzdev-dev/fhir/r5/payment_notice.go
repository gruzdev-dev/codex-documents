package models

import (
	"encoding/json"
	"fmt"
)

// This resource provides the status of the payment for goods and services rendered, and the request and response resource references.
type PaymentNotice struct {
	Id            *string           `json:"id,omitempty" bson:"id,omitempty"`                        // Logical id of this artifact
	Meta          *Meta             `json:"meta,omitempty" bson:"meta,omitempty"`                    // Metadata about the resource
	ImplicitRules *string           `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"` // A set of rules under which this content was created
	Language      *string           `json:"language,omitempty" bson:"language,omitempty"`            // Language of the resource content
	Text          *Narrative        `json:"text,omitempty" bson:"text,omitempty"`                    // Text summary of the resource, for human interpretation
	Contained     []json.RawMessage `json:"contained,omitempty" bson:"contained,omitempty"`          // Contained, inline Resources
	Identifier    []Identifier      `json:"identifier,omitempty" bson:"identifier,omitempty"`        // Business Identifier for the payment notice
	Status        string            `json:"status" bson:"status"`                                    // active | cancelled | draft | entered-in-error
	StatusReason  *string           `json:"statusReason,omitempty" bson:"status_reason,omitempty"`   // Reason for status change
	Request       *Reference        `json:"request,omitempty" bson:"request,omitempty"`              // Request reference
	Response      *Reference        `json:"response,omitempty" bson:"response,omitempty"`            // Response reference
	Created       string            `json:"created" bson:"created"`                                  // Creation date
	Reporter      *Reference        `json:"reporter,omitempty" bson:"reporter,omitempty"`            // Responsible practitioner
	Payment       *Reference        `json:"payment,omitempty" bson:"payment,omitempty"`              // Payment reference
	PaymentDate   *string           `json:"paymentDate,omitempty" bson:"payment_date,omitempty"`     // Payment or clearing date
	Payee         *Reference        `json:"payee,omitempty" bson:"payee,omitempty"`                  // Party being paid
	Recipient     *Reference        `json:"recipient" bson:"recipient"`                              // Party being notified
	Amount        *Money            `json:"amount" bson:"amount"`                                    // Monetary amount of the payment
	PaymentStatus *CodeableConcept  `json:"paymentStatus,omitempty" bson:"payment_status,omitempty"` // Issued or cleared Status of the payment
}

func (r *PaymentNotice) Validate() error {
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
	if r.Request != nil {
		if err := r.Request.Validate(); err != nil {
			return fmt.Errorf("Request: %w", err)
		}
	}
	if r.Response != nil {
		if err := r.Response.Validate(); err != nil {
			return fmt.Errorf("Response: %w", err)
		}
	}
	if r.Created == emptyString {
		return fmt.Errorf("field 'Created' is required")
	}
	if r.Reporter != nil {
		if err := r.Reporter.Validate(); err != nil {
			return fmt.Errorf("Reporter: %w", err)
		}
	}
	if r.Payment != nil {
		if err := r.Payment.Validate(); err != nil {
			return fmt.Errorf("Payment: %w", err)
		}
	}
	if r.Payee != nil {
		if err := r.Payee.Validate(); err != nil {
			return fmt.Errorf("Payee: %w", err)
		}
	}
	if r.Recipient == nil {
		return fmt.Errorf("field 'Recipient' is required")
	}
	if r.Recipient != nil {
		if err := r.Recipient.Validate(); err != nil {
			return fmt.Errorf("Recipient: %w", err)
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
	if r.PaymentStatus != nil {
		if err := r.PaymentStatus.Validate(); err != nil {
			return fmt.Errorf("PaymentStatus: %w", err)
		}
	}
	return nil
}
