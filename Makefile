all:
	rm tmp tmptmp
	go test -benchmem -bench .
