// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type QosResourceType string

// List of QosResourceType
const (
	QosResourceType_NON_GBR          QosResourceType = "NON_GBR"
	QosResourceType_NON_CRITICAL_GBR QosResourceType = "NON_CRITICAL_GBR"
	QosResourceType_CRITICAL_GBR     QosResourceType = "CRITICAL_GBR"
)
