build-frontend:
	docker build --file=app/Dockerfile --rm=true -t training-cal-frontend .

build-backend:
	docker build --file ./Dockerfile -t training-cal-backend .

run: build-frontend build-backend
	docker-compose -f docker-compose.yml up

build-heroku:
	docker build --file ./Dockerfile -t registry.heroku.com/training-calendars/web .
	docker push registry.heroku.com/training-calendars/web
	docker build --file=app/Dockerfile --rm=true -t registry.heroku.com/training-cal/web .
	docker push registry.heroku.com/training-cal/web

	heroku container:release web -a training-calendars
	heroku container:release web -a training-cal
