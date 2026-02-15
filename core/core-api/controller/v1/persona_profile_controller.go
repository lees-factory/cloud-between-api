package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"io.lees.cloud-between/core/core-domain/persona"
)

type PersonaProfileController struct {
	service *persona.PersonaProfileService
}

func NewPersonaProfileController(service *persona.PersonaProfileService) *PersonaProfileController {
	return &PersonaProfileController{service: service}
}

func (ctrl *PersonaProfileController) GetProfiles(c *gin.Context) {
	locale := c.DefaultQuery("locale", "ko")

	profiles, err := ctrl.service.GetProfiles(c.Request.Context(), locale)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := make([]personaProfileResponse, len(profiles))
	for i, p := range profiles {
		responses[i] = toPersonaResponse(p)
	}

	c.JSON(http.StatusOK, responses)
}

func (ctrl *PersonaProfileController) GetProfile(c *gin.Context) {
	typeKey := c.Param("typeKey")
	locale := c.DefaultQuery("locale", "ko")

	profile, err := ctrl.service.GetProfile(c.Request.Context(), typeKey, locale)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "persona profile not found"})
		return
	}

	c.JSON(http.StatusOK, toPersonaResponse(*profile))
}

type personaProfileResponse struct {
	TypeKey   string   `json:"typeKey"`
	Name      string   `json:"name"`
	Emoji     string   `json:"emoji"`
	Subtitle  string   `json:"subtitle"`
	Keywords  []string `json:"keywords"`
	Lore      string   `json:"lore"`
	Strengths []string `json:"strengths"`
	Shadows   []string `json:"shadows"`
}

func toPersonaResponse(p persona.PersonaProfile) personaProfileResponse {
	return personaProfileResponse{
		TypeKey:   p.TypeKey,
		Name:      p.Name,
		Emoji:     p.Emoji,
		Subtitle:  p.Subtitle,
		Keywords:  p.Keywords,
		Lore:      p.Lore,
		Strengths: p.Strengths,
		Shadows:   p.Shadows,
	}
}