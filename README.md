# Welcome to GoCard!

[Powered by Buffalo](http://gobuffalo.io)

## Building the project for local development

First thing we need to do is to create the `.env` file, you can simply copy the values from the example file as it's configured for docker based build.

    $ cp .env.example .env
    $ docker-compose build

## Database Setup

**You can skip this if you are working with the default docker setup**

The first thing you need to do is open up the "database.yml" file and 
check the env vars name to configure the correct host, port, username, password... that are appropriate for your environment.

You can find a `docker-compose.yml` that can start a local database for you.
In the compose file you can configure the username, password, exposed port for the database, these values will be referenced in the `database.yml` file as env vars.

### Create Your Databases

Ok, so you've edited the "database.yml" file and started your database, now Buffalo can create the databases in that file for you:

	$ docker-compose run --rm app buffalo pop create -a

Note that the option `-a` will create all database defined in your `database.yml` file.
In our case that would be the development and test databases.
## Starting the Application

Simply up the services defined in docker compose:

	$ docker-compose up

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a ""Welcome to GoCard!"" page.

you can change the port by changing the exposed port in the `docker-compose.yml` file.

**Congratulations!** You now have your GoCard application up and running.

## Running Tests

If you want to run the full test suit, just execute:

    $ docker-compose exec app buffalo test

You can also check for the coverage by running the following:
    
    $ docker-compose exec app buffalo test -coverprofile=c.out ./..

The current coverage report is as follows:

```
[POP] 2021/07/07 02:09:34 info - dropped database gocard_test
[POP] 2021/07/07 02:09:35 info - created database gocard_test
[POP] 2021/07/07 02:09:35 info - loaded schema for gocard_test
INFO[0000] go test -p 1 -tags development -coverprofile=c.out ./...
?   	gocard	[no test files]
ok  	gocard/actions	0.135s	coverage: 88.0% of statements
?   	gocard/grifts	[no test files]
ok  	gocard/models	0.050s	coverage: 66.7% of statements
```

## APIs

You can find attached in the repo a [Postman file](gocard.json) containing the endpoints and the required params

### Endpoints in API Blueprint format

# GET /status?unhealthy=false

+ Response 200 (application/json; charset=utf-8)

    + Body

            {"status":"ok"}
            


# POST /decks

+ Request (application/json; charset=utf-8)

    + Headers

            Accept: application/json

    + Body

            {
                "shuffled": true,
                "partial_cards": "AH,KC,QS"
            }

+ Response 201 (application/json; charset=utf-8)

    + Body

            {
                "deck_id": "85622fd0-6cf5-4167-a6a2-9a4fb4254b54",
                "remaining": 3,
                "shuffled": true
            }
            


# GET /decks/612eddf9-1000-4bb0-9582-bbcce1182120

+ Request (application/x-www-form-urlencoded; charset=utf-8)

    + Headers

            Accept: application/json



+ Response 200 (application/json; charset=utf-8)

    + Body

            {
                "cards": [
                    { "code": "6H", "suit": "2", "value": "6" },
                    { "code": "KC", "suit": "0", "value": "13" },
                    { "code": "JD", "suit": "1", "value": "11" },
                    { "code": "9H", "suit": "2", "value": "9" },
                    { "code": "KH", "suit": "2", "value": "13" },
                    { "code": "KD", "suit": "1", "value": "13" },
                    { "code": "KS", "suit": "3", "value": "13" },
                    { "code": "10H", "suit": "2", "value": "10" },
                    { "code": "3C", "suit": "0", "value": "3" },
                    { "code": "AH", "suit": "2", "value": "1" },
                    { "code": "AS", "suit": "3", "value": "1" },
                    { "code": "6D", "suit": "1", "value": "6" },
                    { "code": "5H", "suit": "2", "value": "5" },
                    { "code": "2D", "suit": "1", "value": "2" },
                    { "code": "9S", "suit": "3", "value": "9" },
                    { "code": "AC", "suit": "0", "value": "1" },
                    { "code": "5S", "suit": "3", "value": "5" },
                    { "code": "5C", "suit": "0", "value": "5" },
                    { "code": "3H", "suit": "2", "value": "3" },
                    { "code": "JS", "suit": "3", "value": "11" },
                    { "code": "4D", "suit": "1", "value": "4" },
                    { "code": "3S", "suit": "3", "value": "3" },
                    { "code": "8C", "suit": "0", "value": "8" },
                    { "code": "10D", "suit": "1", "value": "10" },
                    { "code": "7S", "suit": "3", "value": "7" },
                    { "code": "4C", "suit": "0", "value": "4" },
                    { "code": "10S", "suit": "3", "value": "10" },
                    { "code": "4H", "suit": "2", "value": "4" },
                    { "code": "4S", "suit": "3", "value": "4" },
                    { "code": "JC", "suit": "0", "value": "11" },
                    { "code": "5D", "suit": "1", "value": "5" },
                    { "code": "2H", "suit": "2", "value": "2" },
                    { "code": "8H", "suit": "2", "value": "8" },
                    { "code": "JH", "suit": "2", "value": "11" },
                    { "code": "10C", "suit": "0", "value": "10" },
                    { "code": "QH", "suit": "2", "value": "12" },
                    { "code": "7D", "suit": "1", "value": "7" },
                    { "code": "2S", "suit": "3", "value": "2" },
                    { "code": "7C", "suit": "0", "value": "7" },
                    { "code": "QD", "suit": "1", "value": "12" },
                    { "code": "6C", "suit": "0", "value": "6" },
                    { "code": "8D", "suit": "1", "value": "8" },
                    { "code": "9C", "suit": "0", "value": "9" },
                    { "code": "QC", "suit": "0", "value": "12" },
                    { "code": "8S", "suit": "3", "value": "8" },
                    { "code": "QS", "suit": "3", "value": "12" },
                    { "code": "7H", "suit": "2", "value": "7" },
                    { "code": "9D", "suit": "1", "value": "9" },
                    { "code": "6S", "suit": "3", "value": "6" },
                    { "code": "2C", "suit": "0", "value": "2" },
                    { "code": "3D", "suit": "1", "value": "3" }
                ],
                "deck_id": "612eddf9-1000-4bb0-9582-bbcce1182120",
                "remaining": 51,
                "shuffled": true
            }
            


# PUT /decks/612eddf9-1000-4bb0-9582-bbcce1182120/draw

+ Request (application/json; charset=utf-8)

    + Headers

            Accept: application/json

    + Body

            {
                "count": 1
            }

+ Response 200 (application/json; charset=utf-8)

    + Body

            {
                "cards": [
                    {
                    "code": "QS",
                    "suit": "3",
                    "value": "12"
                    }
                ]
            }
            


