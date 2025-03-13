# Kramer AI-Assistant

Kramer is that old friend that we can ask the most ridiculous things and they always have a (equally as ridiculous) answer.
So this is basically an AI wrapper, there is no magic in this.
There are probably a million of those tools that do the exact same thing, but I enjoyed writing and intend on keep making it better and more usefull as the time goes by.

## Supported models

- GPT-4

## How to run it

We have a requirement on sqlite, so if you don't have it in your system, please go to [sqlite](https://www.sqlite.org/download.html) and download it.
We also require the Go programming language in order to build the project, so please, if you don't have it go and install.

Also, you need to have your API key in your env, so you can run the following command:
```bash
$ export OPENAI_API_KEY=<your key>
```

To run it you can:
```bash
$ go build cmd/kramer/main.go
$ ./main <context>
```

Context is the name of the context you want to use, if none is given, kramer will start a new one for every time you restart the program.

If you don't know what is your key or how to get one, go [here](https://medium.com/@lorenzozar/how-to-get-your-own-openai-api-key-f4d44e60c327).

## FAQ

- Do you plan on supporting more models?
A: Yea maybe some day, since I believe most of them are easy to implement is just a matter of testing, so yea, I'll add support to some models.

- I found this bug.
A: Please open an issue so I can track it, but I don't promise any features.

- What do you think about feature X?
A: If it is cool I might implement it.
