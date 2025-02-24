package writer

import "go.source.hueristiq.com/logger/levels"

type Writer interface {
	Write(data []byte, level levels.Level)
}
