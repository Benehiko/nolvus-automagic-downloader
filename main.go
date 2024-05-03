package main

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	allocatorContext, cancel := chromedp.NewRemoteAllocator(context.Background(), "ws://localhost:8088")
	defer cancel()

	// build context options
	var opts []chromedp.ContextOption
	// opts = append(opts, chromedp.WithDebugf(log.Printf))
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		allocatorContext,
		opts...,
	)
	defer cancel()

	processBrowser := func() {

		targets, err := chromedp.Targets(ctx)
		if err != nil {
			log.Fatal(err)
			return
		}

		var found bool
		for _, t := range targets {
			if strings.Contains(t.Title, "Skyrim Special Edition Nexus") {
				found = true
				log.Printf("Target: %s", t.Title)
				ctx, cancel := chromedp.NewContext(ctx, chromedp.WithTargetID(t.TargetID))
				defer cancel()

				// create a timeout
				ctx, cancel = context.WithTimeout(ctx, 1*time.Second)
				defer cancel()

				err := chromedp.Run(ctx,
					chromedp.WaitVisible(`button[id="slowDownloadButton"]`),
					chromedp.Click(`button[id="slowDownloadButton"]`),
				)
				if err != nil {
					log.Printf("Could not fulfill new target with error %s", err.Error())
				}
				cancel()
			}
		}

		if !found {
			log.Printf("No target found")
		}
	}

	for {
		processBrowser()
		time.Sleep(5 * time.Second)
	}

	// targets, err := chromedp.Targets(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Println("Targets: ", targets)
	//
	// var currentTarget *target.Info
	// for _, t := range targets {
	// 	if strings.Contains(t.Title, "Skyrim Special Edition Nexus") {
	// 		currentTarget = t
	// 		break
	// 	}
	// }
	//
	// if currentTarget == nil {
	// 	log.Fatal("currentTarget has not been found")
	// }
	//
	// // build context options
	// opts = append(opts, chromedp.WithTargetID(currentTarget.TargetID))
	//
	// // create chrome instance
	// ctx, cancel = chromedp.NewContext(
	// 	allocatorContext,
	// 	opts...,
	// )
	//
	// defer cancel()
	//
	// // create a timeout
	// ctx, cancel = context.WithTimeout(ctx, 1*time.Second)
	// defer cancel()
	//
	// err = chromedp.Run(ctx,
	// 	chromedp.WaitVisible(`button[id="slowDownloadButton"]`),
	// 	chromedp.Click(`button[id="slowDownloadButton"]`),
	// )
	// if err != nil {
	// 	panic(err)
	// }

}
