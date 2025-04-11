#!/bin/bash

# Exit on error
set -e

# Directory paths
ABI_DIR="./abi"
OUTPUT_DIR="./bindings"

# Create output directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# Check if the abi directory exists
if [ ! -d "$ABI_DIR" ]; then
    echo "Error: ABI directory '$ABI_DIR' does not exist."
    exit 1
fi

# Check if abigen is installed
if ! command -v abigen &> /dev/null; then
    echo "Error: abigen is not installed or not in PATH."
    echo "Install it using: go install github.com/ethereum/go-ethereum/cmd/abigen@latest"
    exit 1
fi

# Get list of ABI files
abi_files=($(find "$ABI_DIR" -name "*.json"))
total_files=${#abi_files[@]}

if [ $total_files -eq 0 ]; then
    echo "Warning: No ABI files found in $ABI_DIR"
    exit 0
fi

echo "Generating Go bindings for $total_files ABI files..."

# Process each JSON file in the ABI directory
for abi_file in "${abi_files[@]}"; do
    # Get filename without extension and path
    filename=$(basename "$abi_file" .json)
    
    # Convert filename to PascalCase for Go package naming
    # This removes hyphens and underscores and capitalizes the first letter after them
    contract_name=$(echo "$filename" | sed -E 's/(^|[-_])([a-z])/\U\2/g')
    
    # Create lowercase package name
    package_name="bindings"
    
    current=$((current + 1))
    echo "Processing $filename ($current/$total_files)..."
    
    # Generate Go bindings
    if ! abigen --abi "$abi_file" \
           --pkg "$package_name" \
           --type "$contract_name" \
           --out "$OUTPUT_DIR/${filename}.go"; then
        echo "Error: Failed to generate bindings for $contract_name"
        exit 1
    fi
    
    echo "✓ Generated bindings for $contract_name"
done

echo "All done! Go bindings have been generated in $OUTPUT_DIR"
