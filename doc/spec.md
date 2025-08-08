
# Personal Values Discovery CLI Tool — Developer Specification

---

## 1. Overview

A command-line tool, implemented in Go, designed to help users discover their personal values through a series of simple choice comparisons. The system uses an LLM to generate pairs of options that contrast or drill deeper into the user’s preferences based on their previous choices. After a fixed number of rounds, the LLM synthesizes a ranked list of 5 values with descriptions.

---

## 2. Requirements

### 2.1 Functional Requirements

* **User Outcome:**
  The tool outputs a ranked list of 5 personal values, each with a medium-length description.

* **Choice Interaction:**

  * Each round presents a pair of options in the format:
    `Which feels more important to you right now:`
    `1) <Option A>`
    `2) <Option B>`
  * User responds by typing the number `1` or `2`.
  * Number of options per question is configurable, defaulting to 2.

* **Session Flow:**

  * Fixed number of rounds per session (configurable, e.g., 20).
  * Progress indicator after each round (e.g., “Round 3 of 20”).

* **AI Behavior:**

  * LLM generates options using the entire history of previous choices for context.
  * AI aims to produce contrasting or deeper options, occasionally repeating wording but not frequently.
  * No topic restrictions - AI can use any themes (principles, priorities, hobbies, etc.).

* **Session Completion:**

  * After all rounds, the LLM receives the full choice history and generates a ranked list of exactly 5 values with names and medium-length (2–3 sentence) descriptions.
  * The final output is displayed directly to the user with no extra closing message.

* **User Interaction:**

  * Interface is command-line prompts only.
  * No ability to change previous answers once entered.
  * No initial user input before the first round.

### 2.2 Non-Functional Requirements

* **Tone:** Casual and conversational throughout questions and results.

* **Data Persistence:**

  * All session data (questions, choices, responses, final ranked values) must be saved to a separate, human-friendly timestamped text file per session for troubleshooting and analysis.

* **Configuration:**

  * Use a YAML config file to store parameters such as:

    * Number of rounds
    * Number of options per question
    * Other relevant settings

* **LLM Integration:**

  * Use an abstracted Go package (or internal wrapper) for sending prompts and receiving responses from the LLM API (e.g., OpenAI or Gemini).
  * Two distinct prompts: one for generating options during rounds, one for synthesizing the final ranked list.

* **Error Handling:**

  * If an LLM call fails or times out, print a clear, friendly error message and exit gracefully.

* **User Input:**

  * No validation on user input (accept whatever is typed).

---

## 3. Architecture & Components

### 3.1 Core Components

* **CLI Interface:**
  Handles user prompts, input collection, and progress display.

* **Session Manager:**
  Maintains session state, including:

  * History of all questions and user choices
  * Current round and total rounds
  * Writing session data to file

* **LLM Client Wrapper:**
  Abstracts communication with the LLM API, provides methods for:

  * `GenerateOptions(history []Choice) ([]string, error)`
  * `GenerateFinalValues(history []Choice) (RankedValues, error)`

* **Configuration Loader:**
  Loads YAML config file at startup and validates essential parameters.

### 3.2 Data Structures

```go
type Choice struct {
    QuestionText string   // Full text of the question prompt
    Options      []string // List of options presented
    Selected     int      // Index of user-selected option
}

type RankedValue struct {
    Name        string
    Description string
}

type RankedValues []RankedValue

type SessionData struct {
    Timestamp     time.Time
    Choices       []Choice
    FinalRanking  RankedValues
}
```

### 3.3 Session Data Storage

* Store session in a file named like: `values_session_2025-08-08_14-30-05.txt`
* Include:

  * Timestamp
  * Each question and presented options
  * User’s choice per question
  * Final ranked values with descriptions

---

## 4. Interaction Flow

1. Load config file (YAML).
2. Initialize session and open a new timestamped log file.
3. For each round from 1 to N (configurable):

   * Generate next question with options by passing full history to LLM.
   * Display progress and question with numbered options.
   * Collect user input.
   * Append question, options, and user choice to history and log file.
4. After all rounds, send full choice history to LLM for final ranked values.
5. Display ranked values and descriptions in casual tone.
6. Write final output and all session details to log file.
7. Exit.

---

## 5. Error Handling Strategy

* If LLM call fails, print:
  `Sorry, we’re having trouble connecting to the AI service right now. Please try again later.`
* Cleanly exit the program after logging error.
* No input validation; unexpected inputs are accepted as-is.

---

## 6. Testing Plan

### 6.1 Unit Tests

* Config loader: parse valid/invalid YAML.
* Session manager: appending choices, session state correctness.

### 6.2 Integration Tests

* Simulate full session flow with mock LLM responses:

  * Generate options prompt produces valid options.
  * Final ranking prompt returns valid ranked values.
* Test file logging creates correctly named file with expected contents.
* Verify CLI interaction flow and progress display.

### 6.3 Manual Testing

* Run full sessions with real LLM API (if available).
* Test edge cases: LLM failures, unexpected user inputs, early exit.
* Confirm casual tone consistency in prompts and final results.

---

## 7. Future Enhancements (Out of Scope for Initial Version)

* Variable number of options per question.
* Adaptive stopping based on AI confidence or user input.
* Input validation and correction options.
* Session resume or review of past choices during session.
* GUI or web interface.
* More sophisticated AI prompt engineering or local model fallback.
