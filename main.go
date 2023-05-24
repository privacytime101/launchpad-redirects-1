package main

import (
	"launchpad-redirects/lists"
	"net/url"

	"github.com/gin-gonic/gin"
)

type instance struct {
	URL         string
	FaviconURL  string
	Name        string
	Description string
}

func main() {
	app := gin.Default()

	app.Static("/static", "./static")

	app.LoadHTMLGlob("templates/*")

	app.Use(func(c *gin.Context) {
		// add strict CSP

		c.Header("Content-Security-Policy", "default-src 'self'; script-src 'none'; img-src 'self'; style-src 'self';")

		c.Next()
	})

	app.GET("/", func(c *gin.Context) {
		inputUrl := c.Query("url")
		if inputUrl != "" {
			// parse url domain
			parsed, err := url.Parse(inputUrl)
			if err != nil {
				c.HTML(200, "home.html", gin.H{
					"Error": "Invalid URL",
				})
				return
			}

			originalParsedHost := parsed.Host

			// check if domain is in list

			if lists.ContainsString(lists.RedditDomains, parsed.Host) {
				// replace domain with discuss.whatever.social
				parsed.Host = "discuss.whatever.social"
			} else if lists.ContainsString(lists.TwitterDomains, parsed.Host) {
				// replace domain with read.whatever.social
				parsed.Host = "read.whatever.social"
			} else if lists.ContainsString(lists.YouTubeDomains, parsed.Host) {
				// replace domain with watch.whatever.social
				parsed.Host = "watch.whatever.social"
			} else if lists.ContainsString(lists.StackOverflowDomains, parsed.Host) {
				// replace domain with code.whatever.social
				parsed.Host = "code.whatever.social"
			} else {
				// send error
				c.HTML(200, "home.html", gin.H{
					"Error": "Unsupported domain",
				})
				return
			}

			if parsed.Host != originalParsedHost {
				// redirect
				c.Redirect(302, parsed.String())
				return
			}

		}

		c.HTML(200, "home.html", gin.H{
			"Instances": []instance{
				{
					URL:         "https://code.whatever.social",
					FaviconURL:  "/static/assets/apps/code.png",
					Name:        "AnonymousOverflow",
					Description: "Whatever's own frontend for StackOverflow and StackExchange.",
				},
				{
					URL:         "https://discuss.whatever.social",
					FaviconURL:  "/static/assets/apps/libreddit.png",
					Name:        "Libreddit",
					Description: "Alternative frontend for Reddit.",
				},
				{
					URL:         "https://read.whatever.social",
					FaviconURL:  "/static/assets/apps/nitter.png",
					Name:        "Nitter",
					Description: "Alternative frontend for Twitter.",
				},
				{
					URL:         "https://watch.whatever.social",
					FaviconURL:  "/static/assets/apps/piped.png",
					Name:        "Piped",
					Description: "Alternative frontend for YouTube.",
				},
				{
					URL:         "https://cringe.whatever.social",
					FaviconURL:  "/static/assets/apps/proxitok.png",
					Name:        "ProxiTok",
					Description: "Alternative frontend for TikTok.",
				},
				{
					URL:         "https://listen.whatever.social",
					FaviconURL:  "/static/assets/apps/hyperpipe.png",
					Name:        "Hyperpipe",
					Description: "Alternative frontend for YouTube Music.",
				},
				{
					URL:         "https://sing.whatever.social",
					FaviconURL:  "/static/assets/apps/dumb.png",
					Name:        "Dumb",
					Description: "Alternative frontend for Genius.",
				},
				{
					URL:         "https://rimgo.whatever.social",
					FaviconURL:  "/static/assets/apps/rimgo.png",
					Name:        "Rimgo",
					Description: "Alternative frontend for Imgur.",
				},
				{
					URL:         "https://notavault.com",
					FaviconURL:  "/static/assets/apps/vaultwarden.png",
					Name:        "Vaultwarden",
					Description: "Whatever's self-hosted instance of Vaultwarden, a fork of Bitwarden. 3 copies of data at all times, including one in Scaleway's nuclear fallout shelter.",
				},
			},
		})
	})

	app.Run(":8080")
}
