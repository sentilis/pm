package changelog

import (
	"fmt"
	"strings"
	"time"

	"github.com/josehbez/pm"
	"github.com/josehbez/pm/version"
	"github.com/spf13/cobra"
)

//Command ...
type Command struct {
}

// Run ...
func (command Command) Run(ctx *pm.Ctx) *cobra.Command {
	var requiredAdd = func(kwargs *cobra.Command) {
		val, err := kwargs.Flags().GetString("add")
		if err != nil {
			ctx.Err.Fatal("flag -a required")
		}
		if len(val) == 0 {
			ctx.Err.Fatal("flag -a required")
		}
	}

	var getDateAdd = func(kwargs *cobra.Command) (d bool, a string) {
		d, _ = kwargs.Flags().GetBool("date")
		a, _ = kwargs.Flags().GetString("add")
		return
	}
	var subCLI = &cobra.Command{
		Use:   "changelog",
		Short: "",
		Long:  ``,
		Run: func(kwargs *cobra.Command, args []string) {
			if err := ctx.LoadChangelog(); err != nil {
				ctx.Err.Fatal(err)
			}
			if ok, _ := kwargs.Flags().GetBool("version"); ok {
				requiredAdd(kwargs)
				d, a := getDateAdd(kwargs)
				command.groupedBy(ctx, "version", d, a)
			} else if ok, _ := kwargs.Flags().GetBool("date"); ok {
				requiredAdd(kwargs)
				d, a := getDateAdd(kwargs)
				command.groupedBy(ctx, "date", d, a)
			}
		},
	}
	subCLI.Flags().StringP("add", "a", "", "Add description")
	subCLI.Flags().BoolP("version", "v", false, "Group by version")
	subCLI.Flags().BoolP("date", "d", false, "Group by date")
	//subCLI.Flags().StringP("key", "k", "", "Group by custom key")

	return subCLI
}

// groupedByVersion .. get current version project and set
//
// CLI
// pm changelog -v -a "This is my first changelog"
//
// FILE
// changelog:
//	version-PreRelease:
// 		version: version-PreRelease
//		item:
//			- This is my first changelog
//
// CLI
// pm changelog -v -d -a "This is my first changelog"
//
// FILE
// changelog:
//	version-PreRelease-Date:
// 		version: Version-PreRelease
// 		date: Date
//		items:
//			- This is my first changelog
func (command Command) groupedBy(ctx *pm.Ctx, ttype string, date bool, message string) error {

	index := time.Now().Format("2006-01-02")
	if ttype == "version" {
		d := index
		index = strings.Split(version.Command{}.GetVersion(ctx, "full"), "+")[0]
		if date {
			index = fmt.Sprintf("%s (%s)", index, d)
		}
	} else { // date

	}

	items := ctx.Changelog.GetStringSlice(index)

	items = append(items, message)

	ctx.Changelog.Set(index, items)
	return ctx.Changelog.WriteConfig()
}
