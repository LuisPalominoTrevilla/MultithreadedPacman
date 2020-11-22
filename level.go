package main

import "io/ioutil"

// Level represents a level with all of its contents
type Level struct {
	layout         string
	hTiles, vTiles int
}

func (l *Level) parseLevel(file string) error {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	l.layout = string(dat)
	return nil
}

// InitLevel given a valid level file
func InitLevel(levelFile string) (*Level, error) {
	l := Level{}
	err := l.parseLevel(levelFile)
	return &l, err
}
