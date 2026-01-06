package models

// This operation is used to request the removal of all current and historical versions for all resources in a patient compartment.  The result will be an OperationOutcome with results and/or details about execution.  Following are some common 'issue-type' values: - 'success' the request has been completed to the server's satisfaction - the patient and associated resources are no longer accessible - 'incomplete' the request is partially complete, but additional processing will continue (e.g., the server is continuing to clean out resources)  When supported, it is recommended (though not required) to support an [Asynchronous Request Pattern](async.html).  Note that the deletion of resources typically involves many policy decisions.  Implementers are expected to use this operation in conjunction with their policies for such a request - e.g., soft vs. hard delete, audibility/traceability, evaluation of referential integrity, etc.
type Purge struct {
}

func (r *Purge) Validate() error {
	return nil
}
