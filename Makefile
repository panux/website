all: site

out:
	mkdir out

assets: out/js
	$(MAKE) -C js

out/js: out
	ln -s $(shell pwd)/js out/js

site: assets out
	go run build.go

cleanassets:
	$(MAKE) -C js clean

cleanout:
	rm -rf out

clean: cleanassets cleanout

run:
	$(MAKE) site
	$(MAKE) out/js
	(cd out && quickserve)
