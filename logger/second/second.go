package second

import (
	"log/slog"
	"os"
)

type Second struct {
	name   string
	count  int
	logger *slog.Logger
}

func InitSecond(n string) *Second {
	fp, _ := os.Create("./log/log.txt")

	s := new(Second)
	s.name = n
	s.count = 0
	s.logger = slog.New(slog.NewJSONHandler(fp, nil))

	return s
}

func (s *Second) CountUp() {
	s.count++
	s.logger.Info("Countup!",
		slog.String("name", s.name),
		slog.Int("count", s.count))
}
