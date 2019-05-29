package lamportts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrement(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc   string
		start  Counter
		expect Counter
	}{
		{
			desc:   "increments without increasing length",
			start:  Counter{0x3F},
			expect: Counter{0x40},
		},
		{
			desc:   "increments increasing length",
			start:  Counter{0xFF, 0x7F},
			expect: Counter{0x81, 0x80, 0x00},
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.expect, tC.start.Increment())
		})
	}
}

func TestCompare(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc   string
		a      Counter
		b      Counter
		expect int
	}{
		{
			desc:   "a less than b",
			a:      Counter{0xFF, 0x7A},
			b:      Counter{0xFF, 0x7B},
			expect: -1,
		},
		{
			desc:   "a equal b",
			a:      Counter{0xFF, 0x7A},
			b:      Counter{0xFF, 0x7A},
			expect: 0,
		},
		{
			desc:   "a greater than b",
			a:      Counter{0xFF, 0x7B},
			b:      Counter{0xFF, 0x7A},
			expect: 1,
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.expect, Compare(tC.a, tC.b))
		})
	}
}
