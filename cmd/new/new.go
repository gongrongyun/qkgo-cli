package new

import (
	"github.com/fatih/color"
	"github.com/gongrongyun/qkgo-cli/static"
	"github.com/gongrongyun/qkgo-cli/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
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

var (
	appName string
)

func newApp(cmd *cobra.Command, args []string)  {
	appName = args[0]
	path, err := cmd.Flags().GetString("path")
	if err != nil {
		path = "./"
	}

	// check if the project existed
	appPath := path + appName
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

	err = deepClone(appPath, "template")
	if err != nil {
		_ = os.Remove(appPath)
		color.Red("[Error]---fail to create template files; err=[%v]", err)
		return
	}

	color.Green("successfully create your project!\nplease read the README.md carefully and start your travel!")
}

func deepClone(curAppPath string, dir string) error {
	if isDir(dir) {
		err := os.Mkdir(curAppPath, 0755)
		files, err := static.AssetDir(dir)
		if err != nil {
			return err
		}
		for _, file := range files {
			err = deepClone(curAppPath + "/" + file, dir + "/" + file)
			if err != nil {
				return err
			}
		}
	} else {
		bytes, err := static.Asset(dir)
		if err != nil {
			return err
		}
		text := strings.Replace(string(bytes), oldProjectName, appName, -1)
		err = ioutil.WriteFile(curAppPath, []byte(text), 0755)
		if err != nil {
			return err
		}
		color.Yellow("create " + curAppPath + " successfully!")
	}

	return nil
}

func isDir(path string) bool {
	_, err := static.AssetDir(path)
	if err != nil {
		return false
	}
	return true
}

func init()  {
	Command.Flags().StringP("path", "p", "./", "path of project")
}