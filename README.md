# random-person

A [api](https://rand.lt/api/v0/person) to generate random Lithuanian person.

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

`names.json` example:

```json
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

- To create age weights we need to add json file `age_weights.json` to `/config` directory. If you will leave gaps between ages program will fill them with values from the last weight.

`age_weights.json` example:

```json
[
    {
        "age": "1",
        "weight": "1"
    },
    {
        "age": "3",
        "weight": "2"
    },
]
```

program will convert this to:

```json
[
    {
        "age": "0",
        "weight": "0"
    },
    {
        "age": "1",
        "weight": "1"
    },
    {
        "age": "2",
        "weight": "1"
    },
    {
        "age": "3",
        "weight": "2"
    },
]
```

- To create email domain weights we need to add json file `email_domains.json` to `/config` directory.

`email_domains.json` example:

```json
[
    {
        "data": "gmail.com",
        "weight": 1
    },
    {
        "data": "yahoo.com",
        "weight": 2
    },
]
```

- To create email template weights we need to add json file `email_templates.json` to `/config` directory.

`email_templates.json` example:

```json
[
    {
        "data": "[fn]",
        "weight": 1
    },
    {
        "data": "[fs]",
        "weight": 2
    },
]
```

email template creation docs:

```md
[fn] - inserts full persons name
[fs] - inserts full persons surname
[nws] - inserts name without suffix
[sws] - inserts surname without suffix
[by] - inserts birth year
[pby] - inserts partial birth year (if year is 1985, inserts 85)
if you add a number N after any command it will take N number of characters from the start of a result
e.g. Name is Jonas so [fn2] is Jo
[command{3/2}] - command can be any command , number 3 represents which element, 2 how many time multiply it
e.g. Surname is Kazlauskas so [sws{4/3}] is Kazlllausk
number 3 can be replaced with e for last letter
e.g. Surname is Kazlauskas so [fs{e/3}] is Kazlauskasss
if you add number x after e it will multiply x letter from end
e.g. Surname is Kazlauskas so [sws{e3/4}] is Kazlauuuusk
if you add v (vowel) then [command{v/2}] multiplies first vowel 2 times
e.g. Surname is Kazlauskas so [sws{v/3}] is Kaaazlausk
You can also add e before v
e.g. Surname is Kazlauskas so [fs{ev/4}] is Kaazlauskaaaas
everything what goes after @ symbol is added without checking
if you don't add @ a random popular domain will be added
e.g. @gmail.com @yahoo.com @outlook.com ...
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

`rand.lt/api/v0/person` is a combination of both. You can provide gender and marital status as well. Also it will include weighted random birth date and email.

`rand.lt/api/v0/version` will return version of the api.
