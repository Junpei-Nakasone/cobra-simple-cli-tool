# My CLI Tool

My CLI Tool is a simple command-line interface (CLI) application built with Go and the Cobra library.

## Features

- **dayofweek**: Prints the current day of the week.

## Installation

1. **Clone the repository**:

    ```sh
    git clone https://github.com/Junpei-Nakasone/cobra-simple-cli-tool.git
    cd cobra-simple-cli-tool
    ```

2. **Build the application**:

    ```sh
    go build -o mycli .
    ```

## Usage

After building the application, you can run the CLI tool using the following commands:

- **Print the current day of the week**:

    ```sh
    ./mycli dayofweek
    ```

## Docker

You can also run the CLI tool inside a Docker container.

1. **Build the Docker image**:

    ```sh
    docker build -t mycli .
    ```

2. **Run the Docker container**:

    ```sh
    docker run --rm mycli dayofweek
    ```

