#!/bin/bash

# Script to generate ABI from Solidity contracts
# Requires solc (Solidity compiler) to be installed

# Configuration
CONTRACT_DIR="infrared-contracts/src"  # Directory containing Solidity contracts
ABI_DIR="abi"            # Directory to store generated ABIs
SOLC_VERSION="0.8.26"
# # Check if solc is installed
# if ! command -v solc &> /dev/null; then
#     echo "Error: solc (Solidity compiler) is not installed."
#     echo "Please install it using: npm install -g solc"
#     echo "Or visit https://docs.soliditylang.org/en/latest/installing-solidity.html"
#     exit 1
# fi

# Create abi directory if it doesn't exist
mkdir -p "$ABI_DIR"
echo "ABI directory: $ABI_DIR"

# Check if contracts directory exists
if [ ! -d "$CONTRACT_DIR" ]; then
    echo "Error: Contract directory '$CONTRACT_DIR' does not exist."
    exit 1
fi

# Find all Solidity files in the contract directory
CONTRACT_FILES=$(find "$CONTRACT_DIR" -name "*.sol")

# Check if any contract files were found
if [ -z "$CONTRACT_FILES" ]; then
    echo "Error: No Solidity files found in '$CONTRACT_DIR'."
    exit 1
fi

echo "Found Solidity contracts:"
echo "$CONTRACT_FILES"
echo "-------------------------"

# Process each contract file
for CONTRACT_FILE in $CONTRACT_FILES; do
    echo "Processing: $CONTRACT_FILE"
    
    # Extract contract name without path and extension
    CONTRACT_NAME=$(basename -- "$CONTRACT_FILE")
    CONTRACT_NAME="${CONTRACT_NAME%.sol}"
    
    # Generate ABI JSON using solc
    ABI_FILE="$ABI_DIR/${CONTRACT_NAME}.abi.json"
    
    echo "Generating ABI for $CONTRACT_NAME..."
    
    # Use solc to compile and extract ABI
    npx solc@${SOLC_VERSION} --abi "$CONTRACT_FILE" -o "$ABI_DIR"
    
    # Check if ABI was generated successfully
    if [ -f "$ABI_DIR/${CONTRACT_NAME}.abi" ]; then
        # Rename the file to .json extension for better compatibility
        mv "$ABI_DIR/${CONTRACT_NAME}.abi" "$ABI_FILE"
        echo "ABI successfully generated: $ABI_FILE"
    else
        echo "Error: Failed to generate ABI for $CONTRACT_NAME"
    fi
    
    echo "-------------------------"
done

echo "ABI generation completed."