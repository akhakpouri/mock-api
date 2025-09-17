package model

type Pod struct {
	content []byte
	storage IStorage
}

func NewPod(s IStorage) *Pod {
	return &Pod{
		storage: s,
	}
}

func (p *Pod) Save(name string) {
	p.storage.Save(name, p.content)
}

func (p *Pod) Load(name string) {
	p.content = p.storage.Load(name)
}
