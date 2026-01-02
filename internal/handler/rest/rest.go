package rest

import (
	"github.com/gin-gonic/gin"
)

// AuthHandler defines the interface for authentication HTTP handlers
type AuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetProfile(c *gin.Context)
	UpdateProfile(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type CategoryHandler interface {
	CreateCategory(c *gin.Context)
	ListCategories(c *gin.Context)
	GetCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)

	CreateAttribute(c *gin.Context)
	ListAttributes(c *gin.Context)
	GetAttribute(c *gin.Context)
	DeleteAttribute(c *gin.Context)
}

type TransactionHandler interface {
	CreateTransaction(c *gin.Context)
	GetTransaction(c *gin.Context)
	ListTransactionsByCategory(c *gin.Context)
	UpdateTransactionStatus(c *gin.Context)

	AddTransactionAttribute(c *gin.Context)
	ListTransactionAttributes(c *gin.Context)
	GetTransactionAttribute(c *gin.Context)
	DeleteTransactionAttributes(c *gin.Context)
}
