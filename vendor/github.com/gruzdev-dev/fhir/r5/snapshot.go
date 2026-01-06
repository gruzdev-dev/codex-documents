package models

// Generates a [StructureDefinition](structuredefinition.html) instance  with  a snapshot, based on a differential in a specified [StructureDefinition](structuredefinition.html).     If the operation is not called at the instance level, either *definition* or *url* 'in' parameters must be provided. If more than one is specified, servers may raise an error or may resolve with the parameter of their choice. If called at the instance level, these parameters will be ignored. Snapshot generation is dependent on profiles that are referenced - if those profiles change, the snapshot can change. For a frozen package, $snapshot is idempotent.
type Snapshot struct {
}

func (r *Snapshot) Validate() error {
	return nil
}
