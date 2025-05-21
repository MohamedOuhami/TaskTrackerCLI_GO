# 📝 TaskTracker CLI - Built with Go

**TaskTracker CLI** is a simple command-line tool developed in Go to help users manage and keep track of their tasks locally. It supports adding, updating, and tracking tasks with different statuses such as **done**, **in progress**, and **not done** as a challenge present in the Roadmap.sh GO roadmap "https://roadmap.sh/projects/task-tracker".

---

## 🚀 Features

- ✅ Add new tasks
- ✏️ Update existing tasks
- ❌ Delete tasks
- 🔄 Change task status:
  - Mark as **In Progress**
  - Mark as **Not Done**
  - Mark as **Done**
- 📋 List tasks:
  - All tasks
  - In-progress tasks
  - Not done tasks
  - Done tasks

---

## 🛠️ Installation

### Option 1: Using the pre-built executable

If you're on Windows, you can use the included `TaskTrackerCLI_GO.exe` directly from the command line.

### Option 2: Build from source (requires Go installed)

Clone the repository and run:

```bash
go install
```

This will build the executable and place it in your `$GOPATH/bin`.

---

## 🧭 Usage

After building or downloading the executable, run it from the command line using:

```bash
TaskTrackerCLI_GO.exe [command]
```

---

### 📚 Available Commands

#### Show all available commands

```bash
TaskTrackerCLI_GO.exe help
```

#### Add a new task

```bash
TaskTrackerCLI_GO.exe addTask "Description of your task"
```

#### Update a task by ID

```bash
TaskTrackerCLI_GO.exe updateTask [TaskID] "Updated description"
```

#### Delete a task by ID

```bash
TaskTrackerCLI_GO.exe deleteTask [TaskID]
```

---

### 📋 Listing Tasks

#### List all tasks

```bash
TaskTrackerCLI_GO.exe listAll
```

#### List tasks in progress

```bash
TaskTrackerCLI_GO.exe listInProgress
```

#### List tasks marked as done

```bash
TaskTrackerCLI_GO.exe listDone
```

#### List tasks marked as not done

```bash
TaskTrackerCLI_GO.exe listNotDone
```

---

### 🔄 Change Task Status

#### Mark a task as done

```bash
TaskTrackerCLI_GO.exe markAsDone [TaskID]
```

#### Mark a task as not done

```bash
TaskTrackerCLI_GO.exe markAsNotDone [TaskID]
```

#### Mark a task as in progress

```bash
TaskTrackerCLI_GO.exe markAsInProgress [TaskID]
```

---

## 📂 File Storage

Tasks are stored locally in a `tasklist.json` file in the project root. The file is automatically updated as you modify tasks via the CLI.

---

## 👨‍💻 Author

Created by **Mohamed Ouhami**  
Feel free to connect or contribute!
