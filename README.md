# SaFreeBot: ERC20 Token Security Checker and Honeypot Detector

SaFreeBot is a Telegram bot designed to enhance the security analysis of ERC20 tokens by inspecting various functions, identifying potential vulnerabilities, and detecting honeypot tokens. This project is implemented in Golang, JavaScript, and Solidity to provide a comprehensive set of tools for evaluating the security of ERC20 tokens.

## Features

### 1. ERC20 Token Function Analysis

SaFreeBot analyzes the functions of an ERC20 token to determine which functions the deployer can call to initiate a token transfer lock. By identifying these functions, users can gain insights into the potential risks associated with a token.

### 2. Honeypot Detection

SaFreeBot tests ERC20 tokens to check for honeypot characteristics. It helps users assess whether a token exhibits behavior indicative of a honeypot, providing an additional layer of security analysis.

### 3. Activity Monitoring

The bot monitors the activities of an ERC20 token for a specified duration (2 minutes by default). During this period, SaFreeBot observes token transfers and other relevant activities, generating reports to update users on the token's behavior.

### 4. Telegram Address Extraction

SaFreeBot is equipped with the capability to extract embedded Telegram addresses within the ERC20 token. This feature can be useful for identifying communication channels associated with the token.

## Getting Started

To use SaFreeBot, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/Thankgod20/SaFreeBot.git
```

2. Install dependencies:

   - Golang: Follow the official [Golang installation guide](https://golang.org/doc/install).
   - Node.js: Install Node.js and npm from [https://nodejs.org/](https://nodejs.org/).

3. Set up the environment:

   ```bash
   # Go to the project directory
   cd SaFreeBot

   # Install Golang dependencies
   go mod download

   # Install Node.js dependencies
   cd frontend
   npm install
   ```

4. Configure the bot:

   - Set up a Telegram bot on [BotFather](https://core.telegram.org/bots#botfather) and obtain the token.
   - Create a `.env` file in the project root with the following content:

     ```env
     TELEGRAM_BOT_TOKEN=<your-telegram-bot-token>
     ```

5. Run the bot:

   ```bash
   go run main.go
   ```

## Usage

1. Start a conversation with SaFreeBot on Telegram.
2. Use the available commands to analyze ERC20 tokens, detect honeypots, monitor activities, and extract Telegram addresses.

## Contributing

Contributions are welcome! Feel free to open issues, submit pull requests, or suggest improvements.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Note:** SaFreeBot is a tool designed for security analysis and informational purposes. It is not intended for malicious use. Users should use this tool responsibly and adhere to ethical guidelines when evaluating ERC20 tokens.
