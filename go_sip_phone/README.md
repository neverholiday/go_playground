# How to run locally ?
0. Before run locally, create env file first at `.env`
    ```sh
    GO_SIP_PHONE_UAC_SIP_URI="sip:111@uas:5060;transport=tcp"
    ```

1. Compose up
    ```sh
    docker compose up -f docker-compose.local.yml -d --build
    ```

2. Compose down
    ```sh
    docker compose -f docker-compose.local.yml down -v
    ```
