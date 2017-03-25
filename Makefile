TAG = $(shell git rev-parse --short HEAD)
PROJECT = forloop-162523

appbuild:
	GOOS=linux go build -o app

docker:
	docker build -t us.gcr.io/$(PROJECT)/guestbook:$(TAG) .

gcloud:
	gcloud docker -- push us.gcr.io/$(PROJECT)/guestbook:$(TAG)

all: appbuild docker gcloud
