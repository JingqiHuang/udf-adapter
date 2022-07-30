// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Nudm_EE
 *
 * Nudm Event Exposure Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type CreatedEeSubscription struct {
	EeSubscription *EeSubscription    `json:"eeSubscription" yaml:"eeSubscription" bson:"eeSubscription" mapstructure:"EeSubscription"`
	NumberOfUes    int32              `json:"numberOfUes,omitempty" yaml:"numberOfUes" bson:"numberOfUes" mapstructure:"NumberOfUes"`
	EventReports   []MonitoringReport `json:"eventReports,omitempty" yaml:"eventReports" bson:"eventReports" mapstructure:"EventReports"`
}