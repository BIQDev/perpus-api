package booklist

import "github.com/BIQDev/perpus-api/internal/router"

var routes = []router.BIQRoute{
	{Name: "/booklist/create", Methods: []string{"POST"}, Pattern: "/booklist/{username}", HandlerFunc: BooklistHandlers.Create},
	{Name: "/booklist/read", Methods: []string{"GET"}, Pattern: "/booklist/{username}", HandlerFunc: BooklistHandlers.Read},
	{Name: "/booklist/update", Methods: []string{"PUT"}, Pattern: "/booklist/{username}/{id}", HandlerFunc: BooklistHandlers.Update},
}

func init() {
	router.BIQRouteApply(routes)
}
