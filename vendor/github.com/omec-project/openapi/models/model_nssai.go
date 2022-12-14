// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type Nssai struct {
	SupportedFeatures   string   `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures" mapstructure:"SupportedFeatures"`
	DefaultSingleNssais []Snssai `json:"defaultSingleNssais" yaml:"defaultSingleNssais" bson:"defaultSingleNssais" mapstructure:"DefaultSingleNssais"`
	SingleNssais        []Snssai `json:"singleNssais,omitempty" yaml:"singleNssais" bson:"singleNssais" mapstructure:"SingleNssais"`
}
