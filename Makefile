all:

clear:
	clear

serve:
	goop exec go run main.go

test:
	goop exec go test -v ./spec

