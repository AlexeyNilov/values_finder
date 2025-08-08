
# Project Checklist: Personal Values Discovery CLI

This checklist breaks down the development of the `values-cli` tool into manageable steps. Mark each item as complete as you progress.

### Phase 1: Foundation

-   [ ] **Step 1: Project Scaffolding and Core Data Types**
    -   [ ] Initialize Go module: `go mod init values-cli`
    -   [ ] Create `core/types.go` file.
    -   [ ] Define `Choice` struct with fields: `QuestionText`, `Options`, `Selected`.
    -   [ ] Define `RankedValue` struct with fields: `Name`, `Description`.
    -   [ ] Define `RankedValues` as `[]RankedValue`.
    -   [ ] Define `SessionData` struct with fields: `Timestamp`, `Choices`, `FinalRanking`.
    -   [ ] Ensure all necessary fields are exported and have comments.

-   [ ] **Step 2: Configuration Loader**
    -   [ ] Create `config/config_test.go`.
        -   [ ] Test loading a valid YAML file.
        -   [ ] Test handling of a non-existent file.
    -   [ ] Create `config/config.go`.
        -   [ ] Define `Config` struct with `Rounds` and `OptionsPerQuestion` fields.
        -   [ ] Add YAML tags to struct fields for parsing.
        -   [ ] Implement `LoadConfig(path string) (*Config, error)` function.
    -   [ ] Create initial `main.go`.
    -   [ ] Create `config.yml` with default values (`rounds: 20`, `options_per_question: 2`).
    -   [ ] Integrate `config.LoadConfig` call into `main.go`.

### Phase 2: Session Logic & Mocking

-   [ ] **Step 3: Session Manager and File Logger**
    -   [ ] Create `session/session_test.go`.
        -   [ ] Test that `NewManager` creates a file with the correct timestamped format (`values_session_YYYY-MM-DD_HH-mm-ss.txt`).
        -   [ ] Test that `AddChoice` writes the correct choice data to the log file.
        -   [ ] Test that `LogFinalValues` writes the final ranking to the log file.
    -   [ ] Create `session/session.go`.
        -   [ ] Define `Manager` struct to hold `SessionData` and file handle.
        -   [ ] Implement `NewManager()` to create a log file and initialize the session.
        -   [ ] Implement `(m *Manager) AddChoice(choice core.Choice)`.
        -   [ ] Implement `(m *Manager) LogFinalValues(values core.RankedValues)`.
        -   [ ] Implement `(m *Manager) GetHistory() []core.Choice`.
        -   [ ] Ensure log file output is human-friendly.

-   [ ] **Step 4: Mock LLM Client**
    -   [ ] Create `llm/client.go`.
        -   [ ] Define `Client` interface with `GenerateOptions` and `GenerateFinalValues` methods.
    -   [ ] Implement `MockClient` struct that satisfies the `Client` interface.
        -   [ ] `GenerateOptions` returns a hardcoded `[]string`.
        -   [ ] `GenerateFinalValues` returns hardcoded `core.RankedValues`.
        -   [ ] Add a mechanism to force methods to return an error for testing.
    -   [ ] Create `llm/client_test.go`.
        -   [ ] Test `GenerateOptions` success case.
        -   [ ] Test `GenerateFinalValues` success case.
        -   [ ] Test error-forcing mechanism for both methods.

### Phase 3: Core Application Flow

-   [ ] **Step 5: Main Interactive Loop**
    -   [ ] Modify `main.go`.
        -   [ ] Initialize `config`, `session.Manager`, and `llm.MockClient`.
        -   [ ] Implement the main `for` loop to run for the number of rounds from config.
        -   [ ] Display progress indicator in each round (e.g., "Round 3 of 20").
        -   [ ] Call mock `GenerateOptions` and display the prompt and options.
        -   [ ] Collect user input (no validation required).
        -   [ ] Create a `core.Choice` struct and log it with the session manager.

-   [ ] **Step 6: Generating and Displaying Final Values**
    -   [ ] Modify `main.go`.
        -   [ ] After the loop, call the mock `GenerateFinalValues` method with the full history.
        -   [ ] Create a function to display the `RankedValues` in a casual, numbered format.
        -   [ ] Ensure exactly 5 values are requested/handled.
        -   [ ] Ensure descriptions are 2-3 sentences long (for mock and real prompts).
        -   [ ] Print the formatted results to the console.
        -   [ ] Call `session.LogFinalValues` to write the final results to the log file.
        -   [ ] Ensure no extra closing message is printed after the results.

### Phase 4: Final Integration & Polish

-   [ ] **Step 7: Real LLM Client and Prompt Engineering**
    -   [ ] Create `llm/gemini_client.go` (or other provider).
    -   [ ] Implement the `Client` interface on a new struct (e.g., `GeminiClient`).
    -   [ ] Implement `GenerateOptions`.
        -   [ ] Craft a system prompt to generate two contrasting or deeper options based on history.
        -   [ ] Format choice history for the prompt context.
        -   [ ] Implement API call and parse the response.
    -   [ ] Implement `GenerateFinalValues`.
        -   [ ] Craft a system prompt to synthesize a ranked list of exactly 5 values with names and descriptions.
        -   [ ] Specify a strict, parsable output format in the prompt.
        -   [ ] Implement API call and parse the structured response into `core.RankedValues`.

-   [ ] **Step 8: Error Handling and Final Wiring**
    -   [ ] Modify `main.go`.
        -   [ ] Add logic to select between the mock and real LLM client (e.g., via flag or env var).
        -   [ ] Wrap all LLM client calls (`GenerateOptions`, `GenerateFinalValues`) in error-checking blocks.
        -   [ ] On LLM failure, print the exact message: `Sorry, we’re having trouble connecting to the AI service right now. Please try again later.`.
        -   [ ] On LLM failure, exit the program gracefully.
    -   [ ] Manually test the full application with the real LLM client.

-   [ ] **Step 9: Final Review & Manual Testing**
    -   [ ] Run a full session and check the timestamped log file for correctness and completeness.
    -   [ ] Verify the tone is casual and conversational throughout.
    -   [ ] Test edge cases: unexpected user input (e.g., text instead of numbers), early exit (Ctrl+C).
    -   [ ] Confirm the final output matches the specification exactly.