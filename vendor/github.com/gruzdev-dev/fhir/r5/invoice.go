package models

import (
	"encoding/json"
	"fmt"
)

// Invoice containing collected ChargeItems from an Account with calculated individual and total price for Billing purpose.
type Invoice struct {
	Id                  *string              `json:"id,omitempty" bson:"id,omitempty"`                                     // Logical id of this artifact
	Meta                *Meta                `json:"meta,omitempty" bson:"meta,omitempty"`                                 // Metadata about the resource
	ImplicitRules       *string              `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`              // A set of rules under which this content was created
	Language            *string              `json:"language,omitempty" bson:"language,omitempty"`                         // Language of the resource content
	Text                *Narrative           `json:"text,omitempty" bson:"text,omitempty"`                                 // Text summary of the resource, for human interpretation
	Contained           []json.RawMessage    `json:"contained,omitempty" bson:"contained,omitempty"`                       // Contained, inline Resources
	Identifier          []Identifier         `json:"identifier,omitempty" bson:"identifier,omitempty"`                     // Business Identifier for item
	Status              string               `json:"status" bson:"status"`                                                 // draft | issued | balanced | cancelled | entered-in-error
	CancelledReason     *string              `json:"cancelledReason,omitempty" bson:"cancelled_reason,omitempty"`          // Reason for cancellation of this Invoice
	Type                *CodeableConcept     `json:"type,omitempty" bson:"type,omitempty"`                                 // Type of Invoice
	Subject             *Reference           `json:"subject,omitempty" bson:"subject,omitempty"`                           // Recipient(s) of goods and services
	Recipient           *Reference           `json:"recipient,omitempty" bson:"recipient,omitempty"`                       // Recipient of this invoice
	Creation            *string              `json:"creation,omitempty" bson:"creation,omitempty"`                         // When posted
	PeriodDate          *string              `json:"periodDate,omitempty" bson:"period_date,omitempty"`                    // Billing date or period
	PeriodPeriod        *Period              `json:"periodPeriod,omitempty" bson:"period_period,omitempty"`                // Billing date or period
	Participant         []InvoiceParticipant `json:"participant,omitempty" bson:"participant,omitempty"`                   // Participant in creation of this Invoice
	Issuer              *Reference           `json:"issuer,omitempty" bson:"issuer,omitempty"`                             // Issuing entity
	Account             *Reference           `json:"account,omitempty" bson:"account,omitempty"`                           // Account that is being balanced
	LineItem            []InvoiceLineItem    `json:"lineItem,omitempty" bson:"line_item,omitempty"`                        // Line items of this Invoice
	TotalPriceComponent []MonetaryComponent  `json:"totalPriceComponent,omitempty" bson:"total_price_component,omitempty"` // Components of Invoice total
	TotalNet            *Money               `json:"totalNet,omitempty" bson:"total_net,omitempty"`                        // Net total of this Invoice
	TotalGross          *Money               `json:"totalGross,omitempty" bson:"total_gross,omitempty"`                    // Gross total of this Invoice
	PaymentTerms        *string              `json:"paymentTerms,omitempty" bson:"payment_terms,omitempty"`                // Payment details
	Note                []Annotation         `json:"note,omitempty" bson:"note,omitempty"`                                 // Comments made about the invoice
}

func (r *Invoice) Validate() error {
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
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	if r.Recipient != nil {
		if err := r.Recipient.Validate(); err != nil {
			return fmt.Errorf("Recipient: %w", err)
		}
	}
	if r.PeriodPeriod != nil {
		if err := r.PeriodPeriod.Validate(); err != nil {
			return fmt.Errorf("PeriodPeriod: %w", err)
		}
	}
	for i, item := range r.Participant {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Participant[%d]: %w", i, err)
		}
	}
	if r.Issuer != nil {
		if err := r.Issuer.Validate(); err != nil {
			return fmt.Errorf("Issuer: %w", err)
		}
	}
	if r.Account != nil {
		if err := r.Account.Validate(); err != nil {
			return fmt.Errorf("Account: %w", err)
		}
	}
	for i, item := range r.LineItem {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("LineItem[%d]: %w", i, err)
		}
	}
	for i, item := range r.TotalPriceComponent {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("TotalPriceComponent[%d]: %w", i, err)
		}
	}
	if r.TotalNet != nil {
		if err := r.TotalNet.Validate(); err != nil {
			return fmt.Errorf("TotalNet: %w", err)
		}
	}
	if r.TotalGross != nil {
		if err := r.TotalGross.Validate(); err != nil {
			return fmt.Errorf("TotalGross: %w", err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}

type InvoiceParticipant struct {
	Id    *string          `json:"id,omitempty" bson:"id,omitempty"`     // Unique id for inter-element referencing
	Role  *CodeableConcept `json:"role,omitempty" bson:"role,omitempty"` // Type of involvement in creation of this Invoice
	Actor *Reference       `json:"actor" bson:"actor"`                   // Individual who was involved
}

func (r *InvoiceParticipant) Validate() error {
	if r.Role != nil {
		if err := r.Role.Validate(); err != nil {
			return fmt.Errorf("Role: %w", err)
		}
	}
	if r.Actor == nil {
		return fmt.Errorf("field 'Actor' is required")
	}
	if r.Actor != nil {
		if err := r.Actor.Validate(); err != nil {
			return fmt.Errorf("Actor: %w", err)
		}
	}
	return nil
}

type InvoiceLineItem struct {
	Id                        *string             `json:"id,omitempty" bson:"id,omitempty"`                              // Unique id for inter-element referencing
	Sequence                  *int                `json:"sequence,omitempty" bson:"sequence,omitempty"`                  // Sequence number of line item
	ServicedDate              *string             `json:"servicedDate,omitempty" bson:"serviced_date,omitempty"`         // Service data or period
	ServicedPeriod            *Period             `json:"servicedPeriod,omitempty" bson:"serviced_period,omitempty"`     // Service data or period
	ChargeItemReference       *Reference          `json:"chargeItemReference" bson:"charge_item_reference"`              // Reference to ChargeItem containing details of this line item or an inline billing code
	ChargeItemCodeableConcept *CodeableConcept    `json:"chargeItemCodeableConcept" bson:"charge_item_codeable_concept"` // Reference to ChargeItem containing details of this line item or an inline billing code
	PriceComponent            []MonetaryComponent `json:"priceComponent,omitempty" bson:"price_component,omitempty"`     // Components of total line item price
}

func (r *InvoiceLineItem) Validate() error {
	if r.ServicedPeriod != nil {
		if err := r.ServicedPeriod.Validate(); err != nil {
			return fmt.Errorf("ServicedPeriod: %w", err)
		}
	}
	if r.ChargeItemReference == nil {
		return fmt.Errorf("field 'ChargeItemReference' is required")
	}
	if r.ChargeItemReference != nil {
		if err := r.ChargeItemReference.Validate(); err != nil {
			return fmt.Errorf("ChargeItemReference: %w", err)
		}
	}
	if r.ChargeItemCodeableConcept == nil {
		return fmt.Errorf("field 'ChargeItemCodeableConcept' is required")
	}
	if r.ChargeItemCodeableConcept != nil {
		if err := r.ChargeItemCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ChargeItemCodeableConcept: %w", err)
		}
	}
	for i, item := range r.PriceComponent {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PriceComponent[%d]: %w", i, err)
		}
	}
	return nil
}
