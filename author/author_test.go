package author_test

import (
	"log"
	"os"
	"testing"

	"github.com/josehbez/pm"
	"github.com/josehbez/pm/author"
)

func TestAuthor(t *testing.T) {
	workingDir, _ := os.Getwd()
	ctx := pm.NewCtx(workingDir)

	ic := pm.InitCommand{}
	if err := ic.Initialized(ctx); err != nil {
		log.Fatalln(err)
	}
	if err := ctx.Author.Load(); err != nil {
		log.Fatalln(err)
	}

	for _, val := range []string{"author", "maitainer"} {
		if _, err := author.Add(ctx, val, "jose hbez", []string{"https://github.com/josehbez"}); err != nil {
			log.Fatalln(err)
		}
		if _, err := author.Show(ctx, val); err != nil {
			log.Fatalln(err)
		}
	}

}
