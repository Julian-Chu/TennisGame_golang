package test

import (
	"TennisGame/tennisgame"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LoveAll(t *testing.T) {
	g := &tennisgame.TennisGame{}
	assert.Equal(t, "Love All", g.Score(), "")
}
