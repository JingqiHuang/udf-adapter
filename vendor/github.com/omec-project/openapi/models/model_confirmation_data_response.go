// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * AUSF API
 *
 * OpenAPI specification for AUSF
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ConfirmationDataResponse struct {
	AuthResult AuthResult `json:"authResult" yaml:"authResult" bson:"authResult"`
	Supi       string     `json:"supi,omitempty" yaml:"supi" bson:"supi"`
	Kseaf      string     `json:"kseaf,omitempty" yaml:"kseaf" bson:"kseaf"`
}
