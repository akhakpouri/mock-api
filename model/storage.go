package model

type IStorage interface {
	Load(string) []byte
	Save(string, []byte)
}
