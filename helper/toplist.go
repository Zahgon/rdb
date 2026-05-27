package helper

type Sized interface {
	GetSize() int
}

type topList struct {
	list     []Sized
	capacity int
}

func (tl *topList) add(x Sized) { _ = "STUB: not implemented"; return }

func newToplist(cap int) *topList { _ = "STUB: not implemented"; return nil }
