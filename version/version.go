package version

import (
	"fmt"
	"os"
	"regexp"

	"github.com/josehbez/pm"
	"github.com/spf13/cobra"
)

//Command ...
type Command struct {
}

// Run ...
func (c Command) Run(ctx *pm.Ctx) *cobra.Command {

	var run = func(cmd *cobra.Command, args []string) {
		if ok, _ := cmd.Flags().GetBool("major"); ok {
			c.major(ctx, cmd.Name())
		} else if ok, _ := cmd.Flags().GetBool("minor"); ok {
			c.minor(ctx, cmd.Name())
		} else if ok, _ := cmd.Flags().GetBool("patch"); ok {
			c.patch(ctx, cmd.Name())
		} else if ok, _ := cmd.Flags().GetString("label"); len(ok) > 0 {
			c.label(ctx, cmd.Name(), ok)
		} else if ok, _ := cmd.Flags().GetBool("full"); ok {
			c.status(ctx, "full")
			os.Exit(0)
		} else if ok, _ := cmd.Flags().GetString("check"); len(ok) > 0 {
			ctx.Out.Println(c.check(ok))
			os.Exit(0)
		}
		c.status(ctx, cmd.Name())
	}
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "Semantic versioning specification",
		Long: `
Examples:
 $ pm version -x
 $ pm version -y
 $ pm version -z

 $ pm version -f 
 $ 1.0.1-alfa.1+exp.sha.1
 
 $ pm version -c 1.0.1-alfa.1+exp.sha.1
 $ true
 `,
		Run: func(cmd *cobra.Command, args []string) {
			run(cmd, args)
		},
	}
	usageMajor := "version when you make incompatible API changes"
	usageMinor := "version when you add functionality in a backwards compatible manner"
	usagePatch := "version when you make backwards compatible bug fixe"
	usageRemove := "remove node metadata"
	//cmd.Flags().StringP("label", "l", "", "")
	cmd.Flags().BoolP("major", "x", false, usageMajor)
	cmd.Flags().BoolP("minor", "y", false, usageMinor)
	cmd.Flags().BoolP("patch", "z", false, usagePatch)
	//cmd.Flags().BoolP("remove", "r", false, "")
	cmd.Flags().BoolP("full", "f", false, "print full version Version-PreRealase+Build")
	cmd.Flags().StringP("check", "c", "", "check if version is based on semver.org")

	var cmdRelease = &cobra.Command{
		Use:   "pre-release",
		Short: "version indicates that the version is unstable",
		Long: `
Software release cycle life: [[alfa < beta < release-candidate] release | release-tls < discontinued]

Examples:
 $ pm version pre-release -l alfa
 $ pm version pre-release -x 
 
 $ pm version pre-relase
 $ alfa.1
`,
		Run: func(cmd *cobra.Command, args []string) {
			run(cmd, args)
		},
	}
	cmdRelease.Flags().StringP("label", "l", "", "set pre-release name")
	cmdRelease.Flags().BoolP("major", "x", false, usageMajor)
	//cmdRelease.Flags().BoolP("minor", "y", false, "")
	//cmdRelease.Flags().BoolP("patch", "z", false, "")
	cmdRelease.Flags().BoolP("remove", "r", false, usageRemove)
	cmd.AddCommand(cmdRelease)

	var cmdBuild = &cobra.Command{
		Use:   "build",
		Short: "build metadata",
		Long: `
Examples:
 $ pm version build -l exp.sha
 $ pm version build -x 
 
 $ pm version build
 $ exp.sha.1

		`,
		Run: func(cmd *cobra.Command, args []string) {
			run(cmd, args)
		},
	}
	cmdBuild.Flags().StringP("label", "l", "", "set build name")
	cmdBuild.Flags().BoolP("major", "x", false, usageMajor)
	//cmdBuild.Flags().BoolP("minor", "y", false, "")
	//cmdBuild.Flags().BoolP("patch", "z", false, "")
	cmdBuild.Flags().BoolP("remove", "r", false, usageRemove)
	cmd.AddCommand(cmdBuild)

	return cmd
}

//Version ...
type Version struct {
	Major int
	Minor int
	Patch int
	Label string
}

func getManifestPath(name string) string {
	if name == "version" {
		return name
	}
	return fmt.Sprintf("version.%s", name)
}

func (c *Command) check(version string) bool {
	r, _ := regexp.Compile("^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$")
	return r.MatchString(version)
}

func (c *Command) status(ctx *pm.Ctx, name string) {
	version := ""
	var getVersion = func(name string) string {
		v := Version{
			Label: ctx.Manifest.GetString(fmt.Sprintf("%s.label", getManifestPath(name))),
			Major: ctx.Manifest.GetInt(fmt.Sprintf("%s.major", getManifestPath(name))),
			Minor: ctx.Manifest.GetInt(fmt.Sprintf("%s.minor", getManifestPath(name))),
			Patch: ctx.Manifest.GetInt(fmt.Sprintf("%s.patch", getManifestPath(name))),
		}
		if name == "version" {
			return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
		}

		v2 := ""
		if len(v.Label) > 0 {
			v2 = v.Label
		}
		if v.Major > 0 {
			if len(v.Label) > 0 {
				v2 = fmt.Sprintf("%s.%d", v.Label, v.Minor)
			} else {
				v2 = fmt.Sprintf("%d", v.Minor)
			}
		}
		return v2
	}
	if name == "full" {
		version = getVersion("version")
		p := getVersion("pre-release")
		b := getVersion("build")
		if len(p) > 0 {
			version = fmt.Sprintf("%s-%s", version, p)
		}
		if len(b) > 0 {
			version = fmt.Sprintf("%s+%s", version, b)
		}
	} else {
		version = getVersion(name)
	}

	ctx.Out.Printf(version)
}

func (c *Command) label(ctx *pm.Ctx, name, value string) error {
	ctx.Manifest.Set(fmt.Sprintf("%s.label", getManifestPath(name)), value)
	return ctx.Manifest.WriteConfig()
}

func (c *Command) major(ctx *pm.Ctx, name string) error {

	major := ctx.Manifest.GetInt(fmt.Sprintf("%s.major", getManifestPath(name))) + 1
	ctx.Manifest.Set(fmt.Sprintf("%s.major", getManifestPath(name)), major)
	ctx.Manifest.Set(fmt.Sprintf("%s.minor", getManifestPath(name)), 0)
	ctx.Manifest.Set(fmt.Sprintf("%s.patch", getManifestPath(name)), 0)
	return ctx.Manifest.WriteConfig()
}

func (c *Command) minor(ctx *pm.Ctx, name string) error {
	minor := ctx.Manifest.GetInt(fmt.Sprintf("%s.minor", getManifestPath(name))) + 1
	ctx.Manifest.Set(fmt.Sprintf("%s.minor", getManifestPath(name)), minor)
	ctx.Manifest.Set(fmt.Sprintf("%s.patch", getManifestPath(name)), 0)
	return ctx.Manifest.WriteConfig()
}

func (c *Command) patch(ctx *pm.Ctx, name string) error {
	patch := ctx.Manifest.GetInt(fmt.Sprintf("%s.patch", getManifestPath(name))) + 1
	ctx.Manifest.Set(fmt.Sprintf("%s.patch", getManifestPath(name)), patch)
	return ctx.Manifest.WriteConfig()

}

func (c *Command) remove(ctx *pm.Ctx, name string) error {
	ctx.Manifest.SetDefault(fmt.Sprintf("%s", getManifestPath(name)), nil)
	return ctx.Manifest.WriteConfig()
}
