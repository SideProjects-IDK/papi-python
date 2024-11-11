package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// Project structure for the new Python API app
type Project struct {
	API     string            `json:"api"`
	Version string            `json:"version"`
	Paths   map[string]string `json:"paths"`
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "papi",
		Short: "papi is a tool to manage Python API projects",
	}

	// Add commands
	rootCmd.AddCommand(NewCmd())
	rootCmd.AddCommand(RunCmd())
	rootCmd.AddCommand(TestCmd())

	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// NewCmd creates a new project with a template Python API and configuration file
func NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "new [api-app-name]",
		Short: "Create a new Python API project with example files",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Create the new directory structure
			projectName := args[0]
			projectDir := filepath.Join(".", projectName)

			// Check if directory already exists
			if _, err := os.Stat(projectDir); err == nil {
				fmt.Println("Project already exists!")
				return
			}

			// Create directories
			err := os.MkdirAll(filepath.Join(projectDir, "templates", "api-example"), 0755)
			if err != nil {
				fmt.Println("Error creating directories:", err)
				return
			}

			// Create a papi.config.json file
			configContent := `{
				"api": "flask",
				"version": "1.0.0"
			}`
			err = ioutil.WriteFile(filepath.Join(projectDir, "papi.config.json"), []byte(configContent), 0644)
			if err != nil {
				fmt.Println("Error creating config file:", err)
				return
			}

			// Create a papi.paths.json file for API endpoints
			pathsContent := `{
				"/api/greet": "GET"
			}`
			err = ioutil.WriteFile(filepath.Join(projectDir, "papi.paths.json"), []byte(pathsContent), 0644)
			if err != nil {
				fmt.Println("Error creating paths file:", err)
				return
			}

			// Add example Python API code (Flask API)
			apiCode := `from flask import Flask
app = Flask(__name__)

@app.route('/api/greet', methods=['GET'])
def greet():
    return {"message": "Hello, welcome to my API!"}

if __name__ == '__main__':
    app.run(debug=True)
`
			err = ioutil.WriteFile(filepath.Join(projectDir, "templates", "api-example", "app.py"), []byte(apiCode), 0644)
			if err != nil {
				fmt.Println("Error creating Python app file:", err)
				return
			}

			fmt.Printf("Project '%s' created successfully!\n", projectName)
		},
	}
}

// RunCmd runs the Python API application using python3
func RunCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Run the Python API app",
		Run: func(cmd *cobra.Command, args []string) {
			// Correct way to use exec.Command without assigning it to cobra.Command
			command := exec.Command("python3", "templates/api-example/app.py")
			output, err := command.CombinedOutput() // Runs the command and captures output
			if err != nil {
				fmt.Println("Error running the app:", err)
				return
			}

			// Output the result of running the command
			fmt.Println("API is running:\n", string(output))
		},
	}
}

// TestCmd tests all API endpoints from the papi.paths.json file
func TestCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "test",
		Short: "Test all API endpoints from papi.paths.json",
		Run: func(cmd *cobra.Command, args []string) {
			// Read the paths from papi.paths.json
			data, err := ioutil.ReadFile("papi.paths.json")
			if err != nil {
				fmt.Println("Error reading paths file:", err)
				return
			}

			// Parse the JSON data
			var paths map[string]string
			err = json.Unmarshal(data, &paths)
			if err != nil {
				fmt.Println("Error unmarshalling JSON:", err)
				return
			}

			// Test each API endpoint
			for endpoint, method := range paths {
				fmt.Printf("Testing endpoint: %s (%s)\n", endpoint, method)
				resp, err := http.Get("http://127.0.0.1:5000" + endpoint)
				if err != nil {
					fmt.Println("Error testing API:", err)
					return
				}
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("Error reading response:", err)
					return
				}

				// Output the result of the test
				fmt.Printf("Test result for '%s' endpoint:\n%s\n", endpoint, body)
			}
		},
	}
}
