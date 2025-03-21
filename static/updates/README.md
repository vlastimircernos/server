# This file contains documentation related to the update files and their structure. 
# It may include instructions for adding new updates.

## Updates Directory

This directory is intended for storing update files for the client application. Each update should be structured as follows:

1. **Update Files**: Each update should be placed in its own directory named with the version number (e.g., `v1.0.1`).
2. **Update Manifest**: Each version directory should contain an `update.json` file that describes the update, including:
   - Version number
   - Release notes
   - Download URL for the update package
3. **Update Packages**: The actual update files (e.g., binaries, scripts) should be stored in the version directory.

## Adding a New Update

To add a new update, follow these steps:

1. Create a new directory for the version (e.g., `v1.0.1`).
2. Inside the new directory, create an `update.json` file with the following structure:

```json
{
  "version": "1.0.1",
  "release_notes": "Description of the changes in this update.",
  "download_url": "https://example.com/path/to/update-package.zip"
}
```

3. Place the update package files in the same directory.

## Example Structure

```
updates/
├── v1.0.0/
│   ├── update.json
│   └── update-package.zip
└── v1.0.1/
    ├── update.json
    └── update-package.zip
``` 

Ensure that all updates are tested before deployment to maintain application stability.