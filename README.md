# charming-todo
Project created in order to learn Go by creating a Todo list from scratch using [Charm.sh](https://github.com/charmbracelet/charm).

## Description
This small Todo app is my vision of a TUI/CLI tool that aims at fulfulling the following needs:
- it should be **fast** and **performant**,
- implement **keyboard shortcuts** to be **efficient**,
- while still being as **beautiful** as a backend dev can do.

The app allows you to create new lists, indent todo items, edit them, and much more.

https://user-images.githubusercontent.com/56871713/173691950-1cdefec2-8b29-4822-86b2-af6d80e83953.mp4



## State
The Todo app is mostly working, but is **missing some core mechanics** that would need to be implemented to make it useable; please refer to the "improvements" section.
Unfortunately, I've lost interest in this project. I will reflect more in details about this in later sections.

## Project goals
The initial goal of this project was to dip my toes in Go, which looked like an interesting programming language to me.
This project was a perfect introduction to Go + TUIs with Charm, which I would then use in other personal/work projects. 

## Learnings & Mistakes
Go is pretty awesome. The syntax is clear and easy to learn. The formatter and testing is very interesting.
The Elm architecture (implemented by Charm components) is lovely, I hope to encounter it in the future again.
  
➡️ I'm proud of having learnt Go this fast and super happy to have written some code with it.
  
❎ The main mistake I've made in this project is that I didn't spec out the project before starting its development. As in multiple previous projects, I have been too ambitious. There are so many things I would like to implement, but I lost the motivation to do it.

❎ Also, I made the mistake of wanting to write this from scratch whilst using the least amount of Charm premade components. This is a mistake, as I ended up recreating the wheel (ex. with lists or inputs) and wasting time on things where I shouldn't have.
  
## Improvements
1. Saving/loading file(s) containing the todo lists,
2. More keyboard shortcuts (CTRL+V, HOME, END, CTRL+SHIFT+left/right, ...),
3. Dialog that would allow opening previously closed Todo lists,
5. Tests
  
