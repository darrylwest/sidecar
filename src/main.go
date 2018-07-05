//
// main
//
// @author darryl.west <darwest@ebay.com>
// @created 2017-09-01 09:43:14
//

package main

import "app"

func main() {
	app.CreateLogger()
	cfg := app.ParseArgs()
	if cfg == nil {
		app.ShowHelp()
		return
	}

	service, err := app.NewService(cfg)
	if err != nil {
		panic(err)
	}

	service.Start()
}
