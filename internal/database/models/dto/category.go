package dto

import "database/sql"

type CreateCategoryRequest struct {
    Name string `json:"name" binding:"required"`
    Description *string `json:"description"`
}

type UpdateCategoryRequest struct {
    Name *string `json:"name"`
    Description *string `json:"description"`
}

type CreateCategoryAttributeRequest struct {
    Name string `json:"name" binding:"required"`
    DataType string `json:"data_type" binding:"required"`
    IsRequired *bool `json:"is_required"`
}

func ToNullString(v *string) sql.NullString {
    if v == nil { return sql.NullString{} }
    return sql.NullString{String: *v, Valid: true}
}

func ToNullBool(v *bool) sql.NullBool {
    if v == nil { return sql.NullBool{} }
    return sql.NullBool{Bool: *v, Valid: true}
}

