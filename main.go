package main

import (
	// "encoding/json"
	"log"
	"os"
	// "time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type Car struct {
	Name       string   `json:"name"`
	NameUrl    string   `json:"nameUrl"`
	Company    string   `json:"company"`
	CompanyUrl string   `json:"companyUrl"`
	About      *string  `json:"about,omitempty"`
	Hull       *string  `json:"hull,omitempty"`
	Year       int      `json:"year"`
	Premium    bool     `json:"premium"`
	MinPrice   *float64 `json:"minPrice,omitempty"`
	MaxPrice   *float64 `json:"maxPrice,omitempty"`
	Priority   int      `json:"priority"`
}

func main() {
	app := pocketbase.New()

	// app.OnBootstrap().BindFunc(func(e *core.BootstrapEvent) error {
	// 	if err := e.Next(); err != nil {
	// 		return err
	// 	}
	//
	// 	plan, _ := os.ReadFile("./stale_data/carros.json")
	// 	var items []Car
	// 	err := json.Unmarshal(plan, &items)
	//
	// 	collection, err := e.App.FindCollectionByNameOrId("products")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	for _, item := range items {
	// 		record := core.NewRecord(collection)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	//
	// 		record.Set("name", item.Name)
	// 		record.Set("nameUrl", item.NameUrl)
	// 		record.Set("company", item.Company)
	// 		record.Set("companyUrl", item.CompanyUrl)
	// 		record.Set("about", item.About)
	// 		record.Set("hull", item.Hull)
	// 		record.Set("year", time.Date(item.Year, time.January, 1, 0, 0, 0, 0, time.UTC))
	// 		record.Set("premium", item.Premium)
	// 		record.Set("minPrice", item.MinPrice)
	// 		record.Set("maxPrice", item.MaxPrice)
	// 		record.Set("priority", item.Priority)
	//
	// 		if err := e.App.Save(record); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 	}
	//
	// 	return nil
	// })

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
