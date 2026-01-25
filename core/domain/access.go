package domain

type GenerateTmpTokenRequest struct {
	Payload    map[string]string
	TtlSeconds int64
}

type GenerateTmpTokenResponse struct {
	TmpToken string
}
