package pkg

import (
	"database/sql"
	"time"
)

// NullStringToPointer Helper function to convert sql.NullString to *string
func NullStringToPointer(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

// NullTimeToPointer Helper function to convert sql.NullTime to *time.Time
func NullTimeToPointer(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}

// NullTimeToTime Helper function to convert sql.NullTime to time.Time with a fallback
func NullTimeToTime(nt sql.NullTime) time.Time {
	if nt.Valid {
		return nt.Time
	}
	return time.Time{} // Or use a zero value or specific fallback
}
