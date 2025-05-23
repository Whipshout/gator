# Configuration file:
Configuration file in ~/.gatorconfig.json<br>

This format. It must have db_url from the beginning, you setup the current user using commands:<br>

`{"db_url":"URLTOCONNECTTOYOURPOSTGRES","current_user_name":"WHATEVER"}`

# Launch docker
`docker compose up -d`

# To run migrations:
`cd sql/schema`<br>

`goose postgres postgres://postgres:postgres@localhost:5432/gator up`

# To generate sqlc database code:
`sqlc generate`

# Application commmands:
`gator reset` - Reset database for testing purposes<br>

`gator register <name>` - Register user<br>

`gator login <name>` - Log in user<br>

`gator addfeed <name> <url>` - Add new feed for current user<br>

`gator follow <url>` - Follow new feed for current user<br>

`gator unfollow <url>` - Unfollow feed for current user<br>

`gator following` - List current user followed feeds<br>

`gator agg <time>` - Time in format 1s, 1h, 1ms. Fetch current user feeds and save them to database<br>

`gator browse <limit>` - List posts saved for current user<br>
