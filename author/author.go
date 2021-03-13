package author

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/josehbez/pm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type commandFlag struct {
	Usage string
}

type commandArgs struct {
	Use     string
	Short   string
	Example string
	FlagAdd commandFlag
}

//Command ...
type Command struct {
}

// Run ...
func (c Command) Run(ctx *pm.Ctx) *cobra.Command {
	ca := commandArgs{
		Use:   "author",
		Short: "Show & add authors",
		Example: `
pm author --add "Jose Hbez" https://github.com/josehbez
pm author --add "Jose Hbez" email@email.com email2@email.com
`,
		FlagAdd: commandFlag{Usage: "add author"},
	}
	return run(ctx, ca)
}

//CommandMaintainer ...
type CommandMaintainer struct {
}

func (c CommandMaintainer) Run(ctx *pm.Ctx) *cobra.Command {
	ca := commandArgs{
		Use:   "maintainer",
		Short: "Show & add maintainers",
		Example: `
pm maintainer --add "Jose Hbez" https://github.com/josehbez
`,
		FlagAdd: commandFlag{Usage: "add maintainer"},
	}
	return run(ctx, ca)
}

func run(ctx *pm.Ctx, ca commandArgs) *cobra.Command {
	var checkLoad = func(e error) {
		if e != nil {
			ctx.Err.Fatalln(e)
		}
	}

	var subCommand = &cobra.Command{
		Use:     ca.Use,
		Short:   ca.Short,
		Example: ca.Example,
		Run: func(kwargs *cobra.Command, args []string) {

			if kwargs.Name() == "maintainer" {
				checkLoad(ctx.Maintainer.Load())
			} else { // Default authors
				checkLoad(ctx.Author.Load())
			}

			show := true

			if val, _ := kwargs.Flags().GetString("add"); len(val) > 0 {
				a, aErr := Add(ctx, kwargs.Name(), val, args)
				if aErr != nil {
					ctx.Err.Fatalln(aErr)
				}
				ctx.Out.Println(a)
				show = false
			}
			if show {
				s, sErr := Show(ctx, kwargs.Name())
				if sErr != nil {
					ctx.Err.Fatalln(sErr)
				}
				ctx.Out.Println(s)
			}

		},
	}

	subCommand.Flags().StringP("add", "a", "", ca.FlagAdd.Usage)

	return subCommand
}

func Show(ctx *pm.Ctx, commandName string) (string, error) {

	authors := map[string][]string{}
	manifestPath := ""

	if commandName == "maintainer" {
		manifestPath = ctx.Maintainer.GetPath()
	} else {
		manifestPath = ctx.Author.GetPath()
	}

	file, fileErr := ioutil.ReadFile(manifestPath)
	if fileErr != nil {
		return "", fileErr
	}

	yaml.Unmarshal([]byte(file), &authors)

	val := ""
	for index, items := range authors {
		if fmt.Sprintf("%ss", commandName) == index {
			for _, item := range items {
				val += fmt.Sprintln(" - " + item)
			}
		}
	}
	return val, nil
}

// Add ...
func Add(ctx *pm.Ctx, commandName, name string, args []string) (string, error) {

	index := fmt.Sprintf("%ss", commandName)
	val := name
	emailURL := strings.Join(args, ",")
	if len(emailURL) > 0 {
		val = fmt.Sprintf("%s <%s>", val, emailURL)
	}
	manifestFile := *&viper.Viper{}
	if commandName == "maintainer" {
		manifestFile = *ctx.Maintainer.File
	} else {
		manifestFile = *ctx.Author.File
	}
	items := manifestFile.GetStringSlice(index)
	items = append(items, val)
	manifestFile.Set(index, items)

	return val, manifestFile.WriteConfig()

}
