# Femas Auto Punch
This is a tool for Freedomer to auto-punch.


### Run with docker
```bash
    docker pull wpted/fds_punch
    docker run -e USERNAME='username' -e USERPWD='userpwd' wpted/fds_punch
```

# Can trigger it using Azure Functions or Virtual Machines(crontab).

## Create shell script to run image
```sh
    docker run -e USERNAME='username' -e USERPWD='userpwd' wpted/fds_punch
```

## Trigger with crontab
```crontab
    0 9 * * 1-5 /path/to/run_image.sh // clock-in
    0 18 * * 1-5 /path/to/run_image.sh // clock-out
```