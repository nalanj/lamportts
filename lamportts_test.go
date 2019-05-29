package lamportts

import (
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAfter(t *testing.T) {
	t.Parallel()

	rep0Count0 := &Timestamp{
		ReplicaID: uuid.UUID{},
		Counter:   Counter{},
	}

	rep0Count1 := &Timestamp{
		ReplicaID: uuid.UUID{},
		Counter:   Counter{0x01},
	}

	repNon0Count1 := &Timestamp{
		ReplicaID: uuid.Must(uuid.NewV4()),
		Counter:   Counter{0x01},
	}

	testCases := []struct {
		desc      string
		ts        *Timestamp
		compareTo *Timestamp
		expect    int
	}{
		{
			desc:      "counter before",
			ts:        rep0Count0,
			compareTo: rep0Count1,
			expect:    -1,
		},
		{
			desc:      "counter equal, replica before",
			ts:        rep0Count1,
			compareTo: repNon0Count1,
			expect:    -1,
		},
		{
			desc:      "counter equal, replica after",
			ts:        repNon0Count1,
			compareTo: rep0Count1,
			expect:    1,
		},
		{
			desc:      "counter equal, replica equal",
			ts:        rep0Count1,
			compareTo: rep0Count1,
			expect:    0,
		},
		{
			desc:      "counter after",
			ts:        rep0Count1,
			compareTo: rep0Count0,
			expect:    1,
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.expect, Compare(tC.ts, tC.compareTo))
		})
	}
}
