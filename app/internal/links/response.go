package links

type ErrorLinksResponse struct {
	Error string `json:"error"`
}

type SuccessAddLinksResponse struct {
	Links []LinkData `json:"links"`
}

type SuccessGetLinkResponse struct {
	Link LinkData `json:"link"`
}

type SuccessDeleteLinkResponse struct {
	Message string `json:"message"`
}

type SuccessGetLinksResponse struct {
	Links []LinkData `json:"links"`
}
