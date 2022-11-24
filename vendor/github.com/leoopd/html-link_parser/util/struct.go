package htmllinkparser

type Href struct {
	Link string
	Text string
	Next *Href
}

type LinkedList struct {
	Head *Href
	Len  int
}

func (l *LinkedList) Insert(link, text string) {
	h := Href{}
	h.Link = link
	h.Text = text

	if l.Len == 0 {
		l.Head = &h
		l.Len++
		return
	}
	ptr := l.Head
	for i := 0; i < l.Len; i++ {
		if ptr.Next == nil {
			ptr.Next = &h
			l.Len++
			return
		}
		ptr = ptr.Next
	}
}
