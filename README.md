# goxercise
A rustlings inspired Go framework to create series of exercises

## Toml config

The toml config will be as follows:

```toml
[[exercises]]
name = "intro1"             # The exercise name
filepath = "filepath"       # The filepath to the exercise itself
testpath = "filepath"       # The filepath for the tests
solutionpath = "optional"   # The path for a solution file (optional)
hint = """              
To finish this exercise, you need to …
These links might help you …"""
```
