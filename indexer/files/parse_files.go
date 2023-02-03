package files

func ParseFilesByPath(filePaths []string) []EmailJsonData {
	var jsonDocs []EmailJsonData

	for _, path := range filePaths {
		jsonDocs = append(jsonDocs, ParseEmailFileToJson(path))
	}

	return jsonDocs
}
