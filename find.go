package versions

import(
	"os"
	"path/filepath"
	)

type FilePath struct{
	Path string
	Name string
	Version *Version
}


func FindByName(rootPath string, subPath string, name string) (*FilePath, os.Error){
	filePaths := FindByNameAndVersion(rootPath, subPath, name, "")
	
	if len(filePaths) == 0 {
		return nil, os.NewError("Found no files named " + name + " in path: " + rootPath)
	}

	newestFilePath := filePaths[0]
	
	for _, thisFilePath := range(filePaths) {
		if thisFilePath.Version.NewerThan(newestFilePath.Version) {
			newestFilePath = thisFilePath
		}
	}

	return newestFilePath, nil
}

func FindByNameAndVersion(rootPath string, subPath string, name string, versionPattern string) (paths []*FilePath) {
	path := filepath.Join(rootPath, subPath)
	query := filepath.Join(path, name) + "*"
	results, _ := filepath.Glob(query)
	
	for _, result := range(results) {
		version, err := GetVersion(result)

		if err == nil {
			if version.Matches(versionPattern) {
				matchingPath := &FilePath{
				Path: result,
				Name: name,
				Version: version,
				}
				paths = append(paths, matchingPath)
			}
		}
	}

	return paths
}