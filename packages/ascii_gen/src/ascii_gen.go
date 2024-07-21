package ascii_gen

func AsciiGenerator(input, banner string) (string, error) {
	bannerSlice, err := BannerParser(banner)

	if err != nil {
		return "", err
	}

	return Writer(input, bannerSlice)
}
