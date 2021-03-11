package changelog

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"time"

	"github.com/josehbez/pm"
	"github.com/josehbez/pm/version"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

//Command ...
type Command struct {
}

// Run ...
func (command Command) Run(ctx *pm.Ctx) *cobra.Command {

	type changelogType struct {
		name, shorthand, value, usage string
	}
	changelogTypes := []changelogType{
		{"added", "a", "", "for new features"},
		{"changed", "c", "", "for changes in existing functionality"},
		{"deprecated", "d", "", "for soon-to-be removed features"},
		{"removed", "r", "", "for now removed features"},
		{"fixed", "f", "", "for any bug fixes"},
		{"security", "s", "", "in case of vulnerabilities"},
	}
	var subCLI = &cobra.Command{
		Use:   "changelog",
		Short: "Show or add changelogs",
		Example: `
pm changelog --added "A1"
pm changelog --fixed "F2"
pm changelog --fixed "F2" --added "A1"

 `,
		Run: func(kwargs *cobra.Command, args []string) {
			index := ""
			for _, i := range changelogTypes {
				if val, _ := kwargs.Flags().GetString(i.name); len(val) > 0 {
					index, _ = command.Add(ctx, i.name, val)
				}
			}

			s, sErr := command.Show(ctx, index)
			if sErr != nil {
				ctx.Err.Fatalln(sErr)
			}
			ctx.Out.Println(s)
		},
	}
	for _, i := range changelogTypes {
		subCLI.Flags().StringP(i.name, i.shorthand, i.value, i.usage)
	}
	return subCLI
}

// Show  ...
func (command Command) Show(ctx *pm.Ctx, indexShow string) (string, error) {
	type ChangelogType struct {
		Added      []string `yaml:"added"`
		Changed    []string `yaml:"changed"`
		Deprecated []string `yaml:"deprecated"`
		Removed    []string `yaml:"removed"`
		Fixed      []string `yaml:"fixed"`
		Security   []string `yaml:"security"`
	}
	changelogs := map[string]ChangelogType{}

	file, fileErr := ioutil.ReadFile(ctx.GetChangelogPath())
	if fileErr != nil {
		return "", fileErr
	}
	yaml.Unmarshal([]byte(file), &changelogs)
	msg := ""
	for index, changelog := range changelogs {

		var printcl = func() {

			msg += fmt.Sprintln(index, ":")
			var printscl = func(value reflect.Value, ttype string) {
				if item, ok := value.Interface().([]string); ok {
					if len(item) > 0 {
						msg += fmt.Sprintln(" ", ttype, ":")
						for _, l4v := range item {
							msg += fmt.Sprintln("    -", l4v)
						}
					}
				}

			}
			for i := 0; i < reflect.TypeOf(changelog).NumField(); i++ {
				fieldName := reflect.TypeOf(changelog).Field(i).Name
				printscl(
					reflect.ValueOf(changelog).FieldByName(fieldName),
					fieldName,
				)
			}

		}

		if len(indexShow) > 0 {
			if index == indexShow {
				printcl()
			}
		} else {
			printcl()
		}
	}
	return msg, nil
}

// Add ..
func (command Command) Add(ctx *pm.Ctx, ttype string, message string) (string, error) {

	index := fmt.Sprintf("%s (%s)",
		strings.Split(version.Command{}.GetVersion(ctx, "full"), "+")[0],
		time.Now().Format("2006-01-02"),
	)
	indexType := fmt.Sprintf("%s::%s", index, ttype)
	items := ctx.Changelog.GetStringSlice(indexType)

	items = append(items, message)

	ctx.Changelog.Set(indexType, items)
	return index, ctx.Changelog.WriteConfig()
}
