package links

import "shortify/internal/db"

type LinkData struct {
	Id   int    `json:"id"`
	Url  string `json:"url" binding:"required,url"`
	Code string `json:"code"`
}

func MakeLinkFromShortLink(shortLink *db.ShortLink) LinkData {
	return LinkData{Id: int(shortLink.ID), Url: shortLink.OriginalUrl, Code: shortLink.Code}
}

type AddLinksRequest struct {
	Links []LinkData `json:"links" binding:"required"`
}
