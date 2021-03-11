package changelog_test

import (
	"log"
	"os"
	"testing"

	"github.com/josehbez/pm"
	"github.com/josehbez/pm/changelog"
)

func TestChangelog(t *testing.T) {
	workingDir, _ := os.Getwd()
	ctx := pm.NewCtx(workingDir)
	ic := pm.InitCommand{}
	if err := ic.Initialized(ctx); err != nil {
		log.Fatalln(err)
	}
	cc := changelog.Command{}
	indexAdded, indexAddedErr := cc.Add(ctx, "added", "Run TEST")
	if indexAddedErr != nil {
		log.Fatalln(indexAddedErr)
	}

	if _, err := cc.Show(ctx, indexAdded); err != nil {
		log.Fatalln(err)
	}

}
