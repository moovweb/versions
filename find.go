package versions

import (
	"errors"
	"path/filepath"
)

type FilePath struct {
	Path    string
	Name    string
	Version *Version
}

func FindByName(rootPath string, name string) (*FilePath, error) {
	return FindByNameLatestVersion(rootPath, name, "")
}

func FindByNameAndPatternLatest(rootPath string, name string, versionPattern *Pattern) (*FilePath, error) {
	filePaths, err := FindByNameAndPattern(rootPath, name, versionPattern)

	if err != nil {
		return nil, err
	}

	if len(filePaths) == 0 {
		return nil, errors.New("Found no files named " + name + " in path: " + rootPath)
	}

	return getLatestVersion(filePaths), nil
}

func FindByNameLatestVersion(rootPath string, name string, versionPattern string) (*FilePath, error) {
	filePaths, err := FindByNameAndVersion(rootPath, name, versionPattern)

	if err != nil {
		return nil, err
	}

	if len(filePaths) == 0 {
		return nil, errors.New("Found no files named " + name + " in path: " + rootPath)
	}

	return getLatestVersion(filePaths), nil
}

func getLatestVersion(filePaths []*FilePath) *FilePath {
	newestFilePath := filePaths[0]

	for _, thisFilePath := range filePaths {
		pattern, _ := NewPattern("> " + newestFilePath.Version.String())

		if pattern.Match(thisFilePath.Version) {
			newestFilePath = thisFilePath
		}
	}

	return newestFilePath
}

func FindByNameAndVersion(rootPath string, name string, versionPattern string) (paths []*FilePath, err error) {

	query := filepath.Join(rootPath, name) + "*"
	results, _ := filepath.Glob(query)

	for _, result := range results {
		version, err := GetVersion(result)

		if err != nil {
			println("Didnt get valid version from:" + result)
			println(err.Error())
			continue
		}

		matched, err := version.Matches(versionPattern)

		if err != nil {
			return nil, err
		}

		if matched {
			matchingPath := &FilePath{
				Path:    result,
				Name:    name,
				Version: version,
			}
			paths = append(paths, matchingPath)
		}
	}

	return
}

func FindByNameAndPattern(rootPath string, name string, versionPattern *Pattern) (paths []*FilePath, err error) {
	// TODO(JE): Refactoring to share the same code in this function and FindByNameAndVersion
	query := filepath.Join(rootPath, name) + "*"
	results, _ := filepath.Glob(query)

	for _, result := range results {
		version, err := GetVersion(result)

		if err != nil {
			println("Didnt get valid version from:" + result)
			println(err.Error())
			continue
		}

		if versionPattern.Match(version) {
			matchingPath := &FilePath{
				Path:    result,
				Name:    name,
				Version: version,
			}
			paths = append(paths, matchingPath)
		}
	}

	return
}

func (fp *FilePath) String() string {
	// TODO(SJ) -- support any extension and save the actual filename!
	return fp.Path
}
