
## 🛠️ Tools Included

### 1. **Project Setup Tool (`papi`)**

The `papi` tool is a command-line utility to help you quickly set up and run Python API projects. It can create project directories with a default structure, run your Python API, and even test the endpoints. It supports Flask-based API templates for quick setup.

#### Commands:
- **`new [api-app-name]`**: Create a new Python API project with the default directory structure, example files, and a configuration file.
- **`run`**: Run the Python API application.
- **`test`**: Test all available API endpoints (listed in `papi.paths.json`), and print the results.

#### Example usage:
```bash
# Create a new project
papi new my-api-app

# Run the Python API
papi run

# Test all API endpoints
papi test
```

### 2. **Flask API Template**

An example Flask application is included that you can use as a base for your own API projects. This template provides a simple `/api/greet` endpoint that returns a greeting message in JSON format.

### 3. **Endpoint Testing**

The repository also includes tools and examples for testing your API endpoints. The tests are automatically read from `papi.paths.json`, which lists all available endpoints to be tested.

---

## 📂 Project Structure

The project follows a simple structure, designed for easy use and extendability:

```bash
my-api-app/
├── templates/
│   └── api-example/
│       └── app.py             # Example Python API code (Flask)
├── papi.config.json           # Configuration file for the project
├── papi.paths.json            # List of all API endpoints for testing
├── papi.go                    # Go CLI tool to manage the project
└── README.md                  # This file
```

- **`templates/api-example/app.py`**: A simple Python API using Flask.
- **`papi.config.json`**: Configuration file that stores metadata about the project.
- **`papi.paths.json`**: Contains a list of all endpoints that can be tested with the `test` command.

---

## 🚀 Getting Started

### Prerequisites

To use this repository, you need the following:
- **Go** (for the `papi` command-line tool)
- **Python 3.x** with **Flask** installed
- **Git** (to clone the repository)

#### Install Flask (if not already installed):
```bash
python3 -m pip install flask
```

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/python-api-helper-tools.git
   cd python-api-helper-tools
   ```

2. Run the Go tool for managing Python API projects:
   ```bash
   go run papi.go
   ```

---

## 🔧 How to Use

### 1. Create a New Python API Project

To create a new Python API project with a basic Flask template, use the `new` command:
```bash
papi new my-api-app
```

This will generate the following structure:

```bash
my-api-app/
├── templates/
│   └── api-example/
│       └── app.py             # Flask API code
├── papi.config.json           # Configuration file
└── papi.paths.json            # API endpoints
```

### 2. Run the API

To run your API, use the `run` command:
```bash
papi run
```

This will start the Flask app, and your API will be available at `http://127.0.0.1:5000`.

### 3. Test the API Endpoints

The `test` command can be used to test all available API endpoints:
```bash
papi test
```

It will automatically fetch the API endpoints listed in `papi.paths.json` and perform GET requests to test them.

---

## 📝 Configuration Files

### `papi.config.json`
This file stores metadata about the project. Example content:
```bash
{
  "api": "flask",
  "version": "1.0.0"
}
```

### `papi.paths.json`
This file contains a list of all API endpoints to be tested. Example content:
```bash
{
  "endpoints": [
    {
      "method": "GET",
      "url": "/api/greet"
    }
  ]
}
```

---

## 💬 Contributing

Contributions are welcome! If you'd like to improve this project or add more tools, please follow these steps:

1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Make your changes
4. Commit your changes (`git commit -am 'Add new feature'`)
5. Push to the branch (`git push origin feature-branch`)
6. Create a new Pull Request

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🛠️ Troubleshooting

If you encounter any issues, please check the following:

- Make sure **Python 3.x** is installed and available in your `PATH`.
- Ensure that **Flask** is installed using `pip install flask`.
- If you're using `./papi.exe` and it's throwing an error, check for errors in the Go tool output and make sure Go is properly installed.

Feel free to open an issue if you need further assistance!

---

## 📣 Follow Me

- GitHub: [hmZa-Sfyn](https://github.com/hmZa-Sfyn)
- Documentation: [papi-link](https://github.com/SideProjects-IDK/papi-python)

---

Thank you for checking out this project! ✨
