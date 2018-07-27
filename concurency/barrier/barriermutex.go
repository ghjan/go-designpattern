package barrier

import "sync"

/*
GO 利用sync库实现Barrier
https://studygolang.com/articles/5377
 */
type Barrier struct {
	curCnt int
	maxCnt int
	cond   *sync.Cond
}

func NewBarrier(maxCnt int) *Barrier {
	mutex := new(sync.Mutex)
	cond := sync.NewCond(mutex)
	return &Barrier{curCnt: maxCnt, maxCnt: maxCnt, cond: cond}
}

func (barrier *Barrier) BarrierWait() {
	barrier.cond.L.Lock()
	if barrier.curCnt--; barrier.curCnt > 0 { //如果curCnt还不是0，还需要等待
		barrier.cond.Wait()
	} else { //如果curCnt是0，表示全部凑齐了，可以唤醒所有等待的goroutine了
		barrier.cond.Broadcast()
		barrier.curCnt = barrier.maxCnt
	}
	barrier.cond.L.Unlock()
}
