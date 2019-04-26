NAME=oauth2-confidential-client
GOPKGNAME=github.com/kg0r0/oauth2-confidential-client

build:
	cd $(GOPATH)/src/$(GOPKGNAME) && go build

clean:
	rm $(NAME)
