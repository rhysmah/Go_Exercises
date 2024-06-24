## Choose Your Own Adventure

This exercise is based on Exercise #3: Choose Your Own Adventure on the [gophercises GitHub](https://github.com/gophercises/cyoa)

### Overview

"Choose Your Own Adventure" books are interactive stories where the reader makes a choice, then flips to another page based on this choice.

This exercises is a web application that recreates this experience. The story is read from a JSON file. At the end of every page, the user is presented options to choose from; selecting an option will take the user to the corresponding page.

### Running the Program

1. Open a terminal at the `backend` directory, compile the Go code, and run the program. This will start a backend server at `localhost:8080`.

```bash
make
./adventure
```

2. Open a terminal at the `frontend` directory and run the program. This will start a frontend server at `localhost:3000`.
    
```bash
cd adventure-game/
npm start
```