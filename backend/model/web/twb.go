package web

import "mime/multipart"

type TwbCreateRequest struct {
	Frame *multipart.FileHeader
}

type TwbGetRequest struct {
	ID string
}

type TwbResponse struct {
	Path string `json:"path"`
}
