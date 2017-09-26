all: site

out:
	mkdir out

assets:
	$(MAKE) -C js

site: assets out
	go run build.go

cleanassets:
	$(MAKE) -C js clean

cleanout:
	rm -rf out

clean: cleanassets cleanout
