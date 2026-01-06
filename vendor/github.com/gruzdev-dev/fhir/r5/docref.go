package models

// This operation is used to return all the references to documents related to a patient.    The operation requires a patient id and takes the optional input parameters:    - start date   - end date   - document type     - on demand     - profile    and returns a [Bundle](bundle.html) of type "searchset" containing [DocumentReference](documentreference.html) resources for the patient. If the server has or can create documents that are related to the patient, and that are available for the given user, the server returns the DocumentReference resources needed to support the records.  The principle intended use for this operation is to provide a provider or patient with access to their available document information.    This operation is *different* from a search by patient and type and date range because:    1. It is used to request a server to *generate* a document based on the specified parameters.    1. If no parameters are specified, the server SHALL return a DocumentReference to the patient's most current summary    1. If the server cannot *generate* a document based on the specified parameters, the operation will return an empty search bundle.    Unless the client indicates they are only interested in 'on-demand' documents using the on-demand parameter, the server SHOULD return DocumentReference instances for existing documents that meet the request parameters. In this regard, this operation is similar to a FHIR RESTful query.
type Docref struct {
}

func (r *Docref) Validate() error {
	return nil
}
