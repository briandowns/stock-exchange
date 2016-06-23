# stock-exchange

This repo serves as a learning experience on how to build a stock exchange and all of the individual components that make one up.  We're going to have primary and secondary components.  All primary components should share as much code as possible from the repo however be deployable independantly.

## Primary Components

### Engine

The engine is responsible for taking in orders, managing the securities being trades, executing trades, and notifying the reporting and clearing systems of those things.

#### Endpoints

| Method | Resource             | Description
| :----- | :-------             | :----------
| HEAD   | /api/v1/healthcheck  | Verifies the service is up and running | 
| GET    | /api/v1/stats        | Returns API processing statistics |
| POST   | /api/v1/order        | Add an order |
| DELETE | /api/v1/order/{id}   | Cancel an order |
| GET    | /api/v1/book         | Display the current book |
| GET    | /api/v1/book/{id}    | Retrieve an individual book entry |
| GET    | /api/v1/symbols      | Display all valid securities and associated data |
| GET    | /api/v1/symbols/{id} | Retrieve a symbol and it's data |

### Reporter

The reporter is responsible for relaying trade data to anyone subscribed.  This data can be provided in any number of ways from a REST API, websocket, etc.

### Clearing

The clearing system is responsible for receiving newly executed trades and running it through some rules to validate the trade.  If the trade was made erroneously, it'll be cancelled and the reporter will need to know and disseminate.  

## Secondary Components

### Redis

Redis can be used as a cache store for symbols and potentially for the Order Book.

### NSQ

NSQ can be used to notify components of events happening starting or completing in other components.


## Configuration

Configuration will start with a single configuration file with sections for each primary component.  In later versions we should look at more distributed means of sharing configuration and communicating changes.

## System Stand-Up

To get the entire environment stood up in order with all the needed components, we should look at using something like Docker Compose.
