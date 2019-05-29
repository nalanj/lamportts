package lamportts

import (
	"bytes"

	"github.com/gofrs/uuid"
)

// Timestamp is a Lamport  timestamp
type Timestamp struct {

	// ReplicaID
	ReplicaID uuid.UUID

	// Counter is number of events that have occurred prior to this timestamp
	Counter Counter
}

// New returns a new lamport timestamp from scratch
func New() *Timestamp {
	return &Timestamp{
		ReplicaID: uuid.Must(uuid.NewV4()),
		Counter:   Counter{},
	}
}

// Next moves the timestamp forward one tick and returns a new timestamp
func (t *Timestamp) Next() *Timestamp {
	return &Timestamp{
		ReplicaID: t.ReplicaID,
		Counter:   t.Counter.Increment(),
	}
}

// Update compares the current timestamp to the comparison timestamp and
// returns the greater of the comparison or the current plus a tick
func (t *Timestamp) Update(compareTo *Timestamp) *Timestamp {
	if CompareCounters(t.Counter, compareTo.Counter) < 0 {
		return &Timestamp{
			ReplicaID: t.ReplicaID,
			Counter:   compareTo.Counter.Increment(),
		}
	}
	return t.Next()
}

// Compare returns > 0 if a > b, < 0 if a < b, or 0 if a == b
func Compare(a, b *Timestamp) int {
	counterCompare := CompareCounters(a.Counter, b.Counter)

	if counterCompare == 0 {
		return bytes.Compare(a.ReplicaID[:], b.ReplicaID[:])
	}
	return counterCompare
}
