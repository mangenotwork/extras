package model

var (
	tCPHeartBeatCacheMap = make(map[string]struct{})
)

type TCPHeartBeatCache struct {

}

func (t *TCPHeartBeatCache) Add(token string) {
	tCPHeartBeatCacheMap[token] = struct{}{}
}

func (t *TCPHeartBeatCache) Has(token string) bool {
	_, ok := tCPHeartBeatCacheMap[token]
	return ok
}

func (t *TCPHeartBeatCache) Del(token string) {
	delete(tCPHeartBeatCacheMap, token)
}