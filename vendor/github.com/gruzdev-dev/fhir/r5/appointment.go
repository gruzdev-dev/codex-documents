package models

import (
	"encoding/json"
	"fmt"
)

// A booking of a healthcare event among patient(s), practitioner(s), related person(s) and/or device(s) for a specific date/time. This may result in one or more Encounter(s).
type Appointment struct {
	Id                     *string                         `json:"id,omitempty" bson:"id,omitempty"`                                          // Logical id of this artifact
	Meta                   *Meta                           `json:"meta,omitempty" bson:"meta,omitempty"`                                      // Metadata about the resource
	ImplicitRules          *string                         `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                   // A set of rules under which this content was created
	Language               *string                         `json:"language,omitempty" bson:"language,omitempty"`                              // Language of the resource content
	Text                   *Narrative                      `json:"text,omitempty" bson:"text,omitempty"`                                      // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage               `json:"contained,omitempty" bson:"contained,omitempty"`                            // Contained, inline Resources
	Identifier             []Identifier                    `json:"identifier,omitempty" bson:"identifier,omitempty"`                          // External Ids for this item
	Status                 string                          `json:"status" bson:"status"`                                                      // proposed | pending | booked | arrived | fulfilled | cancelled | noshow | entered-in-error | checked-in | waitlist
	CancellationReason     *CodeableConcept                `json:"cancellationReason,omitempty" bson:"cancellation_reason,omitempty"`         // The coded reason for the appointment being cancelled
	Class                  []CodeableConcept               `json:"class,omitempty" bson:"class,omitempty"`                                    // Classification when becoming an encounter
	ServiceCategory        []CodeableConcept               `json:"serviceCategory,omitempty" bson:"service_category,omitempty"`               // A broad categorization of the service that is to be performed during this appointment
	ServiceType            []CodeableReference             `json:"serviceType,omitempty" bson:"service_type,omitempty"`                       // The specific service that is to be performed during this appointment
	Specialty              []CodeableConcept               `json:"specialty,omitempty" bson:"specialty,omitempty"`                            // The specialty of a practitioner that would be required to perform the service requested in this appointment
	AppointmentType        *CodeableConcept                `json:"appointmentType,omitempty" bson:"appointment_type,omitempty"`               // The style of appointment or patient that has been booked in the slot (not service type)
	Reason                 []CodeableReference             `json:"reason,omitempty" bson:"reason,omitempty"`                                  // Reason this appointment is scheduled
	Priority               *CodeableConcept                `json:"priority,omitempty" bson:"priority,omitempty"`                              // Used to make informed decisions if needing to re-prioritize
	Description            *string                         `json:"description,omitempty" bson:"description,omitempty"`                        // Shown on a subject line in a meeting request, or appointment list
	Replaces               []Reference                     `json:"replaces,omitempty" bson:"replaces,omitempty"`                              // Appointment replaced by this Appointment
	VirtualService         []VirtualServiceDetail          `json:"virtualService,omitempty" bson:"virtual_service,omitempty"`                 // Connection details of a virtual service (e.g. conference call)
	SupportingInformation  []Reference                     `json:"supportingInformation,omitempty" bson:"supporting_information,omitempty"`   // Additional information to support the appointment
	PreviousAppointment    *Reference                      `json:"previousAppointment,omitempty" bson:"previous_appointment,omitempty"`       // The previous appointment in a series
	OriginatingAppointment *Reference                      `json:"originatingAppointment,omitempty" bson:"originating_appointment,omitempty"` // The originating appointment in a recurring set of appointments
	Start                  *string                         `json:"start,omitempty" bson:"start,omitempty"`                                    // When appointment is to take place
	End                    *string                         `json:"end,omitempty" bson:"end,omitempty"`                                        // When appointment is to conclude
	MinutesDuration        *int                            `json:"minutesDuration,omitempty" bson:"minutes_duration,omitempty"`               // Can be less than start/end (e.g. estimate)
	RequestedPeriod        []Period                        `json:"requestedPeriod,omitempty" bson:"requested_period,omitempty"`               // Potential date/time interval(s) requested to allocate the appointment within
	Slot                   []Reference                     `json:"slot,omitempty" bson:"slot,omitempty"`                                      // The slots that this appointment is filling
	Account                []Reference                     `json:"account,omitempty" bson:"account,omitempty"`                                // The set of accounts that may be used for billing for this Appointment
	Created                *string                         `json:"created,omitempty" bson:"created,omitempty"`                                // The date that this appointment was initially created
	CancellationDate       *string                         `json:"cancellationDate,omitempty" bson:"cancellation_date,omitempty"`             // When the appointment was cancelled
	Note                   []Annotation                    `json:"note,omitempty" bson:"note,omitempty"`                                      // Additional comments
	PatientInstruction     []CodeableReference             `json:"patientInstruction,omitempty" bson:"patient_instruction,omitempty"`         // Detailed information and instructions for the patient
	BasedOn                []Reference                     `json:"basedOn,omitempty" bson:"based_on,omitempty"`                               // The request this appointment is allocated to assess
	Subject                *Reference                      `json:"subject,omitempty" bson:"subject,omitempty"`                                // The patient or group associated with the appointment
	Participant            []AppointmentParticipant        `json:"participant" bson:"participant"`                                            // Participants involved in appointment
	RecurrenceId           *int                            `json:"recurrenceId,omitempty" bson:"recurrence_id,omitempty"`                     // The sequence number in the recurrence
	OccurrenceChanged      bool                            `json:"occurrenceChanged,omitempty" bson:"occurrence_changed,omitempty"`           // Indicates that this appointment varies from a recurrence pattern
	RecurrenceTemplate     []AppointmentRecurrenceTemplate `json:"recurrenceTemplate,omitempty" bson:"recurrence_template,omitempty"`         // Details of the recurrence pattern/template used to generate occurrences
}

func (r *Appointment) Validate() error {
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
	if r.CancellationReason != nil {
		if err := r.CancellationReason.Validate(); err != nil {
			return fmt.Errorf("CancellationReason: %w", err)
		}
	}
	for i, item := range r.Class {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Class[%d]: %w", i, err)
		}
	}
	for i, item := range r.ServiceCategory {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ServiceCategory[%d]: %w", i, err)
		}
	}
	for i, item := range r.ServiceType {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ServiceType[%d]: %w", i, err)
		}
	}
	for i, item := range r.Specialty {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Specialty[%d]: %w", i, err)
		}
	}
	if r.AppointmentType != nil {
		if err := r.AppointmentType.Validate(); err != nil {
			return fmt.Errorf("AppointmentType: %w", err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	if r.Priority != nil {
		if err := r.Priority.Validate(); err != nil {
			return fmt.Errorf("Priority: %w", err)
		}
	}
	for i, item := range r.Replaces {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Replaces[%d]: %w", i, err)
		}
	}
	for i, item := range r.VirtualService {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("VirtualService[%d]: %w", i, err)
		}
	}
	for i, item := range r.SupportingInformation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SupportingInformation[%d]: %w", i, err)
		}
	}
	if r.PreviousAppointment != nil {
		if err := r.PreviousAppointment.Validate(); err != nil {
			return fmt.Errorf("PreviousAppointment: %w", err)
		}
	}
	if r.OriginatingAppointment != nil {
		if err := r.OriginatingAppointment.Validate(); err != nil {
			return fmt.Errorf("OriginatingAppointment: %w", err)
		}
	}
	for i, item := range r.RequestedPeriod {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RequestedPeriod[%d]: %w", i, err)
		}
	}
	for i, item := range r.Slot {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Slot[%d]: %w", i, err)
		}
	}
	for i, item := range r.Account {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Account[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	for i, item := range r.PatientInstruction {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PatientInstruction[%d]: %w", i, err)
		}
	}
	for i, item := range r.BasedOn {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BasedOn[%d]: %w", i, err)
		}
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	if len(r.Participant) < 1 {
		return fmt.Errorf("field 'Participant' must have at least 1 elements")
	}
	for i, item := range r.Participant {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Participant[%d]: %w", i, err)
		}
	}
	for i, item := range r.RecurrenceTemplate {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RecurrenceTemplate[%d]: %w", i, err)
		}
	}
	return nil
}

type AppointmentRecurrenceTemplateMonthlyTemplate struct {
	Id             *string `json:"id,omitempty" bson:"id,omitempty"`                            // Unique id for inter-element referencing
	DayOfMonth     *int    `json:"dayOfMonth,omitempty" bson:"day_of_month,omitempty"`          // Recurs on a specific day of the month
	NthWeekOfMonth *Coding `json:"nthWeekOfMonth,omitempty" bson:"nth_week_of_month,omitempty"` // Indicates which week of the month the appointment should occur
	DayOfWeek      *Coding `json:"dayOfWeek,omitempty" bson:"day_of_week,omitempty"`            // Indicates which day of the week the appointment should occur
	MonthInterval  int     `json:"monthInterval" bson:"month_interval"`                         // Recurs every nth month
}

func (r *AppointmentRecurrenceTemplateMonthlyTemplate) Validate() error {
	if r.NthWeekOfMonth != nil {
		if err := r.NthWeekOfMonth.Validate(); err != nil {
			return fmt.Errorf("NthWeekOfMonth: %w", err)
		}
	}
	if r.DayOfWeek != nil {
		if err := r.DayOfWeek.Validate(); err != nil {
			return fmt.Errorf("DayOfWeek: %w", err)
		}
	}
	if r.MonthInterval == 0 {
		return fmt.Errorf("field 'MonthInterval' is required")
	}
	return nil
}

type AppointmentRecurrenceTemplateYearlyTemplate struct {
	Id           *string `json:"id,omitempty" bson:"id,omitempty"`  // Unique id for inter-element referencing
	YearInterval int     `json:"yearInterval" bson:"year_interval"` // Recurs every nth year
}

func (r *AppointmentRecurrenceTemplateYearlyTemplate) Validate() error {
	if r.YearInterval == 0 {
		return fmt.Errorf("field 'YearInterval' is required")
	}
	return nil
}

type AppointmentParticipant struct {
	Id       *string           `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Type     []CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`         // Role of participant in the appointment
	Period   *Period           `json:"period,omitempty" bson:"period,omitempty"`     // Participation period of the actor
	Actor    *Reference        `json:"actor,omitempty" bson:"actor,omitempty"`       // The individual, device, location, or service participating in the appointment
	Required bool              `json:"required,omitempty" bson:"required,omitempty"` // The participant is required to attend (optional when false)
	Status   string            `json:"status" bson:"status"`                         // accepted | declined | tentative | needs-action
}

func (r *AppointmentParticipant) Validate() error {
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	if r.Actor != nil {
		if err := r.Actor.Validate(); err != nil {
			return fmt.Errorf("Actor: %w", err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	return nil
}

type AppointmentRecurrenceTemplate struct {
	Id                    *string                                       `json:"id,omitempty" bson:"id,omitempty"`                                         // Unique id for inter-element referencing
	Timezone              *CodeableConcept                              `json:"timezone,omitempty" bson:"timezone,omitempty"`                             // The timezone of the occurrences
	RecurrenceType        *CodeableConcept                              `json:"recurrenceType" bson:"recurrence_type"`                                    // The frequency of the recurrence
	LastOccurrenceDate    *string                                       `json:"lastOccurrenceDate,omitempty" bson:"last_occurrence_date,omitempty"`       // The date when the recurrence should end
	OccurrenceCount       *int                                          `json:"occurrenceCount,omitempty" bson:"occurrence_count,omitempty"`              // The number of planned occurrences
	OccurrenceDate        []string                                      `json:"occurrenceDate,omitempty" bson:"occurrence_date,omitempty"`                // Specific dates for a recurring set of appointments (no template)
	WeeklyTemplate        *AppointmentRecurrenceTemplateWeeklyTemplate  `json:"weeklyTemplate,omitempty" bson:"weekly_template,omitempty"`                // Information about weekly recurring appointments
	MonthlyTemplate       *AppointmentRecurrenceTemplateMonthlyTemplate `json:"monthlyTemplate,omitempty" bson:"monthly_template,omitempty"`              // Information about monthly recurring appointments
	YearlyTemplate        *AppointmentRecurrenceTemplateYearlyTemplate  `json:"yearlyTemplate,omitempty" bson:"yearly_template,omitempty"`                // Information about yearly recurring appointments
	ExcludingDate         []string                                      `json:"excludingDate,omitempty" bson:"excluding_date,omitempty"`                  // Any dates that should be excluded from the series
	ExcludingRecurrenceId []int                                         `json:"excludingRecurrenceId,omitempty" bson:"excluding_recurrence_id,omitempty"` // Any recurrence IDs that should be excluded from the recurrence
}

func (r *AppointmentRecurrenceTemplate) Validate() error {
	if r.Timezone != nil {
		if err := r.Timezone.Validate(); err != nil {
			return fmt.Errorf("Timezone: %w", err)
		}
	}
	if r.RecurrenceType == nil {
		return fmt.Errorf("field 'RecurrenceType' is required")
	}
	if r.RecurrenceType != nil {
		if err := r.RecurrenceType.Validate(); err != nil {
			return fmt.Errorf("RecurrenceType: %w", err)
		}
	}
	if r.WeeklyTemplate != nil {
		if err := r.WeeklyTemplate.Validate(); err != nil {
			return fmt.Errorf("WeeklyTemplate: %w", err)
		}
	}
	if r.MonthlyTemplate != nil {
		if err := r.MonthlyTemplate.Validate(); err != nil {
			return fmt.Errorf("MonthlyTemplate: %w", err)
		}
	}
	if r.YearlyTemplate != nil {
		if err := r.YearlyTemplate.Validate(); err != nil {
			return fmt.Errorf("YearlyTemplate: %w", err)
		}
	}
	return nil
}

type AppointmentRecurrenceTemplateWeeklyTemplate struct {
	Id           *string `json:"id,omitempty" bson:"id,omitempty"`                      // Unique id for inter-element referencing
	Monday       bool    `json:"monday,omitempty" bson:"monday,omitempty"`              // Recurs on Mondays
	Tuesday      bool    `json:"tuesday,omitempty" bson:"tuesday,omitempty"`            // Recurs on Tuesday
	Wednesday    bool    `json:"wednesday,omitempty" bson:"wednesday,omitempty"`        // Recurs on Wednesday
	Thursday     bool    `json:"thursday,omitempty" bson:"thursday,omitempty"`          // Recurs on Thursday
	Friday       bool    `json:"friday,omitempty" bson:"friday,omitempty"`              // Recurs on Friday
	Saturday     bool    `json:"saturday,omitempty" bson:"saturday,omitempty"`          // Recurs on Saturday
	Sunday       bool    `json:"sunday,omitempty" bson:"sunday,omitempty"`              // Recurs on Sunday
	WeekInterval *int    `json:"weekInterval,omitempty" bson:"week_interval,omitempty"` // Recurs every nth week
}

func (r *AppointmentRecurrenceTemplateWeeklyTemplate) Validate() error {
	return nil
}
