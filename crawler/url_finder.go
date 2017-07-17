package crawler

import "github.com/mvdan/xurls"

func findUrls(s string) []string {
	return xurls.Strict.FindAllString(s, -1)
}
