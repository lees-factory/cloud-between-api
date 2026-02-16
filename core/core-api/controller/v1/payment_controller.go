package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io.lees.cloud-between/core/core-api/controller/v1/request"
	"io.lees.cloud-between/core/core-api/controller/v1/response"
	"io.lees.cloud-between/core/core-domain/payment"
)

type PaymentController struct {
	paymentService *payment.PaymentService
}

func NewPaymentController(paymentService *payment.PaymentService) *PaymentController {
	return &PaymentController{paymentService: paymentService}
}

func (ctrl *PaymentController) CreateOrder(c *gin.Context) {
	var req request.CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userId"})
		return
	}

	orderID, approveURL, err := ctrl.paymentService.CreateOrder(c.Request.Context(), userID, req.Amount, req.Currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.CreatePaymentResponse{
		OrderID:    orderID,
		ApproveURL: approveURL,
	})
}

func (ctrl *PaymentController) CaptureOrder(c *gin.Context) {
	var req request.CapturePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.paymentService.CaptureOrder(c.Request.Context(), req.OrderID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "payment captured successfully"})
}

func (ctrl *PaymentController) CancelOrder(c *gin.Context) {
	var req request.CancelPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.paymentService.CancelOrder(c.Request.Context(), req.OrderID, req.Reason); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "payment cancellation recorded"})
}
