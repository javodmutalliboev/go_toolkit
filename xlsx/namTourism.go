package xlsx

import (
	"log"

	"github.com/javodmutalliboev/go_toolkit/environment"
	"github.com/javodmutalliboev/go_toolkit/postgresql"
	"github.com/javodmutalliboev/go_toolkit/struct_package"
	tealeg_xlsx "github.com/tealeg/xlsx"
)

func init() {
	environment.Load()
}

func Export() {
	// Open the xlsx file
	xlFile, err := tealeg_xlsx.OpenFile("namangan-travel.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgresql.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("TRUNCATE TABLE event RESTART IDENTITY CASCADE;")
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the sheets
	for _, sheet := range xlFile.Sheets {
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		// Iterate through the rows
		for i, row := range sheet.Rows {
			if i == 0 {
				continue
			}

			if i == 50 {
				break
			}

			var event struct_package.Event
			// Iterate through the cells
			for j, cell := range row.Cells {
				switch j {
				case 1:
					event.Name = cell.String()
				case 2:
					event.Description = cell.String()
					/*
						case 3:
							event.Photos = append(event.Photos, struct_package.EventPhoto{File: []byte(cell.Value)})
						case 4:
							event.Photos = append(event.Photos, struct_package.EventPhoto{File: []byte(cell.Value)})
						case 5:
							event.Photos = append(event.Photos, struct_package.EventPhoto{File: []byte(cell.Value)})
						case 6:
							event.Photos = append(event.Photos, struct_package.EventPhoto{File: []byte(cell.Value)})
					*/
				case 3:
					video := cell.String()
					event.Video = &video
				case 4:
					mapData := cell.String()
					event.Map = &mapData
				}
			}

			// Insert the event and return the ID
			err = tx.QueryRow("INSERT INTO event (name, description, video, map) VALUES ($1, $2, $3, $4) RETURNING id", event.Name, event.Description, event.Video, event.Map).Scan(&event.ID)
			if err != nil {
				tx.Rollback()
				log.Fatal(err)
			}

			// Insert the photos
			for _, photo := range event.Photos {
				_, err = tx.Exec("INSERT INTO event_photo (event_id, file) VALUES ($1, $2)", event.ID, photo.File)
				if err != nil {
					tx.Rollback()
					log.Fatal(err)
				}
			}
		}
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}
	}
}
