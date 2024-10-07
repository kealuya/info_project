package main

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "http://40bc87a3d761e5faa168806d816c65ae@123.187.240.255:9000/2",

		EnableTracing:    true,
		Debug:            false,
		AttachStacktrace: true,
		TracesSampleRate: 1.0,
		Release:          "mytest v3.0",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	//defer sentry.Flush(2 * time.Second)

	//sentry.CaptureMessage("It works!")

	v()

	time.Sleep(4 * time.Second)
	fmt.Println("LastEventID::", sentry.LastEventID())
}

func v() {

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("發送")

			sentry.WithScope(func(scope *sentry.Scope) {

				// Add request information
				scope.SetRequest(&http.Request{
					Method: "GET",
					URL: &url.URL{
						Scheme: "http",
						Host:   "www.baidu.com",
					},
				})

				// Add user information if available
				scope.SetUser(sentry.User{
					ID:    "renhao",
					Email: "renhao@example.com",
				})

				// Add custom tags or extra data
				scope.SetExtra("extra-info", "some extra information")
				scope.SetExtras(map[string]any{
					"picket": "sdf",
					"age":    23,
				})

				// Capture the error
				eventId := sentry.CaptureException(fmt.Errorf("%v", err))
				fmt.Println("eventId::", eventId)
			})
			sentry.Flush(time.Second * 5)
		}
	}()

	response, err := http.Post("http://127.0.0.1:8000/hello", "text/plain", strings.NewReader("hello world"))
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
