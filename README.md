# Imguessr
The goal of this project is to create a web based game where one would guess the content of a pixelated image sent by another player. The image will get less pixelated over time in order to make it easier to guess.

This project is made in go, following good practices in order for me to become more familiar with the language.

# .env
The `.env` file is used to store environment variables. It is not versioned and should be created by the user. The following variables are required:
- `JWT_SECRET`: The secret used to sign the JWT tokens

# Tests
Tests are run using the following command in the `test` folder:
```bash
go test
```

Go needs to be installed on the machine and the dependencies need to be installed using `go get`.
