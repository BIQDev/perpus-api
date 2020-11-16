package booklist

import "github.com/BIQDev/perpus-api/internal/router"

var routes = []router.BIQRoute{
	{Name: "/booklist/create", Methods: []string{"POST"}, Pattern: "/booklist/{username}/create", HandlerFunc: BooklistHandlers.Create},
}

func init() {
	router.BIQRouteApply(routes)
}
