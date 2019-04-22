/**
 * @author [Hilton Li]
 * @email [liminghilton@gmail.com]
 * @create date 2019-04-15 16:17:23
 * @modify date 2019-04-15 16:17:23
 * @desc [description]
 */

package module

import (
	// "io/ioutil"

	"os"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/lmlala/Devent/backend/conf"
)

type Incident struct {
	ID            int      `json:"ID" csv:"ID"`
	StartTime     DateTime `json:"start" csv:"start_time"`
	EndTime       DateTime `json:"end" csv:"end_time"`
	LastTime      int      `json:"last_time" csv:"last_time"`
	DownTime      int      `json:"down_time" csv:"down_time"`
	IsDown        bool     `json:"is_down" csv:"is_down"`
	IsServiceDown bool     `json:"is_service_down" csv:"is_service_down"`
	Description   string   `json:"description" csv:"description"`
	RootCause     string   `json:"root_cause" csv:"root_cause"`
	Priority      string   `json:"priority" csv:"priority"`
	Assignee      int      `json:"assignee	" csv:"Assignee	"`
	Postmortem    string   `json:"url" csv:"postmortem"`
	Env           string   `json:"env" csv:"env"`
	Component     []string `json:"component" csv:_`
	// Title	string	`json:"title" csv:_`
}

type Deployment struct {
	ID          int      `json:"ID" csv:"ID"`
	Summary     string   `json:"summary" csv:"Summary"`
	IssueKey    string   `json:"issue_key" csv:"Issue key"`
	Priority    string   `json:"priority" csv:"Priority"`
	Assignee    int      `json:"assignee	" csv:"Assignee	"`
	Description string   `json:"description" csv:"description"`
	StartTime   DateTime `json:"start" csv:"Created"`
	EndTime     DateTime `json:"end" csv:"Created"`
	LastTime    int      `json:"last_time" csv:"Time Spent"`
	Component   []string `json:"component" csv:_`
	// Title	string	`json:"title" csv:_`
}

func ReadIncident() (r []*Event, err error) {

	filename := config.CFG.EventData["csv_path"] + "/" + config.CFG.EventData["incident_file"]

	f, err := os.OpenFile(filename, os.O_RDONLY, 0664)

	var d []*Incident
	// a, _ := ioutil.ReadAll(f)
	// fmt.Println("read file: ", a)
	if err != nil {
		return r, err
	}
	defer f.Close()

	if err = gocsv.UnmarshalFile(f, &d); err != nil {
		return r, err
	}

	for _, i := range d {
		var event Event

		event.Type = "Incident"
		event.ID = 10000 + i.ID // unique id for differenet type
		event.Summary = i.Description
		event.Url = i.Postmortem
		event.Priority = i.Priority
		event.Assignee = i.Assignee
		event.Description = ""
		event.StartTime = i.StartTime.Time
		event.EndTime = i.EndTime.Time
		event.LastTime = i.LastTime
		event.Title = event.Type + ": " + i.Priority
		event.Component = make([]string, 0)
		event.Env = i.Env

		// fmt.Println(i.EndTime)
		// fmt.Println(event.EndTime)
		r = append(r, &event)
	}

	// return r, errors.New("my error")
	return r, nil

}

func ReadDepolyment() (r []*Event, err error) {

	filename := config.CFG.EventData["csv_path"] + "/" + config.CFG.EventData["deployment_file"]

	f, err := os.OpenFile(filename, os.O_RDONLY, 0664)

	var d []*Deployment
	if err != nil {
		return r, err
	}
	defer f.Close()

	if err = gocsv.UnmarshalFile(f, &d); err != nil {
		return r, err
	}

	for _, i := range d {
		var event Event
		var timeDuration = time.Duration(i.LastTime) * time.Second

		event.Type = "Hotfix"
		event.ID = 20000 + i.ID // unique id for differenet type
		event.Summary = i.Summary
		event.Url = "https://jira.com/dop-" + i.IssueKey
		event.Priority = i.Priority
		event.Assignee = i.Assignee
		event.Description = i.Description
		event.StartTime = i.StartTime.Time
		event.EndTime = i.StartTime.Time.Add(timeDuration) // todo: endtime
		event.LastTime = i.LastTime
		event.Title = event.Type + ": " + i.Summary
		event.Component = make([]string, 0)
		event.Env = "todo: i.Env"

		r = append(r, &event)

	}

	// return r, errors.New("my error")
	return r, nil

}
