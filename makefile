up: start-backend start-frontend

down: 
	cd backend && \
	$(MAKE) down

start-frontend:
	cd frontend && \
	$(MAKE) up

start-backend:
	cd backend && \
	$(MAKE) run-backend

publish-backend:
	cd backend && \
	$(MAKE) build publish

publish-frontend:
	cd frontend && \
	$(MAKE) build publish

publish: publish-backend publish-frontend