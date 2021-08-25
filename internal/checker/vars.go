package checker

import (
	"net/http"
	"sync"
)

var (
	client *http.Client
	myip   myIP
	wg     sync.WaitGroup

	endpoint = "https://www.walmart.com/ip/lol/790390437"
)
