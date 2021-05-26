package 单例Singleton

import "sync"

// 单例模式
// 线程安全
// 共享同一个变量
type SingleClient struct {
	Name  string
	Count int64
	mux   sync.Mutex
}

var (
	singleClient *SingleClient
	once         sync.Once
)

func init() {
	once.Do(func() {
		singleClient = &SingleClient{
			Name:  "single",
			Count: 0,
		}
	})
}

func GetInstance() *SingleClient {
	return singleClient
}

func (p *SingleClient) SetName(name string) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.Name = name
}

func (p *SingleClient) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.Name
}

func (p *SingleClient) IncrCount(i int64) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.Count += i
}

func (p *SingleClient) GetCount() int64 {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.Count
}
