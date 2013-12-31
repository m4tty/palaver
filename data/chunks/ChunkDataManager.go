package data

type ChunkDataManager interface {
	GetChunks() (results []*Chunk, err error)
	GetChunkById(id string) (result Chunk, err error)
	SaveChunk(chunk *Chunk) (key string, err error)
	DeleteChunk(id string) (err error)
}
