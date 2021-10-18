help:
	@echo use command deploy to deploy app to gcloud

deploy: static
	gcloud app deploy

static:
	gsutil -m rsync -r ./css gs://leros-capital.appspot.com/css
	gsutil -m rsync -r ./js gs://leros-capital.appspot.com/js
	gsutil -m rsync -r ./img gs://leros-capital.appspot.com/img
