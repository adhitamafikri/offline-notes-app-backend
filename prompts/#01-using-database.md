# Application Database Spec

## Required DBs and their purposes

1. Postgres:
   - used to manage the User data
   - used to manage user auth (registration, login, refresh token)
2. CouchDB/PouchDB
   - used to manage the Note data
   - used to manage the User Settings data
   - can be used by our frontend service to sync their data when they are online
