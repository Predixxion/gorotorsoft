# RotorSoft Go Client

This Go client is designed to interact with the RotorSoft API, providing a convenient way to make requests and manage data. The client supports multiple versions of the API, making it flexible for various use cases.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Multiple API Versions**: Supports interaction with all of the RotorSoft API versions.
- **Simple Authentication**: Use your username and password for easy access to the API.
- **Custom HTTP Client**: Option to provide a custom `http.Client` for more control over HTTP requests.

## Installation

To install the RotorSoft Go client, ensure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/).

1. Clone the repository:

   ```bash
   go get github.com/Predixxion/gorotorsoft
   ```

2. Use Go modules to install dependencies:

   ```bash
   go mod tidy
   ```

3. Import the client in your project:

   ```go
   import "github.com/Predixxion/gorotorsoft"
   ```

## Usage

Here’s an example of how to create a new `Client` and use it to interact with the RotorSoft API:

```go
package main

import (
    "net/http"
    "fmt"
    "github.com/Predixxion/gorotorsoft"
)

func main() {
    httpClient := &http.Client{}
    url := "https://api.rotorsoft.com" // Replace with actual API URL
    username := "your-username"
    password := "your-password"

    client := NewClient(httpClient, url, username, password)

    // Example usage of the API client
    endpoints := client.V0.GetAllEndPoints()

    fmt.Println("Available endpoints:", endpoints)
}
```

## Contributing

Contributions are welcome! If you would like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Commit your changes.
4. Open a pull request.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

