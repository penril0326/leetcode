package numberofrecentcalls

type RecentCounter struct {
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	return 0
}
