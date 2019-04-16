/**
 * @author [Hilton Li]
 * @email [liminghilton@gmail.com]
 * @create date 2019-04-16 14:54:33
 * @modify date 2019-04-16 14:54:33
 * @desc [description]
 */

package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	// "path/filepath"
	// "os"
)

type mainConfig struct {
	EventData map[string]string	`yaml:"event_data"`
}

var (
	CFG	mainConfig
)

func InitConfig(absfilepath string) (error) {
	// fmt.Println(filepath.Abs(filepath.Dir(os.Args[0])))
	f, err := ioutil.ReadFile(absfilepath + "/config.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(f, &CFG)

	return err
}