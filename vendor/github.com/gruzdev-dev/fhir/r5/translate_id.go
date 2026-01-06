package models

// This operation returns an identifier of the target type. The operation takes 5 parameters:       * a source identifier value - either a URI, an OID, or a v2 table 0396 (other) code   *  a code for what type of identifier the source identifier is       * a code for what kind of identifier is desired (URI, OID, v2 table 0396 identifier)       * an optional parameter preferredOnly for whether only the preferred identifier is desired       * an optional date to return only identifiers that have a validity period that includes that date     and returns either the requested identifier(s), or an HTTP errors response with an OperationOutcome because either the provided identifier was not recognized, or the requested identiifer type is not known.
type TranslateId struct {
}

func (r *TranslateId) Validate() error {
	return nil
}
