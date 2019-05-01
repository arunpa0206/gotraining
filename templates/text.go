package main

import (
    "text/template"
    "os"
)


type Joke struct {
    Who string
    Punchline string
}

func main() {
    t := template.New("Knock Knock Joke")
    text := `Knock Knock\nWho's there?
             {{.Who}}
             {{.Who}} who?
             {{.Punchline}}
            `
    t.Parse(text)

    jokes := []Joke{
        {"Etch", "Bless you!"},
        {"Cow goes", "No, cow goes moo!"},
    }

    for _, joke := range jokes {
        t.Execute(os.Stdout, joke)
    }
}
