package leetcode_1472

type node struct {
	previous *node
	current  string
	next     *node
}

type BrowserHistory struct {
	url *node
}

func Constructor(url string) BrowserHistory {
	return BrowserHistory{
		url: &node{
			previous: nil,
			current:  url,
			next:     nil,
		},
	}
}

func (bh *BrowserHistory) Visit(url string) {
	bh.url.next = &node{
		previous: bh.url,
		current:  url,
		next:     nil,
	}
	bh.url = bh.url.next
}

func (bh *BrowserHistory) Back(steps int) string {
	for i := 0; i < steps; i++ {
		if bh.url.previous == nil {
			break
		}
		bh.url = bh.url.previous
	}
	return bh.url.current
}

func (bh *BrowserHistory) Forward(steps int) string {
	for i := 0; i < steps; i++ {
		if bh.url.next == nil {
			break
		}
		bh.url = bh.url.next
	}
	return bh.url.current
}
