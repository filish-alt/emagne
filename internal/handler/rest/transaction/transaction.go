package transaction

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/filagot/emagne/internal/database/models/dto"
	"github.com/filagot/emagne/internal/handler/rest"
	"github.com/filagot/emagne/internal/handler/middleware"
	"github.com/filagot/emagne/internal/module/transaction"
)

type Handler struct {
	mod transaction.Module
}

func Init(mod transaction.Module) rest.TransactionHandler {
	return &Handler{mod: mod}
}

func (h *Handler) CreateTransaction(c *gin.Context) {
	var req dto.CreateTransaction
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tx, err := h.mod.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, tx)
}

func (h *Handler) GetTransaction(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	tx, err := h.mod.GetWithCategory(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tx)
}

func (h *Handler) ListTransactionsByCategory(c *gin.Context) {
	catIDStr := c.Param("categoryID")
	catID, err := uuid.Parse(catIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category id"})
		return
	}
	var limit int32 = 20
	var offset int32 = 0
	if v := c.Query("limit"); v != "" {
		if n, e := parseInt32(v); e == nil {
			limit = n
		}
	}
	if v := c.Query("offset"); v != "" {
		if n, e := parseInt32(v); e == nil {
			offset = n
		}
	}
	list, err := h.mod.ListByCategory(c.Request.Context(), catID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateTransactionStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req dto.UpdateTransactionStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	status := sql.NullString{String: req.Status, Valid: req.Status != ""}
	tx, err := h.mod.UpdateStatus(c.Request.Context(), id, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tx)
}

func (h *Handler) ConfirmTransaction(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	email := middleware.GetUserEmail(c)
	if email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user email not found"})
		return
	}
	tx, err := h.mod.ConfirmBySeller(c.Request.Context(), id, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tx)
}

func (h *Handler) MarkPaid(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	email := middleware.GetUserEmail(c)
	if email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user email not found"})
		return
	}
	tx, err := h.mod.MarkPaid(c.Request.Context(), id, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tx)
}

func (h *Handler) MarkShipped(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	email := middleware.GetUserEmail(c)
	if email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user email not found"})
		return
	}
	tx, err := h.mod.MarkShipped(c.Request.Context(), id, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tx)
}

func (h *Handler) MarkDelivered(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	email := middleware.GetUserEmail(c)
	if email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user email not found"})
		return
	}
	tx, err := h.mod.MarkDelivered(c.Request.Context(), id, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tx)
}

func (h *Handler) StartInspection(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	tx, err := h.mod.StartInspection(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tx)
}

func (h *Handler) CloseTransaction(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	tx, err := h.mod.MarkClosed(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tx)
}

func (h *Handler) AddTransactionAttribute(c *gin.Context) {
	idStr := c.Param("id")
	txID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction id"})
		return
	}
	var req dto.AddTransactionAttributeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attr, err := h.mod.AddAttribute(c.Request.Context(), txID, req.AttributeID, req.Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, attr)
}

func (h *Handler) ListTransactionAttributes(c *gin.Context) {
	idStr := c.Param("id")
	txID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction id"})
		return
	}
	list, err := h.mod.ListAttributes(c.Request.Context(), txID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) GetTransactionAttribute(c *gin.Context) {
	idStr := c.Param("id")
	txID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction id"})
		return
	}
	attrIDStr := c.Param("attrID")
	attrID, err := uuid.Parse(attrIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid attribute id"})
		return
	}
	attr, err := h.mod.GetAttribute(c.Request.Context(), txID, attrID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, attr)
}

func (h *Handler) DeleteTransactionAttributes(c *gin.Context) {
	idStr := c.Param("id")
	txID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction id"})
		return
	}
	if err := h.mod.DeleteAllAttributes(c.Request.Context(), txID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func parseInt32(s string) (int32, error) {
	var n int64
	var err error
	n, err = strconv.ParseInt(s, 10, 32)
	return int32(n), err
}
