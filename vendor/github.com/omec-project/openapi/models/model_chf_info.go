// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ChfInfo struct {
	SupiRangeList []SupiRange     `json:"supiRangeList,omitempty" yaml:"supiRangeList" bson:"supiRangeList" mapstructure:"SupiRangeList"`
	GpsiRangeList []IdentityRange `json:"gpsiRangeList,omitempty" yaml:"gpsiRangeList" bson:"gpsiRangeList" mapstructure:"GpsiRangeList"`
	PlmnRangeList []PlmnRange     `json:"plmnRangeList,omitempty" yaml:"plmnRangeList" bson:"plmnRangeList" mapstructure:"PlmnRangeList"`
}
