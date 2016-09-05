package confloader

import "os"

type IniLoader struct {
	iniFilePath string = "conf.ini"
}

func (loader IniLoader) Load() {
	if _, err := os.Stat(loader.iniFilePath); err == nil {
		
	}
}
