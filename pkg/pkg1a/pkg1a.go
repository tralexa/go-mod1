package pkg1a

import (
	"github.com/tralexa/go-mod1/pkg/pkg1b"
	"github.com/tralexa/go-mod1/pkg/pkg1c"
	"github.com/tralexa/go-mod1/pkg/pkg1d"
)

func Do() {
	pkg1b.Do()
	pkg1c.Do()
	pkg1d.Do()
}