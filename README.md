# mcp-notify

A simple MCP server for sending desktop notifications.

## Supported Environments

Currently supports:

- macOS

Support for Linux and Windows is planned for future releases.

## Installation

```bash
go install github.com/5n7/mcp-notify@latest
```

## Configuration

For [Cursor](https://www.cursor.com/):

```json
{
  "mcpServers": {
    "notify": {
      "command": "/path/to/mcp-notify",
      "args": []
    }
  }
}
```

## Usage

The following parameters are available according to the MCP tool specification:

- `title`: The title of the notification
- `message`: The content of the notification
