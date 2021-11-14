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

## API endpoints
| Endpoint                                   | Request Method           | Auth | Description                      |
| ------------------------------------------ | ------------------------ | ---- | -------------------------------- |
| `/`                                        | `GET`                    | no   | `n/a` |
| `/auth/google`                             | `GET`                    | no   | Redirects to Google oauth endpoint |
| `/auth/google/redirect`                    | `GET`                    | no   | Generates a token given the Google oauth callback |
| `/auth/profile`                            | `GET`                    | yes  | Given a token, returns the profile details for the authenticated user |
| `/products`                                | `GET`                    | no   | Returns a list of products with given a set of query parameters to tune the results |
| `/events`                                  | `GET`, `POST`            | yes  | Create and fetch events for an authenticated user |
| `/events/:id`                              | `GET`, `PATCH`, `DELETE` | yes  | Fetch, update, or delete a specific event for an authenticated user. Updating or deleting an event requires the user to be an event orgranizer |
| `/events/get-details/:linkCode`            | `GET`                    | yes  | Returns the name and description (if exists) for a specific event |
| `/events/invites`                          | `GET`                    | yes  | Returns a list of all pending invites for an authenticated user |
| `/events/invites/accept/:eventId`          | `GET`                    | yes  | Accepts the event invite for an authenticated user |
| `/events/invites/decline/:eventId`         | `GET`                    | yes  | Declines the event invite for an authenticated user |
| `/events/get-link/:eventId`                | `GET`                    | yes  | Generates the invite link to a specific event |
| `/events/get-link/:eventId`                | `GET`                    | yes  | Generates the invite link to a specific event |
| `/events/verify-invite-code/:inviteCode`   | `GET`                    | yes  | Verify the invite code for a specific event |
| `/events/invite-code/:inviteCode`          | `GET`                    | yes  | Add the event to the user's pending invites list |
| `/participants/manage`                     | `PATCH`, `DELETE`        | yes  | Update participant details, or remove them from event. Requires that the participant is also an organizer |
| `/participants/:eventId/:participantId`    | `GET`                    | yes  | Fetch participant information, given the user is part of the same event |
| `/participants/:participantId`             | `PATCH`, `DELETE`        | yes  | Allows user to manage their participant information, including leaving event, and updating address |

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