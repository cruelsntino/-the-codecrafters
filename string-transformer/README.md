# Operation Gopher Protocol: String Transformer

### Status: MISSION CRITICAL
This module is a specialized text-processing unit designed to repair and reformat corrupted intelligence reports intercepted by **SENTINEL**. It provides six core transformations to ensure data is readable by CodeCrafters analysts and machine-ready for HQ.

---

## Features & Commands

The program runs in an interactive loop. Type a **command** followed by your **text**.

| Command | Action | Example Input | Example Output |
| :--- | :--- | :--- | :--- |
| `upper` | Converts text to ALL CAPS | `sentinel online` | `SENTINEL ONLINE` |
| `lower` | Converts text to lowercase | `ALERT LEVEL 5` | `alert level 5` |
| `cap` | Capitalizes every word | `director adaeze` | `Director Adaeze` |
| `title` | Title Case (ignores small words) | `the fall of grid` | `The Fall of the Grid` |
| `snake` | lowercase_with_underscores | `Alert! Level 5.` | `alert_level_5` |
| `reverse` | Reverses letters in each word | `Lagos Nigeria` | `sogaL airegiN` |
| `help` | Displays supported commands | `help` | (List of commands) |
| `exit` | Safely shuts down the system | `exit` | `Shutting down...` |

---

##  Snake Case Specification
To ensure coordinates and agent IDs are machine-readable, the `snake` command follows these strict rules:
1. **Normalization**: All characters are converted to lowercase.
2. **Underscores**: Internal spaces are replaced with a single underscore (`_`).
3. **Sanitization**: Non-alphanumeric characters (punctuation like `!`, `@`, `#`) are removed to prevent data corruption.

---

##  How to Run

1. **Prerequisites**: Ensure Go (Golang) is installed on your system.
2. **Navigate** to the project directory:
   ```bash
   cd string-transformer