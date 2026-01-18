package domain

type GetPresignedUrlsRequest struct {
	UserId      string
	ContentType string
	Size        int64
}

type PresignedUrlsResponse struct {
	FileId      string
	UploadUrl   string
	DownloadUrl string
}
