package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"io.lees.cloud-between/core/core-domain/chemistry"
)

type ChemistryController struct {
	service *chemistry.ChemistryService
}

func NewChemistryController(service *chemistry.ChemistryService) *ChemistryController {
	return &ChemistryController{service: service}
}

func (ctrl *ChemistryController) GetChemistry(c *gin.Context) {
	type1 := c.Query("type1")
	type2 := c.Query("type2")

	if type1 == "" || type2 == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type1 and type2 are required"})
		return
	}

	result, err := ctrl.service.GetChemistry(c.Request.Context(), type1, type2)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "chemistry not found"})
		return
	}

	c.JSON(http.StatusOK, toChemistryResponse(*result))
}

func (ctrl *ChemistryController) GetAllChemistries(c *gin.Context) {
	results, err := ctrl.service.GetAllChemistries(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := make([]chemistryResponse, len(results))
	for i, r := range results {
		responses[i] = toChemistryResponse(r)
	}

	c.JSON(http.StatusOK, responses)
}

type chemistryResponse struct {
	PersonaType1   string                 `json:"personaType1"`
	PersonaType2   string                 `json:"personaType2"`
	SkyName        string                 `json:"skyName"`
	SkyNameKo      string                 `json:"skyNameKo,omitempty"`
	Phenomenon     string                 `json:"phenomenon"`
	Narrative      string                 `json:"narrative"`
	Warning        string                 `json:"warning,omitempty"`
	PhenomenonName map[string]string      `json:"phenomenonName,omitempty"`
	VibeTags       map[string][]string    `json:"vibeTags,omitempty"`
	StoryBeats     map[string]interface{} `json:"storyBeats,omitempty"`
	Premium        map[string]interface{} `json:"premium,omitempty"`
}

func toChemistryResponse(c chemistry.Chemistry) chemistryResponse {
	resp := chemistryResponse{
		PersonaType1: c.PersonaType1,
		PersonaType2: c.PersonaType2,
		SkyName:      c.SkyName,
		SkyNameKo:    c.SkyNameKo,
		Phenomenon:   c.Phenomenon,
		Narrative:    c.Narrative,
		Warning:      c.Warning,
	}

	if len(c.Content) > 0 {
		var content map[string]json.RawMessage
		if err := json.Unmarshal(c.Content, &content); err == nil {
			if raw, ok := content["phenomenonName"]; ok {
				json.Unmarshal(raw, &resp.PhenomenonName)
			}
			if raw, ok := content["vibeTags"]; ok {
				json.Unmarshal(raw, &resp.VibeTags)
			}
			if raw, ok := content["storyBeats"]; ok {
				json.Unmarshal(raw, &resp.StoryBeats)
			}
			if raw, ok := content["premium"]; ok {
				json.Unmarshal(raw, &resp.Premium)
			}
		}
	}

	return resp
}
