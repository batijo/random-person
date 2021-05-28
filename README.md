# random-person

A [website](rand.lt) to generate random person

## Dependencies

- `docker`
- `docker-compose`

## Setup project

```sh
docker network create web
```

- Configure `.env.example` files in `./` and `/config` directories and remove `.example` when done
- When you set `RP_PROD=true` in `/config/.env` CERT anf KEY files of TLS must be added to `/config` directory
- To upload data to database you need to add json files `names.json` and `surnames.json` to `/config` directory

`names.json` example:

```json names.json example
[
    {
        "name": "Jonas",
        "gender": 0,
        "normative_status": "teiktinas",
        "origin": "svetimos kilmės, hebrajiškas asmenvardis. svetimos kilmės, trumpinys",
        "note": "šventojo vardas"
    },
        {
        "name": "Jadvyga",
        "gender": 1,
        "normative_status": "teiktinas",
        "origin": "svetimos kilmės, germaniškas asmenvardis",
        "note": "šventosios vardas"
    }
]
```

`surnames.json` example:

```json
[
    {
        "surname": "Kazlauskas"
    },
    {
        "surname": "Stankevičius"
    }
]
```

## Run project

```sh
docker-compose up -d --build
```
