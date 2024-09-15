# PM

This is a simple yet powerful command-line password manager written in Go. It allows users to securely generate, store, and retrieve passwords, using AES encryption for storage security. Perfect for developers and sysadmins who prefer managing passwords locally and securely without third-party services.

## Features

- **Setup**: Initialize the manager with a secure configuration file.
- **Generate**: Create strong, random passwords for any username.
- **Save**: Store passwords securely using AES encryption.
- **List**: Display all stored passwords (encrypted).
- **Show**: Decrypt and display specific passwords when needed.

## Prerequisites

- **Go**: Ensure you have [Go](https://golang.org/doc/install) installed on your machine.
- **Git**: Ensure you have [git](https://git-scm.com/) installed to clone the repo.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/Ace-h121/PM
    cd PM
    ```

2. Build the project:

    ```bash
    go build -o password-manager
    ```

3. Run the setup:

    ```bash
    ./password-manager setup
    ```

## Usage

Once installed, you can begin using the tool with the following commands:

### 1. **Setup**

Before saving any passwords, run the `setup` command to generate the configuration file:

```bash
./password-manager setup
```

This will create a `.config/PM.conf` file with a randomly generated key for encryption.

### 2. **Generate Password**

Generate a new 32-character password for a specific user:

```bash
./password-manager generate <username>
```

Example:

```bash
./password-manager generate alice
```

This will output the generated password for the user `alice`.

### 3. **Save Password**

To save a password securely, use the `save` command. The password will be encrypted and stored.

```bash
./password-manager save <filename> <password>
```

Example:

```bash
./password-manager save mypass.txt abc123
```

This will encrypt and save the password `abc123` to `mypass.txt`.

### 4. **List All Passwords**

To list all saved passwords (still encrypted):

```bash
./password-manager list
```

### 5. **Show Password**

To decrypt and show a saved password:

```bash
./password-manager show <filename>
```

Example:

```bash
./password-manager show mypass.txt
```

This will display the decrypted password stored in `mypass.txt`.

## Why Use This Tool?

- **Security**: Your passwords are encrypted using AES, ensuring they are safe from unauthorized access.
- **Control**: Store passwords locally, without relying on cloud-based solutions that may expose your data.
- **Simplicity**: A straightforward tool for users who prefer the command line and want to manage their passwords without extra frills.

## Contributing

Feel free to fork this repository and submit pull requests. All contributions are welcome to improve functionality or add new features!

## License

This project is licensed under the MIT License.

---

This CLI tool offers a streamlined, secure way to manage your passwords directly from the command line. If you're looking for local control and encryption-based security without the complexities of external password managers, this project is your answer.
