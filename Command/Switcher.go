package main

type Switcher struct {
	command Command
}

func (s *Switcher) press() {
	s.command.execute()
}
