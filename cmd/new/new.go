package new

import (
	"github.com/fatih/color"
	"github.com/gongrongyun/qkgo-cli/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var Command = &cobra.Command{
	Use: "new appName [-p path]",
	Short: "create a new project",
	Long: "Create a new project",
	Args: cobra.RangeArgs(1,1),
	Run: newApp,
}

const (
	oldProjectName = "template"
)

func newApp(cmd *cobra.Command, args []string)  {
	appName := args[0]
	path, err := cmd.Flags().GetString("path")
	if err != nil {
		path = "./"
	}

	// check if the project existed
	appPath := path + "/" + appName
	if utils.IsExist(appPath) {
		color.Red("[Error]---project has already existed")
		return
	}

	// check env correct
	go111module := os.Getenv("GO111MODULE")
	if go111module != "on" {
		color.Red("[Error]---please set GO111MODULE=on first")
		return
	}

	color.Yellow("creating project" + appName + "......")

	err = filepath.Walk("./template", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		path = utils.SlicePath(path, "template")
		templateCurPath := "./template/" + path
		appCurPath := appPath + "/" + path

		if utils.IsDir(templateCurPath) {
			if err = os.Mkdir(appCurPath, 0755); err != nil {
				return err
			}
		} else {
			text, err := ioutil.ReadFile(templateCurPath)
			var copyText string
			if path == "main.go" {
				copyText = strings.Replace(string(text), oldProjectName, "main", 1)
			}
			copyText = strings.Replace(string(text), oldProjectName, appName, -1)
			err = ioutil.WriteFile(appCurPath, []byte(copyText), 0755)
			if err != nil {
				return err
			} else {
				color.Yellow("created file %v", path)
			}
		}

		return nil
	})
	if err != nil {
		color.Red("[Error]---fail to create template files; err=[%v]", err)
		return
	}

	color.Green("successfully create your project!\nplease read the README.md carefully and start your travel!")
}

func init()  {
	Command.Flags().StringP("path", "p", "./", "path of project")
}