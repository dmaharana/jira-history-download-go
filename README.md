# Jira History Download

A command-line utility written in Go that downloads the change history of Jira issues based on a JQL search query and saves it to a CSV file.

## Features

- Search Jira issues using JQL query
- Retrieve complete change history for each issue
- Export history to CSV with the following information:
  - Issue Key
  - Author
  - Created Date
  - Field Changed
  - Old Value
  - New Value
- Configurable via INI file
- Automatic pagination for large result sets
- Timestamped output files

## Prerequisites

- Go 1.22 or higher
- Jira account with API access
- API token for authentication

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd jira-history-download
```

2. Build using make:
```bash
# Show available make targets
make help

# Build for all platforms (Linux and Windows)
make build

# Build for specific platform
make build-linux    # For Linux
make build-windows  # For Windows

# Clean, get dependencies, run tests, and build
make all
```

The compiled binaries will be available in the `build` directory:
- Linux: `build/jira-history-download_unix`
- Windows: `build/jira-history-download.exe`

## Configuration

Create a `configs/config.ini` file with your Jira settings:

```ini
[jira]
url = https://your-jira-instance.atlassian.net
username = your-email@domain.com
token = your-api-token
jql = project = "YOUR-PROJECT" ORDER BY created DESC
```

- `url`: Your Jira instance URL
- `username`: Your Jira email address
- `token`: Your Jira API token (can be generated from your Jira account settings)
- `jql`: JQL query to select the issues you want to process

## Usage

Run the compiled binary:

```bash
./jira-history-download
```

The program will:
1. Read the configuration from `configs/config.ini`
2. Connect to your Jira instance
3. Search for issues using the provided JQL
4. Download the change history for each issue
5. Save the results to a CSV file in the `output` directory

The output file will be named `jira_history_YYYY-MM-DD_HHMMSS.csv` and will be created in the `output` directory.

## Project Structure

```
jira-history-download/
├── cmd/
│   └── jira-history-download/    # Main application
│       └── main.go
├── configs/
│   └── config.ini               # Configuration file
├── internal/
│   ├── config/                  # Configuration handling
│   │   └── config.go
│   ├── jira/                    # Jira client and API interaction
│   │   └── client.go
│   └── helper/                  # Utility functions
│       └── utils.go
└── output/                      # CSV output directory
```

## Error Handling

The application includes comprehensive error handling for common scenarios:
- Configuration file not found or invalid
- Jira connection issues
- API authentication failures
- Invalid JQL queries
- File system operations

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
