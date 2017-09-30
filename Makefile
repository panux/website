all: site

out:
	mkdir out

assets:
	$(MAKE) -C js

out/js: out assets
	if [ ! -e out/js ]; then ln -s $(shell pwd)/js out/js; fi

site: out/js
	go run build.go

cleanassets:
	$(MAKE) -C js clean

cleanout:
	rm -rf out

clean: cleanassets cleanout

run:
	$(MAKE) site
	(cd out && quickserve)

srvupd: site
