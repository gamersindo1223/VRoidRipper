package main

import (
	// "VRoidRipper/requests"
	"fmt"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-resty/resty/v2"

	// "net/http"
	"os"
	"regexp"

	"github.com/go-faker/faker/v4"
)

func main() {
	a := app.NewWithID("VRoid Ripper")

	Icon, _ := fyne.LoadResourceFromURLString("https://i.pinimg.com/originals/69/0d/25/690d2555dc24257e693e43a79a2e1afc.png")
	mainWindow := a.NewWindow("VRoid Ripper")
	mainWindow.SetIcon(Icon)
	mainWindow.Resize(fyne.NewSize(400, 0))
	mainWindow.SetMaster()
	regex := regexp.MustCompile(`https:\/\/hub.vroid.com\/en\/characters\/[\w-]{19}\/models\/`)

	entry := widget.NewEntry()

	button := widget.NewButton("Rip VRM", func() {
		if entry.Text != "" || regex.MatchString(entry.Text) {
			item := regex.ReplaceAllString(entry.Text, "")
			fmt.Println("Downloading...")
			resp, err := doRequest("https://hub.vroid.com/api/character_models/"+item+"/optimized_preview", entry.Text)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = os.WriteFile(item+".vrm", resp.Body(), 0755)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Finished!")
		} else {
			fmt.Println("Bad Input")
		}
	})

	content := container.NewVBox(
		widget.NewLabel("Enter VRoid character model URL:"),
		entry,
		button,
	)

	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}

func doRequest(url, ref string) (*resty.Response, error) {
	ua, err := faker.GetUserAgent().UserAgent(reflect.Value{})
	if err != nil {
		return nil, err
	}
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetHeader("accept", "*/*").
		SetHeader("accept-language", "en-US,en;q=0.9,id;q=0.8").
		SetHeader("cache-control", "no-cache").
		SetHeader("pragma", "no-cache").
		SetHeader("priority", "u=1, i").
		SetHeader("referer", ref).
		// SetHeader("sec-ch-ua", `"Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"`).
		// SetHeader("sec-ch-ua-mobile", "?0").
		// SetHeader("sec-ch-ua-platform", `"Windows"`).
		// SetHeader("sec-fetch-dest", "empty").
		SetHeader("sec-fetch-mode", "cors").
		SetHeader("sec-fetch-site", "same-origin").
		SetHeader("user-agent", ua.(string)).
		SetHeader("x-api-version", "11").
		Get(url)
	if err != nil {
		return nil, err
	}
	return resp, err
}
