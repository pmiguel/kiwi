module github.com/pmiguel/kiwi

go 1.18

require (
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	github.com/pmiguel/kiwi/common v0.0.0
)

replace (
	github.com/pmiguel/kiwi/common v0.0.0 => ./common
)
