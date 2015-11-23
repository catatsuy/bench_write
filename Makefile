all:
	rm -f tmp tmptmp
	go test -benchmem -bench .
