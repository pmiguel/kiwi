module github.com/pmiguel/kiwi

go 1.18

require (
	github.com/pmiguel/kiwi/pkg/protocol v0.0.0
)

replace (
	github.com/pmiguel/kiwi/pkg/protocol v0.0.0 => ./pkg/protocol
)
