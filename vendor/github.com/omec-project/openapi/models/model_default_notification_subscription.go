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

type DefaultNotificationSubscription struct {
	NotificationType   NotificationType   `json:"notificationType" yaml:"notificationType" bson:"notificationType" mapstructure:"NotificationType"`
	CallbackUri        string             `json:"callbackUri" yaml:"callbackUri" bson:"callbackUri" mapstructure:"CallbackUri"`
	N1MessageClass     N1MessageClass     `json:"n1MessageClass,omitempty" yaml:"n1MessageClass" bson:"n1MessageClass" mapstructure:"N1MessageClass"`
	N2InformationClass N2InformationClass `json:"n2InformationClass,omitempty" yaml:"n2InformationClass" bson:"n2InformationClass" mapstructure:"N2InformationClass"`
}
