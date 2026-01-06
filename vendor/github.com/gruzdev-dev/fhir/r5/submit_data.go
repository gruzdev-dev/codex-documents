package models

// The submit-data operation is used to submit data-of-interest for one or more measures for one or more subjects. Each submitted Bundle **SHOULD** contain resources for a single subject and **SHOULD** contain all of the MeasureReports and data of interest related to that subject. There is no expectation that the submitted data represents all the data-of-interest, only that all the data submitted is relevant to the calculation of the measure for a particular subject or population. The dataUpdateType element of the MeasureReport resource is used to indicate whether the data being submitted is a snapshot or incremental update. Additional guidance about data exchange for quality reporting can be found in the Data Exchange for Quality Measures implementation guide. Note that the use of the [X-Provenance header data](https://hl7.org/fhir/6.0.0-ballot3/provenance.html#header) with data that establishes provenance being submitted/collected **SHOULD** be supported. This provides the capability for associating the provider with the data submitted through the $submit-data transaction. If the X-Provenance header is used it should be consistent with the reporter element in the DEQM Data Exchange MeasureReport Profile. This operation is purposefully not allowed on a Measure instance because the MeasureReport included in the Bundle specifies the measure.\n\nNOTE: This operation is being deprecated in favor of just posting the bundles to the root of the server, either as collection or transaction bundles.
type SubmitData struct {
}

func (r *SubmitData) Validate() error {
	return nil
}
