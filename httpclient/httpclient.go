package httpclient

import "sync"

type httpClient struct {
	lock     *sync.Mutex
	withLock bool
}
