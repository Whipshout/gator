# Configuration file:
Configuration file in ~/.gatorconfig.json
This format. It must have db_url from the beginning, you setup the current user using commands:
`{"db_url":"URLTOCONNECTTOYOURPOSTGRES","current_user_name":"WHATEVER"}`

# Launch docker
`docker compose up -d`

# To run migrations:
`cd sql/schema`
`goose postgres postgres://postgres:postgres@localhost:5432/gator up`

# To generate sqlc database code:
`sqlc generate`

# Application commmands:
`gator reset` - Reset database for testing purposes
`gator register <name>` - Register user
`gator login <name>` - Log in user
`gator addfeed <name> <url>` - Add new feed for current user
`gator follow <url>` - Follow new feed for current user
`gator unfollow <url>` - Unfollow feed for current user
`gator following` - List current user followed feeds
`gator agg <time>` - Time in format 1s, 1h, 1ms. Fetch current user feeds and save them to database
`gator browse <limit>` - List posts saved for current user
