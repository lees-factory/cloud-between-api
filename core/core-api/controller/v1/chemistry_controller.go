package v1

import (
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
	PersonaType1 string `json:"personaType1"`
	PersonaType2 string `json:"personaType2"`
	SkyName      string `json:"skyName"`
	Phenomenon   string `json:"phenomenon"`
	Narrative    string `json:"narrative"`
	Warning      string `json:"warning,omitempty"`
}

func toChemistryResponse(c chemistry.Chemistry) chemistryResponse {
	return chemistryResponse{
		PersonaType1: c.PersonaType1,
		PersonaType2: c.PersonaType2,
		SkyName:      c.SkyName,
		Phenomenon:   c.Phenomenon,
		Narrative:    c.Narrative,
		Warning:      c.Warning,
	}
}
