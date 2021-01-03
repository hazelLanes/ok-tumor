package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/hazelLanes/ok_tumor/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
