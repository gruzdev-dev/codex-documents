package models

// The evaluate operation processes the given Measure(s) to produce the corresponding MeasureReport(s). This operation expects that Measure resources used have a computable representation. The value of title elements in the resulting [MeasureReport](clinicalreasoning-quality-reporting.html#measure-report) should be copied from the corresponding elements on the Measure.
type Evaluate struct {
}

func (r *Evaluate) Validate() error {
	return nil
}
