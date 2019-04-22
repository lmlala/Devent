/**
 * @author [Hilton Li]
 * @email [liminghilton@gmail.com]
 * @create date 2019-04-15 14:27:58
 * @modify date 2019-04-15 14:27:58
 * @desc [description]
 */

package main

import (
	"fmt"
	"os"
	"github.com/lmlala/Devent/backend/app"
	"github.com/lmlala/Devent/backend/conf"
	
)

// type Event struct{
// 	StartTime time.Time	`"json:start_time"`
// 	EndTime	time.Time	`"json:end_time"`
// 	EventType	string	`"json:event_type"`
// 	EventTag	string	`"json:event_tag"`
// 	Description	string	`"json:description"`
// }



func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Miss config file path")
		os.Exit(9)
	}
	
	err := config.InitConfig(args[1])

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Read Config Error !!!")
		os.Exit(9)
	}
    app.Run()
}