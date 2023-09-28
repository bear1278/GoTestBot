# GoTestBot
This is telegram bot, that you can use as task manager.
The bot offers functions such as add, delete, mark as done and show all tasks. 
MySQL is used as data storage. 
The bot uses a docker container for deployment.
Before run application add file .env with token to your bot and data source name.
To run the application for the first time, use:
```
docker compose up --build
```
Then use:
```
docker compose up
```
# Stack
- Go 1.2
- MySQL
- Dccker
