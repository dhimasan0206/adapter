/*
 * Integration Adapter
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

type UpdateStockLevel500Response struct {

	Error UpdateStockLevel500ResponseError `json:"error,omitempty"`

	Id string `json:"id,omitempty"`

	Message string `json:"message,omitempty"`

	Status string `json:"status,omitempty"`

	Store UpdateStockLevelRequestStore `json:"store,omitempty"`

	Success bool `json:"success,omitempty"`

	Timestamp float32 `json:"timestamp,omitempty"`
}
