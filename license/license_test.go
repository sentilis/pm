package license_test

import (
	"log"
	"os"
	"testing"

	"github.com/josehbez/pm"
	"github.com/josehbez/pm/license"
)

func TestLicense(t *testing.T) {
	workingDir, _ := os.Getwd()
	ctx := pm.NewCtx(workingDir)
	l := license.Command{}
	licenses, licensesErr := l.List(ctx)
	if licensesErr != nil {
		log.Fatalln(licensesErr)
	}
	if len(licenses) == 0 {
		log.Fatalln("Not found licenses")
	}
	licenseMIT, licenseMITErr := l.Fetch(ctx, "MIT")
	if licenseMITErr != nil {
		log.Fatalln(licenseMIT)
	}
	if err := l.Save(ctx, licenseMIT); err != nil {
		log.Fatalln(err)
	}
	if _, err := l.Show(ctx); err != nil {
		log.Fatalln(err)
	}
	if err := l.Remove(ctx); err != nil {
		log.Fatalln(err)
	}
}
