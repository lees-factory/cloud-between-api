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
	steps, err := ctrl.diagnosisService.GetSteps(c.Request.Context(), locale)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]stepResponse, len(steps))
	for i, s := range steps {
		questions := make([]questionResponse, len(s.Questions))
		for j, q := range s.Questions {
			options := make([]optionResponse, len(q.Options))
			for k, o := range q.Options {
				options[k] = optionResponse{
					Text:      o.Text,
					CloudType: o.CloudType,
				}
			}
			questions[j] = questionResponse{
				ID:      q.ID,
				Text:    q.QuestionText,
				Options: options,
			}
		}
		response[i] = stepResponse{
			ID:        s.ID,
			Title:     s.Title,
			Emoji:     s.Emoji,
			Questions: questions,
		}
	}

	c.JSON(http.StatusOK, response)
}

func (ctrl *DiagnosisController) Analyze(c *gin.Context) {
	var req struct {
		Answers []diagnosis.UserAnswer `json:"answers"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

type stepResponse struct {
	ID        int                `json:"id"`
	Title     string             `json:"title"`
	Emoji     string             `json:"emoji"`
	Questions []questionResponse `json:"questions"`
}

type questionResponse struct {
	ID      int              `json:"id"`
	Text    string           `json:"text"`
	Options []optionResponse `json:"options"`
}

type optionResponse struct {
	Text      string `json:"text"`
	CloudType string `json:"cloudType"`
}
