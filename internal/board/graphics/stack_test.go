package graphics

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCellStack_Render_lines(t *testing.T) {
	a := require.New(t)
	cs := CellStack{}
	cs.add(LinesCell(Lines{
		t: true,
	}))
	cs.add(LinesCell(Lines{
		b: true,
	}))
	a.Equal("╷", cs.Render())
}

func TestCellStack_Render_charHasPriority(t *testing.T) {
	a := require.New(t)
	cs := CellStack{}
	cs.add(LinesCell(Lines{
		t: true,
	}))
	cs.PlaceChar('a')
	a.Equal("a", cs.Render())
}

func TestCellStack_Render_arrowHasPriority(t *testing.T) {
	a := require.New(t)
	cs := CellStack{}
	cs.PlaceArrow(false)
	cs.add(LinesCell(Lines{
		t: true,
	}))
	a.Equal("▷", cs.Render())
}

func TestCellStack_Tag(t *testing.T) {
	a := require.New(t)
	cs := CellStack{}

	a.Equal(false, cs.Is("key", "foo"))
	cs.Tag("key", "bar")
	a.Equal(false, cs.Is("key", "foo"))
	a.Equal(true, cs.Is("key", "bar"))
	cs.Tag("key", "foo")
	cs.Tag("otherKey", "bar")
	a.Equal(true, cs.Is("key", "foo"))
	a.Equal(true, cs.Is("otherKey", "bar"))
}