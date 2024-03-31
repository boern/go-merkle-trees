package merkle

import "testing"

func TestSiblingIndecies(t *testing.T) {
	idxes := []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8}
	t.Logf("raw idxes: %+v", idxes)
	siblingIdxes := siblingIndecies(idxes)
	t.Logf("sibling idxes: %+v", siblingIdxes)

}

func TestParentIndecies(t *testing.T) {
	idxes := []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Logf("raw idxes: %+v", idxes)
	parentIdxes1 := parentIndecies(idxes)
	t.Logf("parent idxes1: %+v", parentIdxes1)
	// idxes2 := []uint64{1, 0, 3, 2, 5, 4, 7, 6, 9, 8}
	siblingIdxes := siblingIndecies(idxes)
	t.Logf("sibling idxes: %+v", siblingIdxes)
	parentIdxes2 := parentIndecies(siblingIdxes)
	t.Logf("parent idxes2: %+v", parentIdxes2)
}
