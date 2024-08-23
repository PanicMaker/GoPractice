package main

type RateLimiter interface {
	TryAcquire() (bool, error)
}
