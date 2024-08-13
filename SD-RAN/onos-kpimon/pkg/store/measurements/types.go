// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package measurements

// MeasurementItem measurement item
type MeasurementItem struct {
	MeasurementRecords []MeasurementRecord
}

// MeasurementRecord measurement record
type MeasurementRecord struct {
	Timestamp        uint64
	MeasurementName  string
	MeasurementValue interface{}
}

// CellIdentity is the ID for each cell
type CellIdentity struct {
	CellID string
}

// Key is the key of monitoring result metric store
type Key struct {
	NodeID       string
	CellIdentity CellIdentity
}

// Entry measurement store entry
type Entry struct {
	Key   Key
	Value interface{}
}

// MeasurementEvent a measurement event
type MeasurementEvent int

const (
	// None none cell event
	None MeasurementEvent = iota
	// Created created measurement event
	Created
	// Updated updated measurement event
	Updated
	// Deleted deleted measurement event
	Deleted
)

func (e MeasurementEvent) String() string {
	return [...]string{"None", "Created", "Updated", "Deleted"}[e]
}
