package suite

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func Run(t *testing.T, s suite.TestingSuite) {
	suite.Run(t, s)
}
