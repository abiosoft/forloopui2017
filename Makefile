TAG = $(shell git rev-parse --short HEAD)
PROJECT = forloop-162523

build:
	GOOS=linux go build -o app

all: build
	docker build -t us.grc.io/$(PROJECT)/guestbook:$(TAG)
	gcloud docker -- push us.grc.io/$(PROJECT)/guestbook:$(TAG)

