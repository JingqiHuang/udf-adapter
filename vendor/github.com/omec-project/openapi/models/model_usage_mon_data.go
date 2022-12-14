// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains remain allowed usage data for a subscriber.
type UsageMonData struct {
	LimitId      string                       `json:"limitId" bson:"limitId"`
	Scopes       map[string]UsageMonDataScope `json:"scopes,omitempty" bson:"scopes"`
	UmLevel      UsageMonLevel                `json:"umLevel,omitempty" bson:"umLevel"`
	AllowedUsage *UsageThreshold              `json:"allowedUsage,omitempty" bson:"allowedUsage"`
	ResetTime    *TimePeriod                  `json:"resetTime,omitempty" bson:"resetTime"`
}
