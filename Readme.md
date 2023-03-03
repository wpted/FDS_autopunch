# Femas Auto Punch
This is a tool for Freedomer to auto-punch.

## Run locally
```bash
    export USER="username"
    export USERPWD="userpwd"

    go run main.go
```


## Run with docker
```bash
    docker pull wpted/fds_punch
    docker run -e USER='username' -e USERPWD='userpwd' wpted/fds_punch
```

### Trigger it using Azure Functions or Virtual Machines(crontab).

- Create shell script(myjob.sh) to run image
```nano
    docker run -e USER='username' -e USERPWD='userpwd' wpted/fds_punch
```

- Make the script executable by running the command
```bash
    chmod +x myjob.sh
```

- Open crontab in terminal
```bash
    crontab -e
```

- Trigger with crontab
```nano
    0 9 * * 1-5 /path/to/run_image.sh // clock-in
    0 18 * * 1-5 /path/to/run_image.sh // clock-out
```