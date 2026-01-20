#!/bin/bash

# KidLang Test Runner
# Builds kidlang and runs all tests

set -e

# Color codes
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ENGINE_DIR="$SCRIPT_DIR/../engine"
KIDLANG_BIN="$ENGINE_DIR/kidlang"

echo "======================================"
echo "  KidLang Test Suite"
echo "======================================"
echo ""

# Build kidlang
echo "Building kidlang..."
cd "$ENGINE_DIR"
if go build -o kidlang; then
    echo -e "${GREEN}✓${NC} Build successful"
else
    echo -e "${RED}✗${NC} Build failed"
    exit 1
fi
echo ""

# Initialize counters
TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0

# Function to run a single test
run_test() {
    local test_file="$1"
    local test_name=$(basename "$test_file" .kid)
    local expected_file="${test_file%.kid}.expected"
    local setup_file="${test_file%.kid}.setup"
    local input_file="${test_file%.kid}.input"
    local output_file="/tmp/kidlang_test_output_${test_name}.txt"
    
    TOTAL_TESTS=$((TOTAL_TESTS + 1))
    
    # Run setup if exists
    if [ -f "$setup_file" ]; then
        bash "$setup_file" 2>/dev/null
    fi
    
    # Run the test with or without input file
    if [ -f "$input_file" ]; then
        # Test has input file - pipe it to kidlang
        if "$KIDLANG_BIN" "$test_file" < "$input_file" > "$output_file" 2>&1; then
            run_status=0
        else
            run_status=1
        fi
    else
        # No input file - run normally
        if "$KIDLANG_BIN" "$test_file" > "$output_file" 2>&1; then
            run_status=0
        else
            run_status=1
        fi
    fi
    
    if [ $run_status -eq 0 ]; then
        # Compare output
        if [ -f "$expected_file" ]; then
            if diff -q "$output_file" "$expected_file" > /dev/null 2>&1; then
                echo -e "${GREEN}✓${NC} $test_name"
                PASSED_TESTS=$((PASSED_TESTS + 1))
            else
                echo -e "${RED}✗${NC} $test_name"
                echo "   Expected output:"
                cat "$expected_file" | sed 's/^/     /'
                echo "   Actual output:"
                cat "$output_file" | sed 's/^/     /'
                FAILED_TESTS=$((FAILED_TESTS + 1))
            fi
        else
            echo -e "${YELLOW}⚠${NC} $test_name (no expected output file)"
        fi
    else
        echo -e "${RED}✗${NC} $test_name (execution failed)"
        echo "   Error output:"
        cat "$output_file" | sed 's/^/     /'
        FAILED_TESTS=$((FAILED_TESTS + 1))
    fi
    
    # Cleanup
    rm -f "$output_file"
}

# Run all tests
echo "Running tests..."
echo ""

for test_file in "$SCRIPT_DIR"/*.kid; do
    if [ -f "$test_file" ]; then
        run_test "$test_file"
    fi
done

# Summary
echo ""
echo "======================================"
echo "  Test Results"
echo "======================================"
echo "Total:  $TOTAL_TESTS"
echo -e "Passed: ${GREEN}$PASSED_TESTS${NC}"
if [ $FAILED_TESTS -gt 0 ]; then
    echo -e "Failed: ${RED}$FAILED_TESTS${NC}"
else
    echo -e "Failed: $FAILED_TESTS"
fi
echo ""

# Cleanup temp files
rm -f /tmp/kidlang_test_data.txt
rm -f /tmp/kidlang_test_output.txt
rm -f /tmp/kidlang_file_test.txt

# Exit with failure code if any tests failed
exit $FAILED_TESTS
