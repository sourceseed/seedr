package generate

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/sourceseed/seedr/internal/pkg/seed"
	"github.com/otiai10/copy"
)

type GenerateConfig struct {
	Seed       *seed.Seed
	TargetDir  string
	Parameters map[string]string
}

func Generate(config GenerateConfig) error {
	logrus.Infof("Copy seed to %s", config.TargetDir)
	if _, err := os.Stat(config.TargetDir); !os.IsNotExist(err) {
		logrus.Error("Target dir already exists")
		return errors.New("Target dir already exists")
	}
	err := copy.Copy(config.Seed.GetPath(), config.TargetDir)
	if err != nil {
		return err
	}

	logrus.Info("Cleaning up Seedfile")
	if sf := seed.FindSeedFilename(config.TargetDir); sf != "" {
		os.Remove(path.Join(config.TargetDir, sf))
	}

	logrus.Info("Cleaning up any vcs files")
	os.RemoveAll(path.Join(config.TargetDir, ".git"))

	logrus.Info("Cleaning up any .seedkeep files")
	r, _ := regexp.Compile("/\\.seedkeep$")
	for _, path := range locateFiles(config.TargetDir) {
		if r.MatchString(path) {
			os.Remove(path)
		}
	}

	logrus.Info("Replacing parameters")
	for _, path := range locateFiles(config.TargetDir) {
		replaceParams(path, config.Parameters)
	}
	logrus.Info("Done")

	return err
}

func locateFiles(dir string) []string {
	result := []string{}
	filepath.Walk(
		dir,
		func(path string, info os.FileInfo, err error) error {
			result = append(result, path)
			return nil
		},
	)
	return result
}

func replaceParams(file string, params map[string]string) {
	info, err := os.Stat(file)
	if err != nil {
		logrus.Error(err)
	}
	if !info.IsDir() {
		read, err := ioutil.ReadFile(file)
		if err != nil {
			logrus.Error(err)
		}
		content := string(read)
		for search, replace := range params {
			searchKey := fmt.Sprintf("__%s__", search)
			content = strings.ReplaceAll(content, searchKey, replace)
		}
		err = ioutil.WriteFile(file, []byte(content), 0)
		if err != nil {
			logrus.Error(err)
		}
	}

	newPath := file
	for search, replace := range params {
		searchKey := fmt.Sprintf("__%s__", search)
		newPath = strings.ReplaceAll(newPath, searchKey, replace)
	}

	if file != newPath {
		err = os.Rename(file, newPath)
		if err != nil {
			logrus.Error(err)
		}
	}
}
