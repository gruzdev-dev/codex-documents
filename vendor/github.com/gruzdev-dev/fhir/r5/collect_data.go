package models

// The collect-data operation is used to collect the data-of-interest for the given measure. Note that the use of the [X-Provenance header data](provenance.html#header) with data that establishes provenance being submitted/collected **SHOULD** be supported.  This provides the capability for associating the provider with the data submitted through the $collect-data transaction. If the X-Provenance header is used it should be consistent with the 'reporter' element in the DEQM Data Exchange MeasureReport Profile.
type CollectData struct {
}

func (r *CollectData) Validate() error {
	return nil
}
