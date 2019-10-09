# Insert Project Name Here
## Insert Name Here
Insert project description here.

# User Stories
- [x] List
- [] Each
- [] User
- [] Story

# Instructions
Insert environment, build, and execution documentation here.


# Incase go deletes some of my imports
import (
	"fmt"

	//links to bank.go
	"github.com/190930-UTA-CW-Go/project-0/bank"

	"strings"

	"database/sql"

	//imports the db
	_ "github.com/lib/pq"
)

# Data Base Notes
cd db
docker build -t bankaccount .
docker run --name bankaccount -d -p 5432:5432 bankaccount
docker run --name bankaccount -it -p 5432:5432 bankaccount


after initial creation
docker exec -d bankaccount psql -U postgres
docker exec -it bankaccount psql -U postgres
docker start bankaccount
docker stop bankaccount
SELECT * FROM bankaccount;
\q

reset
cd into db
docker stop bankaccount
docker rm bankaccount
check with docker ps