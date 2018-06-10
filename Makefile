all: site

godeps:
	go get -u gopkg.in/yaml.v2

out:
	mkdir out

assets:
	$(MAKE) -C js
	$(MAKE) -C css

out/js: out assets
	if [ ! -e out/js ]; then ln -s $(shell pwd)/js out/js; fi

out/css: out assets
	if [ ! -e out/css ]; then ln -s $(shell pwd)/css out/css; fi

site: out/js out/css godeps
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
