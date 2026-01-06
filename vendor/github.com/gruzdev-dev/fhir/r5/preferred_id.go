package models

// This operation returns the preferred identifiers for identifiers, and terminologies. The operation takes 2 parameters:       * a system identifier - either a URI, an OID, or a v2 table 0396 (other) code   * a code for what kind of identifier is desired (URI, OID, v2 table 0396 identifier)      and returns either the requested identifier, or an HTTP errors response with an OperationOutcome because either the provided identifier was not recognized, or the requested identiifer type is not known.      The principle use of this operation is when converting between v2, CDA and FHIR Identifier/CX/II and CodeableConcepts/C(N/W)E/CD but the operation may also find use when converting metadata such as profiles.
type PreferredId struct {
}

func (r *PreferredId) Validate() error {
	return nil
}
