package license

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"text/tabwriter"

	"github.com/josehbez/pm"
	"github.com/spf13/cobra"
)

//Command ...
type Command struct {
}

// Run ...
func (command Command) Run(ctx *pm.Ctx) *cobra.Command {
	// ref: https://www.synopsys.com/blogs/software-security/top-open-source-licenses/
	var subCLI = &cobra.Command{
		Use:   "license",
		Short: "Show or add lincese based on https://spdx.org/licenses",
		Example: `
pm license --list 
pm license --fetch MIT --save
 `,
		Run: func(kwargs *cobra.Command, args []string) {

			show := true
			if ok, _ := kwargs.Flags().GetBool("list"); ok {
				command.list(ctx)
				show = false
			} else if ok, _ := kwargs.Flags().GetString("fetch"); len(ok) > 0 {
				save, _ := kwargs.Flags().GetBool("save")
				command.fetch(ctx, ok, save)
				show = false
			}
			if show {
				command.show(ctx)
			}

		},
	}

	subCLI.Flags().StringP("fetch", "f", "", "fetch license by indentifier")
	subCLI.Flags().BoolP("save", "s", false, "save license")
	subCLI.Flags().BoolP("list", "l", false, "list full name and identifier licenses")

	return subCLI
}

func (command Command) show(ctx *pm.Ctx) {
	pathLicense := path.Join(ctx.WorkingDir, "LICENSE")
	if _, err := os.Stat(pathLicense); os.IsNotExist(err) {
		ctx.Out.Println("Not has LICENSE")
		return
	}
	inFile, err := os.Open(pathLicense)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		ctx.Out.Println(scanner.Text()) // the line
		return
	}
}

func (command Command) list(ctx *pm.Ctx) {
	type License struct {
		LicenseId             string `json:"licenseId"`
		Name                  string `json:"name"`
		IsDeprecatedLicenseId bool   `json:"isDeprecatedLicenseId"`
		IsFsfLibre            bool   `json:"isFsfLibre"`
		IsOsiApproved         bool   `json:"isOsiApproved"`
	}
	type Licenses struct {
		Licenses []License `json:"licenses"`
	}
	url := "https://spdx.org/licenses/licenses.json"
	//url := "https://raw.githubusercontent.com/spdx/license-list-data/master/json/licenses.json"

	res, err := http.Get(url)
	if err != nil {
		ctx.Err.Fatal(err.Error())
	}
	defer res.Body.Close()

	var data Licenses
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		ctx.Err.Fatal(err.Error())
	}
	json.Unmarshal(body, &data)

	w := new(tabwriter.Writer)
	w.Init(ctx.Out.Writer(), 0, 8, 0, '\t', 0)
	for _, l := range data.Licenses {
		if l.IsFsfLibre {
			fmt.Fprintln(w, fmt.Sprintf("%s\t%s", l.LicenseId, l.Name))
		}
	}
	w.Flush()
}
func (command Command) fetch(ctx *pm.Ctx, licenseID string, save bool) {
	url := fmt.Sprintf("https://spdx.org/licenses/%s.txt", licenseID)
	//url := fmt.Sprintf("https://raw.githubusercontent.com/spdx/license-list-data/master/text/%s.txt", licenseID)

	res, err := http.Get(url)
	if err != nil {
		ctx.Err.Fatal(err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		ctx.Err.Fatal(fmt.Errorf("License ID invalid %s ", licenseID))
	}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		ctx.Err.Fatal(err.Error())
	}
	if save {
		f, err := os.OpenFile("LICENSE", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		f.Write(body)
	}
	ctx.Out.Println(string(body))
}
