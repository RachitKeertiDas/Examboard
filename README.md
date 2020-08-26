## EXAM-BOARD

This is a project to help professors with the scheduling of exams. They can see potential confilcts with their exam-dates and schedule exams accordingly.

## Local Installation and Usage

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.  
These assume that you have Golang, npm, and PostrgeSQL installed on your locl system.

1. ` git clone ` the repository.
2. ` cd /path/to/the/cloned/directory/client ` and run `npm install`
3. ` cd /path/to/the/cloned/directory/server ` and
	1. In files main.go and db.go update your postrges Credentials in the const variable that has been declared.
	2. Run ` go build main.go ` and ` go build db.go `
4.  In the ` server directory` , first run `./db ` and then `./main` . Simultaneously, in the `client` directory, run ` npm run start `. 
5. 	Open your browser and navigate to `localhost:3000` 

## Dependencies
	
The Go backend has some third-party dependencies.
` go get ` them. 

Note : Some dependencies require `GOMODULE` support
If some dependencies fail to install via go get, try prefixing ` GO111MODULE=on ` before the command.   

This project utilises a [Create-React-App](https://create-react-app.dev/) Frontend and a [Golang](http://golang.org) backend with a [PostgreSQL DB](https://www.postgresql.org/).


## LICENSE

MIT

## Built with:

### Frontend:

* [Chakra UI](http://chakra-ui.com)

### Backend:

* [go-chi](https://github.com/go-chi/chi)
