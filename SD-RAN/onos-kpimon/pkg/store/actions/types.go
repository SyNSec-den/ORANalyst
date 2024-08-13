// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package actions

// CellIdentity is the ID for each cell
type CellIdentity struct {
	CellID string
}

// SubscriptionID subID is used for creating action definition
type SubscriptionID struct {
	SubID int64
}

// Key is the key of cells store entries
type Key struct {
	SubscriptionID SubscriptionID
}

// Entry store entry
type Entry struct {
	Key   Key
	Value interface{}
}
