// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Namf_Communication
 *
 * AMF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type NonUeN2InfoSubscriptionCreatedData struct {
	N2NotifySubscriptionId string `json:"n2NotifySubscriptionId"`
	SupportedFeatures      string `json:"supportedFeatures,omitempty"`
}
