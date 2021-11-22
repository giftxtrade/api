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
| `/events/:id`                              | `GET`, `PATCH`, `DELETE` | yes  | Fetch, update, or delete a specific event for an authenticated user. Updating or deleting an event requires the user to be an event organizer |
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
| `/wishes`                                  | `POST`, `DELETE`         | yes  | Creates or removes a wishlist item |
| `/wishes/:id`                              | `GET`                    | yes  | Fetch all wishlist items from the user participant given an event id |
| `/wishes/:eventId/:participantId`          | `GET`                    | yes  | Fetch all wishlist items from a participant given the participant's id and an event id |
| `/draws`                                   | `POST`                   | yes  | Creates randomized pairings from the active participants. Requires that user is an organizer |
| `/draws/confirm/:eventId`                  | `GET`                    | yes  | Confirms the generated draws and sends emails to all the participants. Requires that user is an organizer |
| `/draws/:eventId`                          | `GET`                    | yes  | Fetch all pairings for a given event. Requires that user is an organizer |
| `/me/:eventId`                             | `GET`                    | yes  | Fetch user's draw for a given event |
| `/categories`                              | `GET`                    | no   | Fetches all product categories |

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

### Configure API keys
GiftTrade API requires a number of API keys from Google, Sendgrid, (and possible from Amazon in the near future). To configure the authentication tokens file, copy `auth-tokens.sample.json` file from the root of the project, then rename the file to `auth-tokens.json`, or use the following command in the terminal:
```
cp auth-tokens.sample.json ./auth-tokens.json
```
Once `auth-tokens.json` is present, make sure to replace all values with the appropriate values, if working locally, you can leave `FRONTEND_BASE` and `JWT.SECRET` as it is. However, it is essential that at the very least you create an account on [Google Console](https://console.cloud.google.com/) and set up an OAuth key and paste the appropriate keys on the `GOOGLE` section. Do the same for [SendGrid](https://sendgrid.com) and use the free tier to get access to the appropriate API tokens.

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