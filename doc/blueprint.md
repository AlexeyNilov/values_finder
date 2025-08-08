
## Project Blueprint and Development Plan

This plan breaks down the project into small, iterative, and testable chunks, ensuring a stable and robust development process. We will build the application from the ground up, starting with core data structures and configuration, then implementing the main logic with a mock LLM client for safe testing, and finally integrating a real LLM client.

### Development Phases

1.  **Phase 1: Foundation (Steps 1-2)**

      * **Goal:** Set up the project structure, define core data types, and implement configuration file loading.
      * **Outcome:** A Go application that can parse and validate a `config.yml` file.

2.  **Phase 2: Session Logic & Mocking (Steps 3-4)**

      * **Goal:** Implement session data management, including timestamped file logging, and create a mock LLM client that returns predictable data.
      * **Outcome:** A robust session manager that logs all interactions and a mock LLM client that allows us to test application flow without API calls.

3.  **Phase 3: Core Application Flow (Steps 5-7)**

      * **Goal:** Build the main interactive loop of the CLI, including displaying questions, tracking progress, and recording user choices. This phase will use the mock LLM client.
      * **Outcome:** A functional CLI that takes the user through a full, simulated session and correctly logs the entire interaction history.

4.  **Phase 4: Final Integration & Polish (Steps 8-9)**

      * **Goal:** Implement the final value synthesis, add robust error handling, and wire in the real LLM client.
      * **Outcome:** A complete, production-ready CLI tool that fulfills all requirements from the specification.

-----

## Code Generation Prompts

Below are the sequential prompts for a code-generation LLM. Each prompt represents one step in our development plan and follows a test-driven development (TDD) approach.

### **Step 1: Project Scaffolding and Core Data Types**

This initial step creates the project structure and defines the essential data models that will be used throughout the application, as specified in the `spec.md` file.

```text
I am building a command-line tool in Go called "values-cli".

First, create the project structure. Initialize a new Go module named `values-cli`.

Next, create a file named `core/types.go`. In this file, define the Go data structures required by the developer specification. These are:

1.  `Choice`: Contains the question text, a slice of options, and the index of the selected option.
2.  `RankedValue`: Contains the value's name and a description.
3.  `RankedValues`: A slice of `RankedValue`.
4.  `SessionData`: Contains the session timestamp, a slice of all `Choice` structs, and the final `RankedValues`.

Ensure the struct fields are exported (start with a capital letter) and match the specification. Add comments explaining each struct.
```

### **Step 2: Configuration Loader**

This step implements a configuration loader that safely parses a YAML file. We will start by writing a test to ensure the loader works correctly with valid and invalid files, as per TDD best practices.

````text
In our "values-cli" project, I need to manage configuration from a YAML file.

**1. Create the test file:**
First, create a file named `config/config_test.go`. This test should:
- Define a test function, `TestLoadConfig`.
- Create a temporary `config.yml` file with the following content:
  ```yaml
  rounds: 20
  options_per_question: 2
````

  - Call a `LoadConfig` function with the path to this temporary file.
  - Assert that the returned config struct has `Rounds` equal to 20 and `OptionsPerQuestion` equal to 2.
  - Test for an error case where the file does not exist.

**2. Create the implementation file:**
Now, create a file named `config/config.go`. In this file:

  - Define a `Config` struct with fields `Rounds` and `OptionsPerQuestion`, using YAML tags for parsing.
  - Implement the `LoadConfig(path string) (*Config, error)` function.
  - This function should read the YAML file at the given path, unmarshal it into the `Config` struct, and return the struct. Handle potential file reading or parsing errors.

**3. Update main:**
Finally, create a `main.go` file at the root. For now, it should:

  - Call `config.LoadConfig("config.yml")`.
  - If there's an error, log it and exit.
  - If successful, print a message confirming the config was loaded, like "Configuration loaded: X rounds."

<!-- end list -->

````

### **Step 3: Session Manager and File Logger**

This step creates the session manager responsible for tracking state and logging the entire session to a timestamped file for later analysis. The test will ensure files are created with the correct name and content.

```text
Continuing with the "values-cli" project, let's implement the session manager and file logger.

**1. Create the test file:**
Create a file named `session/session_test.go`. The test should:
- Create a test function `TestSessionFlow`.
- Call a function `NewManager()` to initialize a new session manager. This should create a file with a timestamped name like `values_session_...txt`.
- Verify the file exists.
- Use the manager to log a sample `Choice` struct.
- Read the content of the log file and assert that it contains the question and the selected option from the sample choice.
- Clean up by deleting the created log file.

**2. Create the implementation file:**
Now, create `session/session.go`. It should:
- Import the `core` package for our types.
- Define a `Manager` struct that holds the session data (`core.SessionData`) and a file handle.
- Implement `NewManager()`: this function initializes a `SessionData` struct with the current timestamp, creates a uniquely named log file (`values_session_YYYY-MM-DD_HH-mm-ss.txt`), and returns a pointer to the `Manager`. It should handle file creation errors.
- Implement a method `(m *Manager) AddChoice(choice core.Choice)` that appends a choice to the session history and writes the question, options, and user's selection to the log file.
- Implement a method `(m *Manager) LogFinalValues(values core.RankedValues)` that writes the final ranked list to the log file.
- Implement a method `(m *Manager) GetHistory() []core.Choice` that returns the choice history.

The log file should be human-friendly.
````

### **Step 4: Mock LLM Client for Testing**

To build the application flow without relying on a live API, we'll create a mock LLM client. This is critical for fast, predictable, and offline testing. It will implement an interface that a real client will also use later.

```text
For our "values-cli" project, we need an LLM client. We will start with a mock client to test our application logic.

**1. Define the interface:**
First, create `llm/client.go`. In this file:
- Import the `core` package.
- Define an interface named `Client` with two methods:
  - `GenerateOptions(history []core.Choice) ([]string, error)`
  - `GenerateFinalValues(history []core.Choice) (core.RankedValues, error)`

**2. Create the mock implementation:**
In the same `llm/client.go` file, create a `MockClient` struct.
- Implement the `Client` interface on `*MockClient`.
- `GenerateOptions`: should ignore the history for now and return a hardcoded slice of two strings, e.g., `[]string{"Being creative", "Being disciplined"}` and a nil error.
- `GenerateFinalValues`: should ignore the history and return a hardcoded `core.RankedValues` slice with 2-3 sample values (e.g., Name: "Creativity", Description: "You value expressing yourself."). Return a nil error.
- Add a way to make the mock return an error for testing purposes. For example, add an exported boolean field `ShouldFail`. If this field is true, the methods should return a sample error.

**3. Create the test file:**
Create `llm/client_test.go` to unit test the `MockClient`.
- Test that `GenerateOptions` returns the expected hardcoded strings.
- Test that `GenerateFinalValues` returns the expected hardcoded `RankedValues`.
- Test the error-forcing mechanism: set `ShouldFail` to `true` and assert that both methods now return an error.
```

### **Step 5: Main Interactive Loop**

Now we'll wire our components together into the main application loop in `main.go`. This version will use the `MockClient` to simulate the interaction specified.

```text
Let's build the main interactive loop for "values-cli" using the components we've created.

Modify `main.go` to implement the session flow:
1.  **Initialization:**
    - Load the configuration from `config.yml` using your `config` package.
    - Initialize the session manager using `session.NewManager()`. Defer closing its resources.
    - Initialize the `llm.MockClient`.

2.  **Interaction Loop:**
    - Start a loop that runs for the number of rounds specified in the config.
    - Inside the loop:
        - Print the progress indicator, e.g., "Round 3 of 20".
        - Call the mock client's `GenerateOptions` method, passing the current history from the session manager.
        - Display the prompt "Which feels more important to you right now:" followed by the numbered options returned from the client.
        - Read the user's input (a simple `fmt.Scanln` is fine, no validation needed).
        - Create a `core.Choice` struct containing the question text, the options, and the user's selected index (remembering to convert from 1-based input to 0-based index).
        - Add the new `Choice` to the session manager using `AddChoice`.

3.  **Completion:**
    - After the loop, print a message like "Generating your values...".

This step does not yet generate the final values, just the loop itself. We are building incrementally. Run this to manually test the flow.
```

### **Step 6: Generating and Displaying Final Values**

This step completes the application flow by calling the LLM client after the loop, generating the final values, and displaying them to the user as required by the spec.

````text
We will now complete the application flow in `main.go` for "values-cli".

Modify `main.go` to add the final step after the interaction loop:
1.  **Get Final Values:**
    - After the loop finishes, call the `GenerateFinalValues` method from your `llm.MockClient` instance. Pass the entire choice history, which you can get from your session manager.
    - Check for errors.

2.  **Display Final Values:**
    - If successful, get the `RankedValues` slice.
    - Iterate through the slice and print the results to the console in a casual, human-friendly format. For each value, display its rank, name, and description. For example:
      ```
      Here's what seems most important to you:

      1. Creativity
      You value expressing yourself and thinking outside the box. It's a core part of who you are.

      2. Discipline
      You appreciate structure and the power of consistency to achieve long-term goals.
      ```

3.  **Log Final Values:**
    - Call the `LogFinalValues` method on your session manager instance to write the final ranked list to the session log file.

Run the application manually to confirm that the full flow works and the mock final values are displayed and logged correctly.
````

### **Step 7: Real LLM Client and Prompt Engineering**

Now we create a real LLM client. This involves implementing the same interface but making a real API call. The core of this step is crafting the prompts.

```text
It's time to integrate a real LLM into "values-cli". We will implement a real client that satisfies the `llm.Client` interface.

In a new file, `llm/gemini_client.go` (or `openai_client.go`), do the following:

1.  **Define the Client:**
    - Create a struct, e.g., `GeminiClient`, that holds the necessary API client configuration.
    - Implement the `llm.Client` interface on this struct.

2.  **Implement `GenerateOptions`:**
    - This method receives the `[]core.Choice` history.
    - You must craft a system prompt that tells the LLM its role. It should be instructed to generate two contrasting or deeper options based on the user's entire choice history to help them clarify their values. Instruct it to return ONLY the two options, separated by a newline.
    - Format the choice history into a simple text block to be included in the prompt.
    - Make the API call, parse the response (splitting the text by newline), and return the `[]string` options.

3.  **Implement `GenerateFinalValues`:**
    - This method also receives the full history.
    - Craft a different system prompt. This prompt instructs the LLM to act as a helpful coach. It should review the entire session history and synthesize a ranked list of exactly 5 personal values. For each value, it must provide a name and a 2-3 sentence, casual-toned description.
    - The prompt must specify the output format very strictly, for example:
      `1. Value Name: Description of the value...`
      `2. Value Name: Description of the value...`
      etc.
    - Make the API call, parse the structured response, and populate the `core.RankedValues` struct.

Handle API errors by returning them. This client should not have any CLI logic inside it.
```

### **Step 8: Error Handling and Final Wiring**

This final prompt cleans up `main.go` by adding the specified error handling and allowing the user to select the real LLM client.

```text
Let's finalize our "values-cli" tool by adding robust error handling and wiring up the real LLM client.

Modify `main.go`:

1.  **LLM Client Selection:**
    - Add a simple mechanism to switch between `llm.MockClient` and your new `llm.GeminiClient`. An environment variable or a command-line flag is a good choice. Initialize one or the other into your `llm.Client` interface variable.

2.  **Error Handling:**
    - Go to the parts of your code where you call `client.GenerateOptions` and `client.GenerateFinalValues`.
    - If the call returns an error, you must now handle it.
    - As per the specification, print the exact message: `Sorry, we’re having trouble connecting to the AI service right now. Please try again later.`
    - After printing the message, log the actual error to your session file for debugging and exit the program gracefully (`os.Exit(1)`).

3.  **Final Review:**
    - Read through your `main.go` one last time and ensure it matches the interaction flow from the specification perfectly:
        - Load config.
        - Init session.
        - Loop through rounds, getting choices from the LLM and recording user input.
        - After loop, get final values from LLM.
        - Display final values.
        - Log everything and exit.

Ensure there is no extra closing message after the final values are displayed. The program should simply print the list and terminate.
```