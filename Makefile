TAG = $(shell git rev-parse --short HEAD)
PROJECT = forloop-162523
all:
	docker build -t us.grc.io/$(PROJECT)/guestbook:$(TAG)
	gcloud docker -- push us.grc.io/$(PROJECT)/guestbook:$(TAG)

