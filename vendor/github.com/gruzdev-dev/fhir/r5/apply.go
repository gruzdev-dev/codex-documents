package models

// The apply operation applies a PlanDefinition to a given subject or group of subjects, instantiating applicable actions and returning the results as bundles of request resources.
type Apply struct {
}

func (r *Apply) Validate() error {
	return nil
}
