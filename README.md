# Clamper

Clamper is a resource throttling tool designed to manage and limit CPU and RAM usage for applications in test environments.

## Features
- CPU Limitation: Restrict the number of CPU cores available to an application.
- RAM Limitation: Set a maximum amount of RAM usage for an application.
- Automatic Monitoring: Continuously monitor resource usage and terminate applications if limits are exceeded.
- Cross-Platform Support: Available for Linux, macOS, and Windows.

## Installation

To install Clamper, use the following command which will automatically download and install the latest release:

```bash
curl -s https://raw.githubusercontent.com/julianofirme/clamper/main/install.sh | bash
```

## Usage

Once Clamper is installed, you can run commands with resource limits using the run command.

## Running a Command with Resource Limits

```bash
clamper [flags] run [command]
```

## Flags
- `--cores`: Number of CPU cores
- `--clock`: CPU clock speed in MHz 
- `--ram`: RAM in MB

## Example
```bash
clamper --cores 2 --clock 70 --ram 200 run your_app_command
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)