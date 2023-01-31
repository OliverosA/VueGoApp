package files

func ParseFilesByPath(filePaths []string) []EmailJsonData {
	// email struct to save all the email info
	var jsonDocs []EmailJsonData

	for _, path := range filePaths {
		/// adding email info into jsonDocs
		jsonDocs = append(jsonDocs, ParseEmailFileToJson(path))
	}

	return jsonDocs
}
