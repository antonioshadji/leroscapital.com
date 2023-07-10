help:
	@echo 'use <make deploy> to deploy app to gcloud'

deploy: static
	gcloud app deploy

# -m multiprocessing in parallel
# -c checksum to verify if similar files are indeed different files
# -d mirror source and destination
# -r recursive handling of directories
static:
	gsutil -m rsync -c -r -d ./css gs://leros-capital.appspot.com/css
	gsutil -m rsync -c -r -d ./js gs://leros-capital.appspot.com/js
	gsutil -m rsync -c -r -d ./img gs://leros-capital.appspot.com/img
	gsutil -m rsync -c -r -d ./robots.txt gs://leros-capital.appspot.com
