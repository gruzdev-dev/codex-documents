package models

// This operation is used to return all the information related to one or more products described in the resource or context on which this operation is invoked. The response is a bundle of type "searchset". At a minimum, the product resource(s) itself is returned, along with any other resources that the server has that are related to the products(s), and that are available for the given user. This is typically the marketing authorizations, ingredients, packages, therapeutic indications and so on. The server also returns whatever resources are needed to support the records - e.g. linked organizations, document references etc.
type Everything struct {
}

func (r *Everything) Validate() error {
	return nil
}
