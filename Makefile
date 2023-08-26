help:
	@echo 'use <make deploy> to deploy app to gcloud'

deploy: static
	gcloud app deploy

static:
	gcloud storage cp ./robots.txt gs://leros-capital.appspot.com
	gcloud storage rsync --recursive --delete-unmatched-destination-objects js gs://leros-capital.appspot.com/js
	gcloud storage rsync --recursive --delete-unmatched-destination-objects css gs://leros-capital.appspot.com/css
	gcloud storage rsync --recursive --delete-unmatched-destination-objects img gs://leros-capital.appspot.com/img
