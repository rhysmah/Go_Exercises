## CLI Quiz Game

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