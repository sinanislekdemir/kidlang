# 🎨 KidLang - Programming Language for Kids

**KidLang** is a beginner-friendly programming language designed for children aged 7-12. It focuses on making programming fun, intuitive, and forgiving - perfect for young learners taking their first steps into the world of coding!

## 🌟 Why KidLang?

- **Kid-Friendly Syntax**: Uses simple, everyday words that children understand
- **Forgiving Nature**: Doesn't punish kids for minor mistakes like case sensitivity or missing quotes
- **Multilingual**: Available in English, Turkish, Finnish, and German
- **Visual IDE**: Beautiful Curses IDE with syntax highlighting and integrated terminal
- **Learn by Doing**: Comes with 15+ example programs to learn from

## ✨ Key Features

- **Simple Variables**: Store values in "boxes" (e.g., `box age = 10`)
- **Easy Input/Output**: Use `ask` to get user input, `print` to display (or just write text!)
- **Smart Type System**: Automatically handles numbers and text - kids don't worry about types
- **Flexible Comparisons**: `"1"` equals `1` - the language figures it out
- **Lists/Stacks**: Store collections with flexible indexing (e.g., `stack toys[1] = "robot"`)
- **File Operations**: Read and write files easily
- **Control Flow**: Simple if/goto for logic, loops for repetition
- **Comments**: Use `//` to add notes to your code

## 📷 Screenshots!

| | | |
|---|---|---|
| [<img src="https://github.com/user-attachments/assets/9485b04e-4201-4d70-a065-83815f00cf68" width="220"/>](https://github.com/user-attachments/assets/9485b04e-4201-4d70-a065-83815f00cf68) | [<img src="https://github.com/user-attachments/assets/c579146d-84d3-4549-8694-dbbf1b5cfee1" width="220"/>](https://github.com/user-attachments/assets/c579146d-84d3-4549-8694-dbbf1b5cfee1) | [<img src="https://github.com/user-attachments/assets/ebb7315f-d3d2-4941-9c70-4d50bad3f1fe" width="220"/>](https://github.com/user-attachments/assets/ebb7315f-d3d2-4941-9c70-4d50bad3f1fe) |
| [<img src="https://github.com/user-attachments/assets/4ea6ac08-6ba9-4828-a839-ba678581486d" width="220"/>](https://github.com/user-attachments/assets/4ea6ac08-6ba9-4828-a839-ba678581486d) | [<img src="https://github.com/user-attachments/assets/2d65d4da-7626-4f7d-9781-617155117e6b" width="220"/>](https://github.com/user-attachments/assets/2d65d4da-7626-4f7d-9781-617155117e6b) | [<img src="https://github.com/user-attachments/assets/8ad86c31-c006-4ec3-911e-78f0e2fbfe5e" width="220"/>](https://github.com/user-attachments/assets/8ad86c31-c006-4ec3-911e-78f0e2fbfe5e) |
| [<img src="https://github.com/user-attachments/assets/5081b1d1-1cb8-474a-bb1d-796ee648f662" width="220"/>](https://github.com/user-attachments/assets/5081b1d1-1cb8-474a-bb1d-796ee648f662) | [<img src="https://github.com/user-attachments/assets/5703a442-494b-4e46-b24f-21966296c700" width="220"/>](https://github.com/user-attachments/assets/5703a442-494b-4e46-b24f-21966296c700) | [<img src="https://github.com/user-attachments/assets/8e63bcc5-d89a-4b8a-8d25-a24eae320c0b" width="220"/>](https://github.com/user-attachments/assets/8e63bcc5-d89a-4b8a-8d25-a24eae320c0b) |




## 🚀 Quick Start

### Hello World

```kidlang
// Your first program!
print Hello, World!
print Welcome to KidLang!
```

### Interactive Program

```kidlang
// Ask for user's name
ask What is your name?
print Hello box answer
print Nice to meet you!
```

### Using Variables (Boxes)

```kidlang
// Boxes store your values
box age = 10
box next_year = box age + 1

print You are box age years old
print Next year you will be box next_year
```

### Simple Game Logic

```kidlang
ask Pick a number between 1 and 10:
box guess = answer
box secret = 7

if box guess = box secret then
goto winner
end
print Sorry, try again!
goto end

winner:
print You won! The number was box secret

end:
```

## 🛠️ Building from Source

### Engine (Command Line Interpreter)

The KidLang engine is written in Go.

**Prerequisites:**
- Go 1.16 or higher

**Build:**
```bash
cd engine
go build -o kidlang
```

**Run a program:**
```bash
./kidlang myprogram.kid
```

**Run tests:**
```bash
cd engine
go test
```

### UI (Visual IDE)

The KidLang IDE is built with Qt5/C++.

**Prerequisites:**
- Qt5 development libraries
- C++ compiler (g++)
- make

**On Ubuntu/Debian:**
```bash
sudo apt-get install qt5-default qtbase5-dev g++ make
```

**Build:**
```bash
cd ui
qmake
make
```

**Run:**
```bash
./kidlang-ide
```

## 📚 Learning Resources

### Tutorials (Step-by-Step Learning)
- **[TUTORIAL_BEGINNER.md](TUTORIAL_BEGINNER.md)** - Start here! Learn variables, loops, and basics
- **[TUTORIAL_ALGORITHMS.md](TUTORIAL_ALGORITHMS.md)** - Build real algorithms (sorting, searching, math)
- **[TUTORIAL_ADVANCED.md](TUTORIAL_ADVANCED.md)** - Advanced topics, file handling, complex projects
- **[TUTORIAL_PROJECTS.md](TUTORIAL_PROJECTS.md)** - Complete projects with step-by-step builds

### Quick References
- **[COMMAND_REFERENCE.md](COMMAND_REFERENCE.md)** - Complete command documentation
- **[CHEATSHEET_EN.md](CHEATSHEET_EN.md)** - Quick reference (English)
- **[CHEATSHEET_TR.md](CHEATSHEET_TR.md)** - Hızlı referans (Türkçe)
- **[CHEATSHEET_FI.md](CHEATSHEET_FI.md)** - Pikalinkit (Suomi)
- **[CHEATSHEET_DE.md](CHEATSHEET_DE.md)** - Schnellreferenz (Deutsch)

### Example Programs
- **[examples/](examples/)** - 20+ working programs to learn from

## 📖 Example Programs

The `examples/` directory contains programs for various skill levels:

1. **01_hello_world.kid** - Your first program
2. **02_simple_math.kid** - Basic calculations
3. **03_user_input.kid** - Getting user input
4. **04_guessing_game.kid** - Fun number guessing game
5. **05_calculator.kid** - Simple calculator
6. **06_countdown.kid** - Countdown timer with loops
7. **07_multiplication_table.kid** - Math practice
8. **08_temperature_converter.kid** - Unit conversion
9. **09_math_quiz.kid** - Random math questions
10. **10_todo_list.kid** - File-based to-do list
11. **11_story_generator.kid** - Creative storytelling
12. **12_simple_game.kid** - Text adventure game
13. **13_gradebook.kid** - Student grade manager
14. **14_contact_book.kid** - Contact management with stacks
15. **15_inventory.kid** - Game inventory system

## 🌍 Multilingual Support

KidLang supports multiple languages! Keywords are translated:

| English | Turkish | Finnish | German |
|---------|---------|---------|--------|
| box | kutu | laatikko | kiste |
| print | yaz | tulosta | schreib |
| ask | sor | kysy | frag |
| if | eğer | jos | wenn |
| goto | git | mene | geh |
| answer | cevap | vastaus | antwort |
| stack | liste | pino | liste |

## 🎯 Design Philosophy

KidLang is designed with kids in mind:

1. **Forgiving**: Case-insensitive, flexible whitespace, automatic type conversion
2. **Simple**: No complex syntax, no semicolons, no strict types
3. **Intuitive**: Uses real-world metaphors (boxes, stacks)
4. **Fun**: Encourages experimentation and creativity
5. **Educational**: Teaches core programming concepts without overwhelming complexity

## 🤝 Contributing

Contributions are welcome! Whether it's:
- New example programs
- Documentation improvements
- Bug fixes
- New language translations
- Feature suggestions

## 📄 License

See [LICENSE](LICENSE) file for details.

## 🎓 Educational Use

KidLang is perfect for:
- Elementary school programming classes
- After-school coding clubs
- Parents teaching kids at home
- Kids learning independently
- Summer coding camps

## 💡 What's Next?

After mastering KidLang, kids will be ready to move on to languages like Python, JavaScript, or Scratch, having already learned:
- Variables and data storage
- User input/output
- Conditional logic (if statements)
- Loops and repetition
- File operations
- Data structures (lists/stacks)
- Problem-solving and algorithmic thinking

---

**Made with ❤️ for young programmers everywhere!**
