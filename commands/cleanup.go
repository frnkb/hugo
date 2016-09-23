package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// RemoveNonMinified removes the non minified versions of static css/js files from the publish dir
// Only remove the non minified version if the minified version exists.
func RemoveNonMinified(publishDir string) {
	fmt.Println("poop")
	delCount := 0
	files := fileList(publishDir)
	for _, f := range files {
		if f[len(f)-6:] == "min.js" || f[len(f)-7:] == "min.css" {
			nonMinified := strings.Replace(f, ".min", "", -1)
			for _, ff := range files {
				if ff == nonMinified {
					if removeIfExists(ff) {
						delCount++
					}
				}
			}
		}
	}
	fmt.Printf("%d  non-minified files removed.\n ", delCount)
}

func removeIfExists(minified string) bool {
	nonMinified := strings.Replace(minified, ".min", "", -1)
	if err := os.Remove(nonMinified); err == nil {
		return true
	}
	return false
}

// fileList returns a list of .css and .js files that are in the publish directory.
func fileList(dir string) []string {
	fileList := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && (filepath.Ext(path) == ".js" || filepath.Ext(path) == ".css") {
			fileList = append(fileList, path)
		}
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return fileList
}
