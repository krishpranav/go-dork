package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/logrusorgru/aurora/v3"
	log "github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
)

// helper function
func init() {
	flag.StringVar(&query, "q", "", "")
	flag.StringVar(&query, "query", "", "")

	flag.StringVar(&engine, "e", "", "")
	flag.StringVar(&engine, "engine", "google", "")

	flag.IntVar(&page, "p", 1, "")
	flag.IntVar(&page, "page", 1, "")

	flag.Var(&headers, "header", "")
	flag.Var(&headers, "H", "")

	flag.StringVar(&proxy, "x", "", "")
	flag.StringVar(&proxy, "proxy", "", "")

	flag.BoolVar(&silent, "s", false, "")
	flag.BoolVar(&silent, "silent", false, "")

	flag.Usage = func() {
		h := []string{
			"Options:",
			"  -q, --query <query>          Search query",
			"  -e, --engine <engine>        Provide search engine (default: Google)",
			"                               (options: Google, Shodan, Bing, Duck, Yahoo, Ask)",
			"  -p, --page <i>               Specify number of pages (default: 1)",
			"  -H, --header <header>        Pass custom header to search engine",
			"  -x, --proxy <proxy_url>      Use proxy to surfing (HTTP/SOCKSv5 proxy)",
			"  -s, --silent                 Silent mode",
			"\n",
		}
		showBanner()
		fmt.Fprintf(os.Stderr, "%s", aurora.Green(strings.Join(h, "\n")))
	}
	flag.Parse()

	engine = strings.ToLower(engine)

	maxLog := levels.LevelDebug
	if silent {
		maxLog = levels.LevelSilent
	}
	log.DefaultLogger.SetMaxLevel(maxLog)

	showBanner()
}
