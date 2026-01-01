package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/filagot/emagne/internal/database/models/dto"
	"github.com/filagot/emagne/internal/handler/rest"
	"github.com/filagot/emagne/internal/module"
)

type Handler struct {
	mod module.Catagory
}

func Init(mod module.Catagory) rest.CategoryHandler {
	return &Handler{mod: mod}
}

func (h *Handler) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	desc := dto.ToNullString(req.Description)
	cat, err := h.mod.CreateCategory(c.Request.Context(), req.Name, desc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, cat)
}

func (h *Handler) ListCategories(c *gin.Context) {
	cats, err := h.mod.ListCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cats)
}

func (h *Handler) GetCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	cat, err := h.mod.GetCategory(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cat)
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	name := ""
	if req.Name != nil {
		name = *req.Name
	}
	desc := dto.ToNullString(req.Description)
	cat, err := h.mod.UpdateCategory(c.Request.Context(), id, name, desc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cat)
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.mod.DeleteCategory(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) CreateAttribute(c *gin.Context) {
	idStr := c.Param("id")
	categoryID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category id"})
		return
	}
	var req dto.CreateCategoryAttributeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isReq := dto.ToNullBool(req.IsRequired)
	attr, err := h.mod.CreateAttribute(c.Request.Context(), categoryID, req.Name, req.DataType, isReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, attr)
}

func (h *Handler) ListAttributes(c *gin.Context) {
	idStr := c.Param("id")
	categoryID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category id"})
		return
	}
	attrs, err := h.mod.ListAttributes(c.Request.Context(), categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, attrs)
}

func (h *Handler) GetAttribute(c *gin.Context) {
	idStr := c.Param("id")
	categoryID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category id"})
		return
	}
	name := c.Param("name")
	attr, err := h.mod.GetAttribute(c.Request.Context(), categoryID, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, attr)
}

func (h *Handler) DeleteAttribute(c *gin.Context) {
	attrIDStr := c.Param("attrID")
	id, err := uuid.Parse(attrIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid attribute id"})
		return
	}
	if err := h.mod.DeleteAttribute(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
