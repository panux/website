assets = jquery.min.js materialize.min.js

all: $(assets)

jquery.min.js:
	wget https://cdnjs.cloudflare.com/ajax/libs/jquery/3.2.1/jquery.min.js

materialize.min.js:
	wget https://cdnjs.cloudflare.com/ajax/libs/materialize/0.100.2/js/materialize.min.js

clean:
	rm -f $(assets)
