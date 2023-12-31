package support

import (
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/brunoga/unitybridge"
	"github.com/brunoga/unitybridge/support/logger"
	"github.com/brunoga/unitybridge/support/token"
	"github.com/brunoga/unitybridge/unity/key"
	"github.com/brunoga/unitybridge/unity/result"
)

// ResultListener is a helper class to listen for event results from the
// Unity Bridge. It allows callers to wait for new results, to get the
// the last result obtained and to register a callback to be called when
// a new result is available. It is thread safe (and lock free).
type ResultListener struct {
	ub unitybridge.UnityBridge
	l  *logger.Logger
	k  *key.Key
	cb result.Callback

	t token.Token

	m       sync.Mutex
	r       *result.Result
	c       chan struct{}
	started bool
}

// NewResultListener creates a new ResultListener instance.
func NewResultListener(ub unitybridge.UnityBridge, l *logger.Logger,
	k *key.Key, cb result.Callback) *ResultListener {
	if l == nil {
		l = logger.New(slog.LevelError)
	}

	l = l.WithGroup("result_listener").With(
		slog.String("key", k.String()))

	return &ResultListener{
		ub: ub,
		l:  l,
		k:  k,
		cb: cb,
	}
}

// Start starts the listener. If cb is non nil, it will be called when a new
// result is available.
func (ls *ResultListener) Start() error {
	ls.m.Lock()
	defer ls.m.Unlock()

	if ls.started {
		return fmt.Errorf("listener already started")
	}

	ls.c = make(chan struct{})
	ls.r = nil

	var err error

	ls.t, err = ls.ub.AddKeyListener(ls.k, func(r *result.Result) {
		ls.m.Lock()

		ls.r = r
		ls.notifyWaitersLocked()

		ls.m.Unlock()

		if ls.cb != nil && r.Succeeded() {
			go ls.cb(r)
		}
	}, true)

	ls.started = true

	return err
}

// WaitForNewResult blocks until a new result is available, a timeout happens
// or the listener is stopped. IF result is nil, no result was available (for
// example, if the listener is closed). If result is non nil, Callers should
// inspect the result error code and description to check if the result is
// valid.
func (ls *ResultListener) WaitForNewResult(timeout time.Duration) *result.Result {
	ls.m.Lock()
	c := ls.c
	ls.m.Unlock()

	select {
	case <-c:
		return ls.Result()
	case <-time.After(timeout):
		return nil
	}
}

// WaitForAnyResult returns any existing result imemdiatelly or blocks until a
// result is available, a timeout happens or the listener is stopped. IF result
// is nil, no result was available (for example, if the listener is closed). If
// result is non nil, Callers should inspect the result error code and
// description to check if the result is valid.
func (ls *ResultListener) WaitForAnyResult(timeout time.Duration) *result.Result {
	// Make sure we get a correct snapshot of the current channel and result
	// state by obtainignthem inside a lock. This guarantee that we either
	// have a result or that, if we do not, we are going to be listen on a
	// channel that is guaranteed to b ethe one existing when the value
	// was nil so either it is closed now and we do have a non-nil value or
	// it will be closed after we start waiting on it (and we will get a result
	// or a timeout.
	ls.m.Lock()
	if ls.r != nil {
		ls.m.Unlock()
		return ls.r
	}

	c := ls.c

	ls.m.Unlock()

	select {
	case <-c:
		return ls.Result()
	case <-time.After(timeout):
		return nil
	}
}

// Result returns the current result.
func (ls *ResultListener) Result() *result.Result {
	ls.m.Lock()
	defer ls.m.Unlock()

	return ls.r
}

// Stop stops the listener.
func (ls *ResultListener) Stop() error {
	ls.m.Lock()
	defer ls.m.Unlock()

	if !ls.started {
		return fmt.Errorf("listener not started")
	}

	err := ls.ub.RemoveKeyListener(ls.k, ls.t)
	if err != nil {
		return err
	}

	ls.r = nil
	ls.started = false

	ls.notifyWaitersLocked()

	return nil
}

// notifyWaitersLocked closes the current channel and creates a new one.
// The channel mutex must be locked when this is called.
func (ls *ResultListener) notifyWaitersLocked() {
	close(ls.c)
	ls.c = make(chan struct{})
}
