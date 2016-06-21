# stock-exchange

This repo serves as a learning experience on how to build a stock exchange and all of the individual components that make one up.  We're going to have primary and secondary components. 

## Primary Components

### Engine

The engine is responsible for taking in orders, managing the securities being trades, executing trades, and notifying the reporting and clearing systems of those things.

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

We should start with a single configuration file with sections for each primary component.  In later versions we should look at more distributed means of sharing configuration and communicating changes.

## System Stand-Up

To get the entire environment stood up in order with all the needed components, we should look at using something like Docker Compose.
