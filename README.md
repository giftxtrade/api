<p align="center">
    <a href="http://giftxtrade.com/" target="blank">
        <!-- <img src="https://giftxtrade.com/logos/logo_profile_rounded.svg" width='50' alt="GiftTrade Logo" /> -->
        <img src="https://giftxtrade.com/logos/logotype_rounded_color.svg" width='250' alt="GiftTrade Logo" />
    </a>
</p>

<p align="center">
    The platform that aims to simplify your online gift exchange and secret santa for <i>free</i>.
</p>

<br />

## Description
The GiftTrade API repository serves as the REST API for the [giftxtrade.com](https://giftxtrade.com) web app. This repo is designed to work with a fully working MySQL database.


## Set up

### Clone repository
```
git clone git@github.com:giftxtrade/api.git
```

### Install dependencies
```
npm install
```

### Configure database connection
This repo requires a working connection with a MySQL database and uses TypeORM to manage models and connections with the database.
To set up the config file with the connection details, create a file named `ormconfig.json` in the root of the project directory, then copy the code below, replacing all `<...>` with the appropriate values for you local database.
```json
{
    "type": "mysql",
    "host": "localhost",
    "port": 3306,
    "username": "<username>",
    "password": "<password>",
    "database": "<database>",
    "entities": [
        "dist/**/*.entity{.ts,.js}"
    ],
    "synchronize": true
}
```

### Start server in watch-mode
```
npm start:dev
```

### Start server
```
npm start
```

### Build server
```
npm build
```