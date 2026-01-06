package models

// Execute a graphql statement on a since resource or against the entire system. See the [Using GraphQL with FHIR](graphql.html) page for further details.  For the purposes of graphQL compatibility, this operation can also be invoked using a POST with the graphQL as the body, or a JSON body (see [graphQL spec](http://graphql.org/) for details)
type Graphql struct {
}

func (r *Graphql) Validate() error {
	return nil
}
