package limit

import (
	"math"
	"sync"
	"time"
)

// Window menghitung permintaan ke satu endpoint selama 10 detik.
type Window struct {
	mu    sync.Mutex
	end   time.Time
	count int
}

// RetryAfter mengembalikan 0 jika permintaan boleh lewat, atau waktu tunggu dalam detik.
func (w *Window) RetryAfter() int {
	w.mu.Lock()
	defer w.mu.Unlock()

	now := time.Now()
	if !now.Before(w.end) {
		w.end = now.Add(10 * time.Second)
		w.count = 0
	}
	if w.count >= 3 {
		return int(math.Ceil(w.end.Sub(now).Seconds()))
	}
	w.count++
	return 0
}
