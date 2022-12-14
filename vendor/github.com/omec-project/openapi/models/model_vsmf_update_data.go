// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type VsmfUpdateData struct {
	RequestIndication         RequestIndication             `json:"requestIndication"`
	SessionAmbr               *Ambr                         `json:"sessionAmbr,omitempty"`
	QosFlowsAddModRequestList []QosFlowAddModifyRequestItem `json:"qosFlowsAddModRequestList,omitempty"`
	QosFlowsRelRequestList    []QosFlowReleaseRequestItem   `json:"qosFlowsRelRequestList,omitempty"`
	EpsBearerInfo             []EpsBearerInfo               `json:"epsBearerInfo,omitempty"`
	AssignEbiList             []int32                       `json:"assignEbiList,omitempty"`
	RevokeEbiList             []int32                       `json:"revokeEbiList,omitempty"`
	ModifiedEbiList           []EbiArpMapping               `json:"modifiedEbiList,omitempty"`
	Pti                       int32                         `json:"pti,omitempty"`
	N1SmInfoToUe              *RefToBinaryData              `json:"n1SmInfoToUe,omitempty"`
	AlwaysOnGranted           bool                          `json:"alwaysOnGranted,omitempty"`
	Cause                     Cause                         `json:"cause,omitempty"`
	N1smCause                 string                        `json:"n1smCause,omitempty"`
	BackOffTimer              int32                         `json:"backOffTimer,omitempty"`
}
