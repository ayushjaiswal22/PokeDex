package pokecache

import(
    "time"
//    "fmt"
    "sync"
)

type Cache struct{
    mu sync.RWMutex
    cacheMap map[string]CacheVal
}

type CacheVal struct{
    val []byte
    createdAt time.Time
}

func CreateCache(interval time.Duration) Cache{
    c := Cache{ 
        cacheMap: make(map[string]CacheVal), 
    }
    go c.PollCache(interval)
    return c
}

func (c *Cache)AddCacheVal(key string, value []byte) {
    c.mu.Lock()
    c.cacheMap[key] = CacheVal{val:value, createdAt:time.Now().UTC()}
    c.mu.Unlock()
}

func (c *Cache)GetCacheVal(key string) ([]byte, bool) {
    c.mu.RLock()
    c.mu.RUnlock()
    cacheVal, ok := c.cacheMap[key]
    return cacheVal.val, ok
}

func (c *Cache)UpdateCacheVal(key string){
    c.mu.Lock()
    defer c.mu.Unlock()
    cv := c.cacheMap[key]
    cv.createdAt = time.Now().UTC()
    c.cacheMap[key] = cv
}

func (c *Cache)PollCache(interval time.Duration) {
    ticker := time.NewTicker(interval)
    for range ticker.C {
        c.reap(interval)
    }
}
func (c *Cache)reap(interval time.Duration) {
    timeLimit := time.Now().UTC().Add(-interval)
    for key, value := range c.cacheMap {
        if value.createdAt.Before(timeLimit) {
            //fmt.Printf("\ndeleted %s\n", key)
            delete(c.cacheMap, key)
        }
    }
}
