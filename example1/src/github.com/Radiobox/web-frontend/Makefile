HEROKU_REMOTE_ERR_MESSAGE = No heroku remote found.  Please add the heroku git URL as a git remote, and name it heroku.
HEROKU_DEPLOY_MESSAGE = Deploying to heroku for the first time requires an application token, manually generated on github.com.  If this is your first time deploying on this computer, make sure you have one handy.

DB_REQUIREMENTS_MESSAGE = Our application requires PostgreSQL 9.3 and PostGIS - if you are missing either of these, chances are that something will not work
RADIOBOX_DB := $(findstring radiobox,$(shell psql -l -U postgres))
GITMASTERHASH := $(shell git log -n 1 master --pretty=format:"%H")

REMOTES = $(shell git remote)

.PHONY: nothing build-public-dist chek-heroku-remote heroku-deploy db-requirements devdb-requirements devdb-full devdb-blank devdb-data devdb-wipe devdb-reset

nothing:
	$(info What do you want to do, please be more specific)

build-dist:
	$(info Exporting current master commit hash)
	$(info $(shell echo "Version:";tput setaf 1; echo $(GITMASTERHASH); tput sgr 0))
	grunt --filehash $(GITMASTERHASH) --gruntfile ./public/vendor/Gruntfile.js --force
	git add --all public/dist
	git commit -m "Dist build $(GITMASTERHASH)"

check-heroku-remote:
ifeq (,$(findstring heroku,$(REMOTES)))
	$(info Adding heroku remote...)
	git remote add heroku git@heroku.com:radiobox-api-frontend-go.git
endif

heroku-deploy: check-heroku-remote
	$(MAKE) build-dist
	@export $(shell heroku config --shell) && ./deploy.sh

db-requirements:
	$(info $(DB_REQUIREMENTS_MESSAGE))

devdb-requirements: db-requirements
ifeq (,$(RADIOBOX_DB))
	$(info Creating radiobox user and database)
	-@createuser --username postgres radiobox >/dev/null 2>&1
	@createdb --username postgres --owner radiobox radiobox
endif

devdb-dump:
	@pg_dump --username radiobox --dbname radiobox --format c --file ./database/development_full.sql.c
	@pg_dump --username radiobox --dbname radiobox --format c --file ./database/development_schema.sql.c --schema-only
	@pg_dump --username radiobox --dbname radiobox --format c --file ./database/development_data.sql.c --data-only

devdb-updates: devdb-requirements
	@psql --username radiobox --dbname radiobox < ./database/recent_updates.sql 2>/dev/null
	$(info Updates Finished)

devdb-update: devdb-updates

devdb-full: devdb-requirements
	@pg_restore --format c --username postgres --dbname radiobox ./database/development_full.sql.c

devdb-blank: devdb-requirements
	@pg_restore --format c --username postgres --dbname radiobox ./database/development_schema.sql.c

devdb-data: devdb-requirements
	@pg_restore --format c --username postgres --dbname radiobox ./database/development_data.sql.c

devdb-wipe:
	$(info WARNING: This will wipe your current development database.  Use ctrl+c to cancel.)
	@sleep 3
	@dropdb --username postgres radiobox

devdb-reset: devdb-wipe
	$(MAKE) devdb-full

dev-clean:
	@go clean -i

dev-updates: devdb-updates dev-clean
	@go get -u ./...

dev-update: dev-updates

dev-install: dev-clean
	@go install

dev-run: check-heroku-remote dev-install
	@export $(shell heroku config --shell) && web-frontend
