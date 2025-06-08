package main

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// Create a new allocator context for managing Chrome instances
	allocatorContext, cancel := chromedp.NewRemoteAllocator(context.Background(), "ws://localhost:8088")
	defer cancel()

	// Build context options (currently empty, but can be used for configuration)
	var opts []chromedp.ContextOption

	// Create a new Chrome instance using the allocator and options
	ctx, cancel := chromedp.NewContext(
		allocatorContext,
		opts...,
	)
	defer cancel()

	// "processBrowser" processes the browser's targets
	processBrowser := func() {
		// Retrieve the list of targets (tabs/windows) in the Chrome instance
		targets, err := chromedp.Targets(ctx)
		if err != nil {
			log.Fatal(err)
			return
		}

		// Flag to check if the target is found
		var found bool

		// Iterate through the targets to find the desired one
		for _, t := range targets {
			if strings.Contains(t.Title, "Skyrim Special Edition Nexus") {
				found = true
				log.Printf("Target: %s", t.Title)
				ctx, cancel := chromedp.NewContext(ctx, chromedp.WithTargetID(t.TargetID))
				defer cancel()

				// Create a timeout
				ctx, cancel = context.WithTimeout(ctx, 1*time.Second)
				defer cancel()

				// Set a timeout for ad check
				var adVisible bool

				// Check if an ad is visible
				err := chromedp.Run(ctx,
					chromedp.Evaluate(`!!document.querySelector('input.close-btn[type="checkbox"]')`, &adVisible),
				)

				// Log error without exiting
				if err != nil {
					log.Printf("Error checking for ad popup: %s", err.Error())
				}

				// If an ad is visible, set a timeout for closing it
				if adVisible {
					ctxCloseAd, cancelCloseAd := context.WithTimeout(ctx, 10*time.Second)
					defer cancelCloseAd()
					err = chromedp.Run(ctxCloseAd,
						chromedp.WaitVisible(`input.close-btn[type="checkbox"]`), // Wait until the ad's close button is visible
						chromedp.Click(`input.close-btn[type="checkbox"]`),       // Click the ad's close button
					)
					if err != nil {
						log.Printf("Could not close ad popup: %s", err.Error()) // Log error without exiting
					}
				}

				// Attempt to click the download button
				err = chromedp.Run(ctx,
					chromedp.WaitVisible(`button[id="slowDownloadButton"]`),
					chromedp.Click(`button[id="slowDownloadButton"]`),
				)

				// Log error without exiting
				if err != nil {
					log.Printf("Could not fulfill new target with error %s", err.Error())
				}
				cancel()
			}
		}

		// Log if no target was found
		if !found {
			log.Printf("No target found")
		}
	}

	// Log if no target was found
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
