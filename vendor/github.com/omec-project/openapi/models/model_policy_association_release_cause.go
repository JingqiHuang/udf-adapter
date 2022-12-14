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

type PolicyAssociationReleaseCause string

// List of PolicyAssociationReleaseCause
const (
	PolicyAssociationReleaseCause_UNSPECIFIED      PolicyAssociationReleaseCause = "UNSPECIFIED"
	PolicyAssociationReleaseCause_UE_SUBSCRIPTION  PolicyAssociationReleaseCause = "UE_SUBSCRIPTION"
	PolicyAssociationReleaseCause_INSUFFICIENT_RES PolicyAssociationReleaseCause = "INSUFFICIENT_RES"
)
