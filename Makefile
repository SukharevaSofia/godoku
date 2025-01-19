all: run

run: build
	./godoku 2> logfile

#tag libsqlite3 allows to dynamically link to existing
# system sqlite3 installation
build:
	go build -tags 'linux libsqlite3'

clean:
	rm -f godoku godoku.db logfile
