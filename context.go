package semv

import (
	"log"
)

// Ctx ...
type Ctx struct {
	WorkingDir string // Where to execute.
	Manifest   *Manifest
	Out, Err   *log.Logger // Required loggers.

}
