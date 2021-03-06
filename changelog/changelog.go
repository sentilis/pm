package changelog

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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
		Short: "Show or add changelogs",
		Example: `
pm changelog -v -a "Grouped by version"
1.0.1-beta.1: 
- Grouped by version

pm changelog -vd -a "Grouped by version and date"
1.0.1-beta.1 (2021-03-10): 
- Grouped by version and date

pm changelog -d -a "Grouped by date"
2021-03-10: 
 - First changelog
 `,
		Run: func(kwargs *cobra.Command, args []string) {

			show := true
			if ok, _ := kwargs.Flags().GetBool("version"); ok {
				requiredAdd(kwargs)
				d, a := getDateAdd(kwargs)
				command.add(ctx, "version", d, a)
				show = false
			} else if ok, _ := kwargs.Flags().GetBool("date"); ok {
				requiredAdd(kwargs)
				d, a := getDateAdd(kwargs)
				command.add(ctx, "date", d, a)
				show = false
			}
			if show {
				command.show(ctx)
			}

		},
	}

	subCLI.Flags().StringP("add", "a", "", "add description")
	subCLI.Flags().BoolP("version", "v", false, "grouped by version")
	subCLI.Flags().BoolP("date", "d", false, "grouped by date")

	return subCLI
}

func (command Command) show(ctx *pm.Ctx) {
	cmd := exec.Command("less")

	file, err := ioutil.ReadFile(ctx.GetChangelogPath())
	if err != nil {
		ctx.Err.Fatal(err)
	}

	cmd.Stdin = strings.NewReader(string(file))
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		ctx.Err.Fatal(err)
	}
}

// add .. Set log grouped by date|version|version (date)
func (command Command) add(ctx *pm.Ctx, ttype string, date bool, message string) error {

	index := time.Now().Format("2006-01-02")
	if ttype == "version" {
		d := index
		index = strings.Split(version.Command{}.GetVersion(ctx, "full"), "+")[0]
		if date {
			index = fmt.Sprintf("%s (%s)", index, d)
		}
	}

	items := ctx.Changelog.GetStringSlice(index)

	items = append(items, message)

	ctx.Changelog.Set(index, items)
	return ctx.Changelog.WriteConfig()
}
