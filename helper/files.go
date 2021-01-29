package helper

import (
	"fmt"
	"io/ioutil"
	"strings"

	"io"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"gitlab.aws.dev/devops-aws/terraform-ce-cli/box"
)

// BuildPath changes the given path to a more cross platform friendly format
func BuildPath(path string) string {
	sep := string(os.PathSeparator)
	return strings.ReplaceAll(path, "/", sep)
}

// WriteFile writes a file and return true if successful
func WriteFile(filename string, data []byte) bool {
	err := ioutil.WriteFile(filename, data, os.ModePerm)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

// WriteFileFromBox get the file from box's resources and write into the given destination, returns false if not able to.
func WriteFileFromBox(source string, dest string) bool {
	bytes, found := box.Get(BuildPath(source))

	if !found {
		log.Errorf("file \"%s\" not found under box/resources", BuildPath(source))
		return false
	}

	return WriteFile(BuildPath(dest), bytes)
}

// DownloadFileTo downloads a file and saves into the given directory with the given file name
func DownloadFileTo(url string, destination string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(BuildPath(destination))
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// DownloadFile downloads a file and saves into the given directory with the given file name
func DownloadFile(url string, dirPath string, filename string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(dirPath + "/" + filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// FileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(path string) bool {
	info, err := os.Stat(BuildPath(path))
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DirOrFileExists an error is known to report that a file or directory does not exist.
// It is satisfied by ErrNotExist as well as some syscall errors.
func DirOrFileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// CopyFile copies a file from source into a given destination path
// https://github.com/mactsouk/opensource.com/blob/master/cp2.go
func CopyFile(sourceFile string, destinationFile string) {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destinationFile)
		fmt.Println(err)
		os.Exit(1)
	}
}

// CopyFileTo copy a file from sourceFile location to destinationFile location
func CopyFileTo(sourceFile string, destinationFile string) error {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		logrus.Errorf("unable to copy file\n%v\n", err)
		return err
	}

	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		logrus.Errorf("unable to write file\n%v\n", err)
		return err
	}

	return nil
}

// FileSize return the size of the give file path.
// Gives an error if files does not exist
func FileSize(path string) (int64, error) {
	var size int64 = -1
	if FileExists(path) {
		info, err := os.Stat(path)
		if err != nil {
			if err != nil {
				return size, fmt.Errorf("unable to obtain information about file: %s\n%s", path, err)
			}
			return size, err
		}
		size = info.Size()
	} else {
		return size, fmt.Errorf("file does not exist")
	}
	return size, nil
}

// ListFiles list of all file names in the given directory. Pass "." if you want to list at the current directory.
func ListFiles(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	return files
}

// DeleteFile removes the named file or (empty) directory.
// If there is an error, it will be of type *PathError.
func DeleteFile(name string) error {
	return os.Remove(name)
}
