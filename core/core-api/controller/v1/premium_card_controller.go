package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"io.lees.cloud-between/core/core-domain/premiumcard"
)

type PremiumCardController struct {
	service *premiumcard.PremiumCardService
}

func NewPremiumCardController(service *premiumcard.PremiumCardService) *PremiumCardController {
	return &PremiumCardController{service: service}
}

// GetAll returns all premium card templates
func (ctrl *PremiumCardController) GetAll(c *gin.Context) {
	results, err := ctrl.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toPremiumCardGroupedResponse(results))
}

// GetByCategory returns templates for a specific category
func (ctrl *PremiumCardController) GetByCategory(c *gin.Context) {
	category := c.Param("category")
	locale := c.DefaultQuery("locale", "ko")

	results, err := ctrl.service.GetByCategoryAndLocale(c.Request.Context(), category, locale)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	responses := make([]premiumCardResponse, len(results))
	for i, r := range results {
		responses[i] = toPremiumCardResponse(r)
	}
	c.JSON(http.StatusOK, responses)
}

type premiumCardResponse struct {
	ID       int             `json:"id"`
	Category string          `json:"category"`
	SubKey   string          `json:"subKey,omitempty"`
	Locale   string          `json:"locale,omitempty"`
	Content  json.RawMessage `json:"content"`
}

func toPremiumCardResponse(p premiumcard.PremiumCard) premiumCardResponse {
	return premiumCardResponse{
		ID:       p.ID,
		Category: p.Category,
		SubKey:   p.SubKey,
		Locale:   p.Locale,
		Content:  p.Content,
	}
}

// toPremiumCardGroupedResponse groups templates by category
func toPremiumCardGroupedResponse(cards []premiumcard.PremiumCard) map[string][]premiumCardResponse {
	grouped := make(map[string][]premiumCardResponse)
	for _, card := range cards {
		resp := toPremiumCardResponse(card)
		grouped[card.Category] = append(grouped[card.Category], resp)
	}
	return grouped
}
