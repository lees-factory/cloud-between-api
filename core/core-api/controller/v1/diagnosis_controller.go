package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"io.lees.cloud-between/core/core-domain/diagnosis"
)

type DiagnosisController struct {
	diagnosisService *diagnosis.DiagnosisService
}

func NewDiagnosisController(diagnosisService *diagnosis.DiagnosisService) *DiagnosisController {
	return &DiagnosisController{diagnosisService: diagnosisService}
}

func (ctrl *DiagnosisController) GetQuestions(c *gin.Context) {
	locale := c.DefaultQuery("locale", "ko")
	questions, err := ctrl.diagnosisService.GetQuestions(c.Request.Context(), locale)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, questions)
}

func (ctrl *DiagnosisController) Analyze(c *gin.Context) {
	var req struct {
		Answers []diagnosis.UserAnswer `json:"answers"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Member/Non-member detection (Mock)
	var userID *string
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		uid := "member-user-id"
		userID = &uid
	}

	result, err := ctrl.diagnosisService.CalculateResult(c.Request.Context(), userID, req.Answers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
