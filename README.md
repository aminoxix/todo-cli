## todo cli app

a simple command-line interface (cli) for managing todos locally. this application allows you to **create**, **view**, **update**, and **delete** todos, saving your data to a local json file for persistence.

---

### features

- **create a todo**: add new tasks with ease.
- **view all todos**: list all saved todos.
- **update a todo**: modify existing tasks by their id.
- **delete a todo**: remove tasks by their id.

---

### prerequisites

- go installed (1.18 or later).

---

### installation

1. clone this repository:
   ```bash
   git clone https://github.com/aminoxix/todo-cli.git
   ```
2. navigate to the project directory:
   ```bash
   cd todo-cli
   ```
3. run the application:
   ```bash
   go run main.go
   ```

---

### usage

once the program is running, you'll be prompted with the following options:

```plaintext
what operation do you want to perform?

c: Insert a todo
r: View all todos
u: Update a todo
d: Delete a todo
```

1. enter the desired operation (`c`, `r`, `u`, or `d`).
2. follow the on-screen prompts to complete the action.

---

### todo persistence

- all todos are saved to a json file at `./shared/data/todos.json`.
- this ensures that your data persists between sessions.

---

_</3> by [aminos](https://aminoxix.vercel.app)._
