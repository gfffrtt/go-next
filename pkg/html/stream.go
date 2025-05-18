package html

import (
	"sync"
)

type Stream struct {
	Channel   chan string
	WaitGroup sync.WaitGroup
}

func NewStream() *Stream {
	return &Stream{Channel: make(chan string)}
}

func (s *Stream) Write(html string) {
	s.Channel <- html
	s.WaitGroup.Done()
}

func (s *Stream) Add() {
	s.WaitGroup.Add(1)
}

func (s *Stream) Wait() {
	s.WaitGroup.Wait()
	close(s.Channel)
}
