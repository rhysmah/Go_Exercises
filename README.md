# Go Exercises

## Exercise 1: CLI Quiz Game

### Overview

This exercise is based on Exercise #1: Quiz Game on the [gophercises GitHub](https://github.com/gophercises/quiz).

This is a simple quiz game that reads questions and answers from a CSV file, then, in the terminal, asks the user to answer them. The game keeps track of the number of correct answers and prints the final score at the end.

Includes a flag to set the time limit for the quiz; the default time is 30 seconds. To set a custom time, use the `--time` flag following by an integer (representing seconds).

### Running the Program

1. Navigate to the `Quiz_Game` directory and compile the code

```bash
cd Quiz_Game
make
```

2. Run the program; select quiz time limit (optional). The default is 30 seconds.

```bash
./Quiz_Game --time=15
```

## Exercise 2: URL Shortener

This exercise is based on Exercise #2: URL Shortener on the [gophercises GitHub](https://github.com/gophercises/urlshort).

### Overview

URL shorteners are typically used to turn long URLs into shorter, more human-readable URLs (for promotional purposes, as an example). When a user clicks on the shortened URL or types it in, the shortener service checks if the URL is valid and redirects the user to the original, "real" URL.

This exercise simulates a URL shortener service that reads an incoming URL request, check that exists within a map, then serves the corresponding URL. If the URL is not found, the service returns a "fallback" that redirects the user to a 404 error page.

## Exercise 3: Choose Your Own Adventure

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

## Exercise 4: HTML Link Parser

This exercise is based on Exercise #4: HTML Link Parser on the [gophercises GitHub](https://github.com/gophercises/link)

### Overview

This link parser reads an HTML file and extracts all the links within it. The links are then printed to the terminal.

### Running the Program

1. Navigate to the `HTML_Link_Parser` directory and compile the code

```bash
make
```

2. Run the program

```bash
./HTML_Link_Parser
```