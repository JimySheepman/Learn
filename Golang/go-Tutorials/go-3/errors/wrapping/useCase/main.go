package usecase

import "io/ioutil"

func transferFileContents(filename string) error {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("/tmp/filecontents", contents, 0644)
	if err != nil {
		return err
	}
	return nil
}
