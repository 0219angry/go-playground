package first

import (
	"log/slog"
	"os"
)

type First struct {
	name   string
	count  int
	logger *slog.Logger
}

func InitFirst(n string) *First {
	fp, _ := os.Create("./log/log.txt")

	f := new(First)
	f.name = n
	f.count = 0
	f.logger = slog.New(slog.NewJSONHandler(fp, nil))

	return f
}

func (f *First) CountUp() {
	f.count++
	f.logger.Info("Countup!",
		slog.String("name", f.name),
		slog.Int("count", f.count))
}
