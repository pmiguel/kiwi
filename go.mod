module github.com/pmiguel/kiwi

go 1.19

require github.com/pmiguel/kiwi/pkg/protocol v0.0.0

require github.com/google/uuid v1.3.0 // indirect

replace github.com/pmiguel/kiwi/pkg/protocol v0.0.0 => ./pkg/protocol
