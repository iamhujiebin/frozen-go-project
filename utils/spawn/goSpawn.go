package spawn

// 使用 buffered channel 的协程并发数控制工具
type GoSpawn struct {
	queue chan struct{}
}

func NewGoSpawn(max int) *GoSpawn {
	return &GoSpawn{
		queue: make(chan struct{}, max),
	}
}

func (g *GoSpawn) Go(f func()) {
	g.queue <- struct{}{}
	go func() {
		f()
		<-g.queue
	}()
}
