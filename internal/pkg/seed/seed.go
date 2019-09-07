package seed

import (
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"time"

	"gopkg.in/src-d/go-git.v4"
)

type Seed struct {
	path string
}

func (s *Seed) GetPath() string {
	return s.path
}

func (s *Seed) GetSeedfilePath() string {
	return path.Join(
		s.GetPath(),
		FindSeedFilename(s.GetPath()),
	)
}

func NewSeed(str string) (*Seed, error) {
	if str == "" {
		return nil, errors.New("Need to provide a value")
	}

	var seedPath string
	if isSeedDir(str) {
		seedPath = str
	} else {
		if !isGitRepo(str) {
			str = fmt.Sprintf("https://github.com/sourceseed/seed-%s.git", str)
		}
		p, err := cloneRepo(str)
		seedPath = p
		if err != nil {
			return nil, err
		}
	}

	return &Seed{
		path: seedPath,
	}, nil
}

func isGitRepo(str string) bool {
	match, _ := regexp.MatchString("(?:git|ssh|https?|git@[-\\w.]+):(\\/\\/)?(.*?)(\\.git)(\\/?|\\#[-\\d\\w._]+?)$", str)

	return match
}

func cloneRepo(repo string) (string, error) {
	timeString := strconv.FormatInt(time.Now().Unix(), 10)

	tmpDir := path.Join("/tmp", "seedr", timeString)

	_, err := git.PlainClone(tmpDir, false, &git.CloneOptions{
		URL: repo,
	})

	return tmpDir, err
}

func isSeedDir(str string) bool {
	return isDir(str)
}

func FindSeedFilename(str string) string {
	seedFiles := []string{
		"Seedfile.yml",
		"Seedfile",
	}
	for _, f := range seedFiles {
		if fileExists(path.Join(str, f)) {
			return f
		}
	}

	return ""
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func isDir(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
