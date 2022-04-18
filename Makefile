help:
	@echo 'use <make deploy> to deploy app to gcloud'

deploy: static
	gcloud app deploy

static:
	gsutil -m rsync -r -d ./css gs://leros-capital.appspot.com/css
	gsutil -m rsync -r -d ./js gs://leros-capital.appspot.com/js
	gsutil -m rsync -r -d ./img gs://leros-capital.appspot.com/img
