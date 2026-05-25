//go:build !debug

package assets

import "embed"

//go:embed *.html
//go:embed graphicarts
//go:embed javascripts
//go:embed stylesheets
//go:embed tester-env
var embedded embed.FS

func init() {
	FS.embedded = &embedded
}
