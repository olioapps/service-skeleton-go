package middleware

import (
	"net/url"
	"regexp"

	"github.com/gin-gonic/gin"
)

type WhiteList struct {
	routeRegexes []*regexp.Regexp
}

func (this *WhiteList) isWhiteListed(url *url.URL) bool {
	for _, regex := range this.routeRegexes {
		match := regex.FindStringSubmatch(url.Path)
		if len(match) > 0 {
			return true
		}
	}

	return false
}

func (this *WhiteList) Handler(c *gin.Context) {
	if this.isWhiteListed(c.Request.URL) {
		c.Set("whitelisted", true)
	}
	c.Next()
}

func (this *WhiteList) AddWhiteListedURLs(regexes []string) {
	for _, regex := range regexes {
		compiled := this.compile(regex)
		if compiled == nil {
			panic("Invalid url index:" + regex)
		} else {
			this.routeRegexes = append(this.routeRegexes, compiled)
		}
	}
}

func (this *WhiteList) compile(regexStr string) *regexp.Regexp {
	r, err := regexp.Compile(regexStr)
	if err != nil {
		return nil
	}

	return r
}

func NewWhitelist() *WhiteList {
	w := WhiteList{}
	w.routeRegexes = []*regexp.Regexp{}

	return &w
}
