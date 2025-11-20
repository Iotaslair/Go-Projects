package assets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GetAssetIds(alertEventId string) []string {
	// In our system this is just a call to another team
	// so using an example website that doesn't exist

	url := fmt.Sprintf("https://example.com/assets/%s", alertEventId)

	resp, err := http.Get(url)

	checkError(err)
	defer resp.Body.Close()

	return parseAssetIdsFromJson(string(resp.Body))
}

func parseAssetIdsFromJson(jsonResp string) []string {
	assetIds := make([]string, 0)

	decoder := json.NewDecoder(strings.NewReader(jsonResp))

	_, err := decoder.Token()
	checkError(err)
	for decoder.More() {
		var assetId string
		err := decoder.Decode(&assetId)
		checkError(err)
		assetIds = append(assetIds, assetId)
	}
	return assetIds
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
