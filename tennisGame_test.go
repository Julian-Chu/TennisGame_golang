package TennisGame

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TennisGameTestSuite struct {
	game *TennisGame
	suite.Suite
}

func (ts TennisGameTestSuite) SetupTest() {
	ts.game = &TennisGame{}
}

func TestInitTestSuite(t *testing.T) {
	suite.Run(t, new(TennisGameTestSuite))
}

//var game *TennisGameTestSuite

//func Test_LoveAll(t *testing.T) {
//	assert.Equal(t, "Love All", game.Score())
//}

func (suite *TennisGameTestSuite) Test_Love_All() {
	assert.Equal(suite.T(), "Love All", suite.game.Score())
}
