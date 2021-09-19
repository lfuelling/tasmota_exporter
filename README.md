# tasmota_exporter

Prometheus exporter for Tasmota devices.

## Usage

1. Clone this repository
2. Run `go build`
3. Copy `examples/config.json` to `config.json`
   3.1. Edit the configuration to fit your setup
4. (Optional) copy `examples/tasmota_exporter.service` to `/etc/systemd/system/tasmota_exporter.service`
   1. Run `systemctl daemon-reload`
   2. Run `systemctl enable tasmota_exporter.service`
5. Start the exporter either by running the executable directly or starting the service

### Command line params

- `config`
    - Path to the configuration file
    - Default: `./config.json`

