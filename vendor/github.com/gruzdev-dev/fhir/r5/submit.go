package models

// This operation is used to submit an EligibilityRequest for assessment either as a single EligibilityRequest resource instance or as a Bundle containing the EligibilityRequest and other referenced resources, or Bundle containing a batch of EligibilityRequest resources, either as single EligibilityRequests resources or Bundle resources, for processing. The only input parameter is the single EligibilityRequest or Bundle resource and the only output is a single EligibilityResponse, Bundle of EligibilityResponses or an OperationOutcome resource.
type Submit struct {
}

func (r *Submit) Validate() error {
	return nil
}
