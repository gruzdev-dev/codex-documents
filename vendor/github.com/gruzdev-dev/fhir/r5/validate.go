package models

// The validate operation checks whether the attached content would be acceptable either generally, as a create, an update or as a delete to an existing resource.  The action the server takes depends on the mode parameter:    * [mode not provided]: The server checks the content of the resource against any schema, constraint rules, and other general terminology rules  * create: The server checks the content, and then checks that the content would be acceptable as a create (e.g. that the content would not violate any uniqueness constraints)  * update: The server checks the content, and then checks that it would accept it as an update against the nominated specific resource (e.g. that there are no changes to immutable fields the server does not allow to change, and checking version integrity if appropriate)  * delete: The server ignores the content, and checks that the nominated resource is allowed to be deleted (e.g. checking referential integrity rules)    Modes update and delete can only be used when the operation is invoked at the resource instance level.   The return from this operation is an [OperationOutcome](operationoutcome.html)  Note that this operation is not the only way to validate resources - see [Validating Resources](validation.html) for further information.
type Validate struct {
}

func (r *Validate) Validate() error {
	return nil
}
