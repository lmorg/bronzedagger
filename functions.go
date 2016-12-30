package main

/*
var rx_tsv *regexp.Regexp = regexCompile(`\s+`)

func regexCompile(s string) *regexp.Regexp {
	r, err := regexp.Compile(s)
	isErr(err)
	return r
}
*/

func lower(val int) int {
	return int(val/fRounding) * fRounding
}

func noBlankStr(s string) string {
	if s == "" {
		return "-"
	}
	return s
}

/*
func readConfig(filename string) (config [][]string) {
	if filename == "" {
		return
	}

	b, err := ioutil.ReadFile(filename)
	if !isErr(err) {
		lines := strings.Split(string(b), "\n")
		for _, l := range lines {
			config = append(config, rx_tsv.Split(l, -1))
		}
	}

	return
}
*/
