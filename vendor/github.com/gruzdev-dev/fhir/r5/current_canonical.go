package models

// Returns the most current version of the canonical resource with the specified url available on the server.  It optionally also allows filtering to only expose the most current version with a particular status or set of statuses.  Note that 'current' is determined by comparing version values using the specified versionAlgorithm, NOT by looking at lastUpdated.
type CurrentCanonical struct {
}

func (r *CurrentCanonical) Validate() error {
	return nil
}
