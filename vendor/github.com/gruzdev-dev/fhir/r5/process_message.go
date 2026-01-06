package models

// This operation accepts a message, processes it according to the definition of the event in the message header, and returns one or more response messages.    In addition to processing the message event, a server may choose to retain all or some the resources and make them available on a RESTful interface, but is not required to do so.
type ProcessMessage struct {
}

func (r *ProcessMessage) Validate() error {
	return nil
}
