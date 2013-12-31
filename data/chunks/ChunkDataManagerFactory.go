package data

import "appengine"

func GetDataManager(context *appengine.Context) (chunkDataManager ChunkDataManager) {
	var fcdm = NewAppEngineChunkDataManager(context)
	chunkDataManager = ChunkDataManager(fcdm)
	return
}
