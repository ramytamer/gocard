# Welcome to GoCard!

[Powered by Buffalo](http://gobuffalo.io)

## Database Setup

The first thing you need to do is open up the "database.yml" file and 
check the env vars name to configure the correct host, port, username, password... that are appropriate for your environment.

You can find a `docker-compose.yml` that can start a local database for you.
In the compose file you can configure the username, password, exposed port for the database, these values will be referenced in the `database.yml` file as env vars.

To start the services in the docker compose file run:

	$ docker-compose up

### Create Your Databases

Ok, so you've edited the "database.yml" file and started your database, now Buffalo can create the databases in that file for you:

	$ buffalo pop create -a

Note that the option `-a` will create all database defined in your `database.yml` file.
## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

you can change the port by running

	$ PORT=4000 buffalo dev

**Congratulations!** You now have your Buffalo application up and running.

## What Next?

We recommend you heading over to [http://gobuffalo.io](http://gobuffalo.io) and reviewing all of the great documentation there.

Good luck!

[Powered by Buffalo](http://gobuffalo.io)
