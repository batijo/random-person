# random-person

A [website](https://rand.lt) to generate random Lithuanian person.

## Setup

### Dependencies

- `docker`
- `docker-compose`

### Initial app setup

```sh
docker network create web
```

- Configure `.env.example` files in `./` and `/config` directories and remove `.example` when done
- To upload data to database you need to add json files `names.json` and `surnames.json` to `/config` directory on first start. You should remove those files later for better start up performance.
- `docker-compose` will not create config file for [watchtower](https://github.com/containrrr/watchtower) if it does not exist so you should create `config.json` file in `$HOME/.docker/` yourself.

  ```sh
  echo {} > $HOME/.docker/config.json
  ```

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

### Run app

```sh
docker-compose up -d --build
```

## Api usage

### v0

`rand.lt/api/v0/name` will return random name.

`rand.lt/api/v0/name/{gender}` replace `{gender}` with `0` or `male` for only random male names and respectively `1` or `female` for female names

`rand.lt/api/v0/surname/female?m_status=1` surname works the same as names except it can accept query parameter for marital status (female surnames only):

- 0 - not married.
- 1 - married.
- 2 - without marital status.

`rand.lt/api/v0/person` is a combination of both. You can provide gender and marital status as well.
