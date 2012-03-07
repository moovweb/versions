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
	filePaths, err := FindByNameAndVersion(rootPath, subPath, name, "")
	
	if err != nil {
		return nil, err
	}

	if len(filePaths) == 0 {
		return nil, os.NewError("Found no files named " + name + " in path: " + rootPath)
	}

	newestFilePath := filePaths[0]
	

	for _, thisFilePath := range(filePaths) {
		pattern, _ := NewPattern("> " + newestFilePath.Version.String() )

		if pattern.Match(thisFilePath.Version) {
			newestFilePath = thisFilePath
		}
	}

	return newestFilePath, nil
}

func FindByNameAndVersion(rootPath string, subPath string, name string, versionPattern string) (paths []*FilePath, err os.Error) {
	path := filepath.Join(rootPath, subPath)
	query := filepath.Join(path, name) + "*"
	results, _ := filepath.Glob(query)
	
	for _, result := range(results) {
		version, err := GetVersion(result)

		if err != nil {
			return nil, err
		}

		matched, err := version.Matches(versionPattern)
		
		if err != nil {
			return nil, err
		}

		if matched {
			matchingPath := &FilePath{
			Path: result,
			Name: name,
			Version: version,
			}
			paths = append(paths, matchingPath)
		}
	}

	return
}