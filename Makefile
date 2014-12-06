all:

clear:
	clear

serve:
	goop exec go run main.go

test: clean
	goop exec go test -v ./spec

clean:
	rm spec/phony.db || :
