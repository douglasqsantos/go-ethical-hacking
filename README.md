# Go Ethical Hacking Tools

The project is packaged with Docker support.

Download the source code and execute the docker-compose.yml

```bash
cd src/app/
docker compose -f docker-compose.yml up -d --build --remove-orphans
```

After start the containers add the default user
```bash
docker exec -it app_api_1 go run commands/populateUsers.go
```

If you want to test the app use the Collection inside the Postman directory.

You can change the .env file to make it fits your needs.

# Aim

Just to write some tools to with Golang to help in the CTF's ;)