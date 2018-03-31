build:
	docker build -t user-cli .

run:
	docker run -e MICRO_REGISTRY=mdns user-cli command \
		--name="Bob Loblaw" \
		--email="bob.loblaw@craigpeters.me" \
		--password="password@123" \
		--company="Bob Loblaw Law Blog"