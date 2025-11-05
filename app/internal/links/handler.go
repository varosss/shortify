package links

import (
	"fmt"
	"net/http"
	"shortify/internal/cache"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type LinksHandler struct {
	linksService *LinksService
	redisClient  *redis.Client
}

func NewLinksHandler(conn *gorm.DB, redisAddr string) *LinksHandler {
	return &LinksHandler{linksService: NewLinksService(conn), redisClient: cache.NewRedisClient(redisAddr)}
}

func (h *LinksHandler) Redirect(c *gin.Context) {
	code := c.Param("short")

	url, err := h.redisClient.Get(c.Request.Context(), fmt.Sprintf("url_%s", code)).Result()
	if err == nil {
		c.Redirect(http.StatusFound, url)

		return
	}

	shortLink, err := h.linksService.shortLinksRepo.FindOneByCode(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorLinksResponse{Error: "short link doesn't exist"})

		return
	}

	h.redisClient.Set(c.Request.Context(), fmt.Sprintf("url_%s", code), shortLink.OriginalUrl, cache.URL_CACHE_DURATION)

	c.Redirect(http.StatusFound, shortLink.OriginalUrl)
}

func (h *LinksHandler) AddLinks(c *gin.Context) {
	var req AddLinksRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorLinksResponse{Error: err.Error()})

		return
	}

	userId := c.GetInt("user_id")

	links := h.linksService.AddLinks(c.Request.Context(), int(userId), req.Links)

	c.JSON(http.StatusOK, SuccessAddLinksResponse{Links: links})
}

func (h *LinksHandler) GetLink(c *gin.Context) {
	code := c.Param("short")
	userId := c.GetInt("user_id")

	link, err := h.linksService.GetLink(c.Request.Context(), code, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorLinksResponse{Error: err.Error()})

		return
	}

	c.JSON(http.StatusOK, SuccessGetLinkResponse{Link: *link})
}

func (h *LinksHandler) DeleteLink(c *gin.Context) {
	code := c.Param("short")
	userId := c.GetInt("user_id")

	err := h.linksService.DeleteLink(c.Request.Context(), code, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorLinksResponse{Error: err.Error()})

		return
	}

	c.JSON(http.StatusOK, SuccessDeleteLinkResponse{Message: "ok"})
}

func (h *LinksHandler) GetLinks(c *gin.Context) {
	userId := c.GetInt("user_id")

	links, err := h.linksService.GetLinks(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorLinksResponse{Error: err.Error()})

		return
	}

	c.JSON(http.StatusOK, SuccessGetLinksResponse{Links: links})
}
