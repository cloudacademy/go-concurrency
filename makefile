all: goroutines waitgroups channels mutex

channels: buffers close range select timeout pingpong

# ======================
# goroutines

goroutines:
	go run ./cmd/goroutines/goroutines.go

# ======================
# waitgroups

waitgroups:
	go run ./cmd/waitgroups/waitgroups.go

httpfetcher:
	go run ./cmd/waitgroups/http/httpfetcher.go

# ======================
# channels

buffers:
	go run ./cmd/channels/buffers/buffers.go

close:
	go run ./cmd/channels/close/close.go

range:
	go run ./cmd/channels/range/range.go

select:
	go run ./cmd/channels/select/channelselect.go

timeout:
	go run ./cmd/channels/timeout/timeouts.go

pingpong:
	go run ./cmd/channels/pattern/pingpong.go

# ======================
# mutex:

mutex:
	go run ./cmd/mutex/counter.go

