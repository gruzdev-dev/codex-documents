package models

// Attachment Type: For referring to data content defined in other formats.
type Attachment struct {
	Id          *string  `json:"id,omitempty" bson:"id,omitempty"`                    // Unique id for inter-element referencing
	ContentType *string  `json:"contentType,omitempty" bson:"content_type,omitempty"` // Mime type of the content, with charset etc.
	Language    *string  `json:"language,omitempty" bson:"language,omitempty"`        // Human language of the content (BCP-47)
	Data        *string  `json:"data,omitempty" bson:"data,omitempty"`                // Data inline, base64ed
	Url         *string  `json:"url,omitempty" bson:"url,omitempty"`                  // Uri where the data can be found
	Size        *int64   `json:"size,omitempty" bson:"size,omitempty"`                // Number of bytes of content (if url provided)
	Hash        *string  `json:"hash,omitempty" bson:"hash,omitempty"`                // Hash of the data (sha-1, base64ed)
	Title       *string  `json:"title,omitempty" bson:"title,omitempty"`              // Label to display in place of the data
	Creation    *string  `json:"creation,omitempty" bson:"creation,omitempty"`        // Date attachment was first created
	Height      *int     `json:"height,omitempty" bson:"height,omitempty"`            // Height of the image in pixels (photo/video)
	Width       *int     `json:"width,omitempty" bson:"width,omitempty"`              // Width of the image in pixels (photo/video)
	Frames      *int     `json:"frames,omitempty" bson:"frames,omitempty"`            // Number of frames if > 1 (photo)
	Duration    *float64 `json:"duration,omitempty" bson:"duration,omitempty"`        // Length in seconds (audio / video)
	Pages       *int     `json:"pages,omitempty" bson:"pages,omitempty"`              // Number of printed pages
}

func (r *Attachment) Validate() error {
	return nil
}
