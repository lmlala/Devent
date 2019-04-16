/**
 * @author [Hilton Li]
 * @email [liminghilton@gmail.com]
 * @create date 2019-04-15 22:04:36
 * @modify date 2019-04-15 22:04:36
 * @desc [description]
 */

package Events

import (
	"time"
	"fmt"
)

type DateTime struct {
	time.Time
}

// // Convert the internal date as CSV string
// func (date *DateTime) MarshalCSV() (string, error) {
// 	return date.Time.Format("2006-2-01 15:04 PM"), nil
// }


// Convert the CSV string as internal date
func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("2006-01-02 15:04:05", csv)
	if err != nil {
		fmt.Printf("Error when parse time: %s, %s\n", csv, err.Error())
		date.Time, _ = time.Parse("2006-01-02 15:04:05", "2000-01-01 00:00:00")
	}
	return nil
}

type Event struct {
	Type	string	`json:"event_type"`
	ID int	`json:"ID"`
	Summary string	`json:"summary"`
	Url	string	`json:"url"`
	Priority	string	`json:"priority"`
	Assignee int	`json:"assignee	"`
	Description	string	`json:"description"`
	StartTime	DateTime	`json:"start"`
	EndTime	DateTime	`json:"end"`
	LastTime	int	`json:"last_time"`
	Title	string	`json:"title"`
	Component []string	`json:"component"`
	Env	string	`json:"env"`
	// todo: env and component
}

func ListEvent() (r []*Event, err error) {

	incidents, err := ReadIncident()

	if err != nil {
		return r, err
	}

	deployments, err := ReadDepolyment()

	if err != nil {
		return r, err
	}

	r = append(incidents,deployments...)

	return r, nil
}