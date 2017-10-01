assets = jquery.min.js materialize.min.css materialize.min.js

all: $(assets)

jquery.min.js:
	wget https://cdnjs.cloudflare.com/ajax/libs/jquery/3.2.1/jquery.min.js
	echo -e '\n$( document ).ready(function(){$(".button-collapse").sideNav();})' >> jquery.min.js

materialize.min.css:
	wget https://cdnjs.cloudflare.com/ajax/libs/materialize/0.100.2/css/materialize.min.css

materialize.min.js:
	wget https://cdnjs.cloudflare.com/ajax/libs/materialize/0.100.2/js/materialize.min.js

clean:
	rm -f $(assets)
