package primitivetypes

// Rectangle representes a four-sided shape
type Rectangle struct {
	x      int
	y      int
	width  int
	height int
}

func (r Rectangle) findOverlap(r1 Rectangle) Rectangle {

	xo := findRangeOverlap(r.x, r.width, r1.x, r1.width)
	yo := findRangeOverlap(r.y, r.height, r1.y, r1.height)

	if xo.length == 0 || yo.length == 0 {
		return Rectangle{}
	}
	return Rectangle{xo.start, yo.start, xo.length, yo.length}
}

// RangeOverlap represents an overlapping range
type RangeOverlap struct {
	start  int
	length int
}

func findRangeOverlap(p1, l1, p2, l2 int) RangeOverlap {
	hsp := max(p1, p2)
	lsp := min(p1+l1, p2+l2)

	if hsp >= lsp {
		return RangeOverlap{0, 0}
	}
	ol := lsp - hsp
	return RangeOverlap{hsp, ol}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
