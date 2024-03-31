package merkle

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkDifference(b *testing.B) {
	s1 := []uint64{11, 15, 84, 88888888, 999999999999999}
	s2 := []uint64{11, 15, 1333, 7777777777, 999999999999999}
	for n := 0; n < b.N; n++ {
		sliceDifferences(s1, s2)
	}
}

func TestPopFromIndexQueue(t *testing.T) {
	s := []uint64{11, 15, 84, 88888888, 999999999999999}
	t.Logf("raw: %+v", s)
	newS := popFromIndexQueue(s)
	t.Logf("popFromIndexQueue: %+v", newS)

}

func TestSliceDifferences(t *testing.T) {
	sclice1 := []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Logf("raw sclice1: %+v", sclice1)
	sclice2 := []uint64{1, 0, 3, 2, 5, 4, 7, 6, 9, 8}
	t.Logf("raw sclice2: %+v", sclice2)
	sliceDiff := sliceDifferences(sclice1, sclice2)
	require.Empty(t, len(sliceDiff), 0)
	t.Logf("sliceDifferences : %+v", sliceDiff)

}
