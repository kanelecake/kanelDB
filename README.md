## KanelDB

> The project is under development and does not give any guarantees.
> If you want to help, then create a pull request and I will definitely consider them!

```
.-. .-')    ('-.         .-') _   ('-.           _ .-') _ .-. .-')   
\  ( OO )  ( OO ).-.    ( OO ) )_(  OO)         ( (  OO) )\  ( OO )  
,--. ,--.  / . --. /,--./ ,--,'(,------.,--.     \     .'_ ;-----.\  
|  .'   /  | \-.  \ |   \ |  |\ |  .---'|  |.-') ,`'--..._)| .-.  |  
|      /,.-'-'  |  ||    \|  | )|  |    |  | OO )|  |  \  '| '-' /_) 
|     ' _)\| |_.'  ||  .     |/(|  '--. |  |`-' ||  |   ' || .-. `.  
|  .   \   |  .-.  ||  |\    |  |  .--'(|  '---.'|  |   / :| |  \  | 
|  |\   \  |  | |  ||  | \   |  |  `---.|      | |  '--'  /| '--'  / 
`--' '--'  `--' `--'`--'  `--'  `------'`------' `-------' `------'  
```

kanelDB is a lightweight in-memory database 
great for testing and storing any values 
in memory. Also, this database is volatile, 
which means that when you turn off the 
computer, all data is erased.

### Quickstart

```go
package main

import (
	"context"
	"fmt"
	kaneldb "github.com/kanelecake/kanelDB"
	"time"
)

var ctx = context.Background()

func main() {
	// create a database instance
	db := kaneldb.NewClient()

	// set to database
	ttl := 3 * time.Minute
	err := db.Set(ctx, "hello_key", "hello_world", &ttl).Err()
	if err != nil {
		panic(err)
	}

	// get a value
	result, getErr := db.Get(ctx, "hello_key").Result()
	if getErr != nil {
		panic(err)
	}

	fmt.Printf("%s", result)

	// Output: hello_world
}
```

### The main goal at the moment:

* Partial compatibility with Redis for correct debugging