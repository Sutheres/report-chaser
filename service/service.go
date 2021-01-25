package service

import (
	"github.com/Sutheres/report-chaser/internal/sec"
)

type Service interface {
}

type service struct {
	BuildHash  string
	CommitHash string
	//
	sec sec.SEC
}

type option func(s *service)

func NewService(buildHash, commitHash string, opts ...option) Service {
	svc := &service{
		BuildHash:  buildHash,
		CommitHash: commitHash,
	}
	svc.WithOptions(opts...)
	return svc
}

func (s *service) WithOptions(opts ...option) {
	for _, opt := range opts {
		opt(s)
	}
}

func WithSEC(sec sec.SEC) option {
	return func(s *service) {
		s.sec = sec
	}
}
