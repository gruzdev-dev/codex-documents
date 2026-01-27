package domain

type ShareRequest struct {
	ResourceIDs []string
	TTLSeconds  int64
}

type ShareResponse struct {
	Token       string
	ResourceURL string
}

type SharedResourcesResponse struct {
	Observations       []string
	DocumentReferences []string
}
