#!/bin/bash

WORKSPACE_DIR=$(pwd)

find "$WORKSPACE_DIR" -name "*.go" | while read -r gofile; do
    # Apply gofmt
    gofmt -w "$gofile"
    
    goimports -w "$gofile"
    
    gopls format "$gofile"
    
    echo "Formatted: $gofile"
done

echo "Formatting completed for all Go files in $WORKSPACE_DIR"