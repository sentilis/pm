package license

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
pm license --fetch Apache-2.0
pm license --fetch MIT --save
 `,
		Run: func(kwargs *cobra.Command, args []string) {

			show := true
			if ok, _ := kwargs.Flags().GetBool("list"); ok {
				show = false
				l, lErr := command.List(ctx)
				if lErr != nil {
					ctx.Err.Fatalln(lErr)
				}
				w := new(tabwriter.Writer)
				w.Init(ctx.Out.Writer(), 0, 8, 0, '\t', 0)
				for _, l := range l {
					if l.IsFsfLibre {
						fmt.Fprintln(w, fmt.Sprintf("%s\t%s", l.LicenseID, l.Name))
					}
				}
				w.Flush()
			} else if ok, _ := kwargs.Flags().GetString("fetch"); len(ok) > 0 {
				show = false

				f, fErr := command.Fetch(ctx, ok)
				if fErr != nil {
					ctx.Err.Fatalln(fErr)
				}
				if s, _ := kwargs.Flags().GetBool("save"); s {
					if sErr := command.Save(ctx, f); sErr != nil {
						ctx.Err.Fatalln(sErr)
					}
				}
				ctx.Out.Println(f)
			}
			if show {
				s, sErr := command.Show(ctx)
				if sErr != nil {
					ctx.Err.Fatalln(sErr)
				}
				ctx.Out.Println(s)
			}

		},
	}

	subCLI.Flags().StringP("fetch", "f", "", "fetch license by indentifier")
	subCLI.Flags().BoolP("save", "s", false, "save license. required flag -f")
	subCLI.Flags().BoolP("list", "l", false, "list of available licenses")

	return subCLI
}

//Show ...
func (command Command) Show(ctx *pm.Ctx) (string, error) {
	pathLicense := path.Join(ctx.WorkingDir, "LICENSE")
	if _, err := os.Stat(pathLicense); os.IsNotExist(err) {
		return "", err
	}
	inFile, err := ioutil.ReadFile(pathLicense)
	if err != nil {
		return "", err
	}

	return string(inFile), nil
}

//License ..
type License struct {
	LicenseID             string `json:"licenseId"`
	Name                  string `json:"name"`
	IsDeprecatedLicenseID bool   `json:"isDeprecatedLicenseId"`
	IsFsfLibre            bool   `json:"isFsfLibre"`
	IsOsiApproved         bool   `json:"isOsiApproved"`
}

// List ...
func (command Command) List(ctx *pm.Ctx) ([]License, error) {

	type Licenses struct {
		Licenses []License `json:"licenses"`
	}
	url := "https://spdx.org/licenses/licenses.json"
	//url := "https://raw.githubusercontent.com/spdx/license-list-data/master/json/licenses.json"

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data Licenses

	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		return nil, bodyErr
	}
	if jsonErr := json.Unmarshal(body, &data); jsonErr != nil {
		return nil, jsonErr
	}
	return data.Licenses, nil
}

// Fetch ...
func (command Command) Fetch(ctx *pm.Ctx, licenseID string) (string, error) {
	url := fmt.Sprintf("https://spdx.org/licenses/%s.txt", licenseID)
	//url := fmt.Sprintf("https://raw.githubusercontent.com/spdx/license-list-data/master/text/%s.txt", licenseID)

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("License ID invalid %s ", licenseID)
	}
	body, bodyErr := ioutil.ReadAll(res.Body)

	if bodyErr != nil {
		return "", bodyErr
	}
	//if save {

	//}
	return string(body), nil
}

// Save ..
func (command Command) Save(ctx *pm.Ctx, license string) error {
	f, err := os.OpenFile(path.Join(ctx.WorkingDir, "LICENSE"), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, writeErr := f.Write([]byte(license))
	if writeErr != nil {
		return writeErr
	}
	return nil
}

// Remove ..
func (command Command) Remove(ctx *pm.Ctx) error {
	return os.Remove(path.Join(ctx.WorkingDir, "LICENSE"))
}
