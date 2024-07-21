package ascii_gen

import (
	"fmt"
	"os"
	"strings"
)

var Banners = map[string][][]string{}

func CheckBanner(bannerPath string) bool {
	_, err := os.ReadFile(bannerPath)

	return err == nil
}

func BannerParser(bannerPath string) ([][]string, error) {
	_, ok := Banners[bannerPath]
	// Check if the key (bannerPath) present in the map
	// if not exists!
	if !ok {
		bannerContent, err := os.ReadFile(bannerPath)

		if err != nil {
			fmt.Println("Error when we read the file", err)
			return [][]string{}, err
		}
		// Cast byte with string and replace \r with nothing
		bannerString := strings.ReplaceAll(string(bannerContent), "\r", "")

		// Add bannerPath (key) to Map and skip the first line
		Banners[bannerPath] = FileSpliter(bannerString[1:])

	}

	return Banners[bannerPath], nil
}
