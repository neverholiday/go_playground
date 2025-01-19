# GO_SIP_PHONE

## Env file for this repo
Please create `.env` file and copy this content to the file.
```sh
GO_SIP_PHONE_UAC_SIP_URI="sip:111@uas:5060;transport=tcp"
GO_SIP_PHONE_REGISTER_SIP_URI="sip:<user>@<host>:<port>"
```

## How to run locally for uas and uac ?
1. Compose up
    ```sh
    docker compose up -f docker-compose.local.yml -d --build
    ```

2. Compose down
    ```sh
    docker compose -f docker-compose.local.yml down -v
    ```

## How to run register scripts ?
1. Export env by
    ```sh
    export $(grep -v '^#' .env | xargs)
    ```
