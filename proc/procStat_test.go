package proc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRunningProcesses(t *testing.T) {
	lines := processLines()
	assert.True(t, len(lines) > 0, "empty ps list")
}

func TestFilterPs(t *testing.T) {
	lines := parsePsStats([]string{"     1     1  1024", "     2     2  1024"})
	expected := []Process{{1, 1, 1024}, {2, 2, 1024}}
	assert.Equal(t, expected, lines, "wrong count")
}

func TestCalcStat(t *testing.T) {
	main := Process{1, 0, 1}

	child1 := Process{2, main.PID, 1}
	child2 := Process{3, main.PID, 1}
	grandChild := Process{4, child1.PID, 1}
	grandGrandChild := Process{5, grandChild.PID, 1}

	unrelatedProcess := Process{6, 0, 1}

	stats := []Process{
		main,
		child1,
		child2,
		grandChild,
		grandGrandChild,
		unrelatedProcess,
	}

	pidRss, descendantsRss := CalculateRss(stats, 1)

	assert.Equal(t, 1, pidRss, "wrong RSS for pid")
	assert.Equal(t, 4, descendantsRss, "wrong RSS for descendants")
}
