package rope

import (
	"io"
)

// A node representing a contiguous string
type leaf string

func (l leaf) depth() depthT { return 0 }
func (l leaf) length() int64 { return int64(len(l)) }

func (l leaf) WriteTo(w io.Writer) (n int64, err error) {
	n1, err := io.WriteString(w, string(l))
	return int64(n1), err
}

// Precondition: start < end
func (l leaf) slice(start, end int64) node {
	if start < 0 {
		start = 0
	}
	if end > int64(len(l)) {
		end = int64(len(l))
	}
	// The precondition may have been destroyed by above fixes.
	if start >= end {
		// Don't hold a 0-length substring, let the GC have it.
		return emptyNode
	}
	return l[start:end]
}

func (l leaf) dropPrefix(start int64) node {
	switch {
	case start >= int64(len(l)):
		return emptyNode
	case start <= 0:
		return l
	default: // 0 < start < len(l)
		return l[start:]
	}
}

func (l leaf) dropPostfix(end int64) node {
	switch {
	case end >= int64(len(l)):
		return l
	case end <= 0:
		return emptyNode
	default:
		return l[:end]
	}
}
