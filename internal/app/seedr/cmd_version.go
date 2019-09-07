package seedr

import (
	"github.com/NoUseFreak/go-vembed"
)

func init() {
	rootCmd.Version = vembed.Version.String()
}
