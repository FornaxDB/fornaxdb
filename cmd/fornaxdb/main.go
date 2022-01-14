package main

import (
	"fmt"
	// "github.com/FornaxDB/fornaxdb/logger"
	// "github.com/FornaxDB/fornaxdb/errors"
	"github.com/FornaxDB/fornaxdb/schema"
)


func main() {
	// l := logger.New()
	// l.Trace("Hello, World!", map[string]interface{}{"foo": "bar"})
	// err := x()
	// if err != nil {
	// 	l.Fatal(err.Error(), map[string]interface{}{})
	// }
	s := `type Event {
		name: string!,
		website: string,
		cool(bla: string = BLA): boolean,
		participants: int?,
		sponsors: [string?]
	  }
	  
	  type Person {
		name: string
		twitter: string,
		skill_level: float,
	  }
	  
	  type Team {
		name: string!,
		website: string,
		country: string!,
		no_of_members: string | int,
		avg_elo: float!
	  }
	  
	  relation member_of {
		is_leader: boolean,
		__src: Person,
		__des: Team
	  }
	  
	  relation is_participating {
		__src: Team,
		__des: Event
	  }`
	  t := schema.Tokenize(s)
	  fmt.Println(t)
}

// func x() error {
// 	return errors.SchemaAlreadyExists.New("bla")
// }
