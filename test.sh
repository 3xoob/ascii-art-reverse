#!/bin/bash

# Colors for printing
RED='\033[0;31m'
NC='\033[0m' # No Color

# Function to run a test
run_test() {
    command=$1

    echo -e "Running test: ${command}"
    eval $command || echo -e "${RED}Error: Command failed${NC}"
    echo ""
}

run_test "go run . 'Ascii Art Team'"

# Test 1: Basic usage with a string
run_test "go run . 'Hello World!'"

# Test 2: Usage with a specific banner
run_test "go run . 'Hello World!' shadow"

# Test 3: Output to a file
run_test "go run . --output=output01.txt 'Hello World!'"

# Test 4: Output to a file with a banner
run_test "go run . --output=output02.txt 'Hello World!' standard"

# Test 5: Color option with text
run_test "go run . --color=red 'Hello World!'"

# Test 6: Color option with a substring
run_test "go run . --color=red 'Hello' 'Hello World!'"

# Test 7: Color option with a substring and banner
run_test "go run . --color=blue 'World' 'Hello World!' standard"

# Additional Tests for the reverse flag with various example files
run_test "go run . --reverse=TestCases/example.txt"
run_test "go run . --reverse=TestCases/example00.txt"
run_test "go run . --reverse=TestCases/example01.txt"
run_test "go run . --reverse=TestCases/example02.txt"
run_test "go run . --reverse=TestCases/example03.txt"
run_test "go run . --reverse=TestCases/example04.txt"
run_test "go run . --reverse=TestCases/example05.txt"
run_test "go run . --reverse=TestCases/example06.txt"
run_test "go run . --reverse=TestCases/example07.txt"

echo "All tests completed."
