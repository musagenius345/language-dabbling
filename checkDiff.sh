#!/bin/bash

# Loop through each file in the current directory
for file1 in *; do
    # Skip directories
    if [ -f "$file1" ]; then
        for file2 in *; do
            # Skip directories and the same file
            if [ -f "$file2" ] && [ "$file1" != "$file2" ]; then
                # Compare file contents using diff
                if diff "$file1" "$file2" >/dev/null; then
                    echo "Similar files: $file1 and $file2"
                fi
            fi
        done
    fi
done
