# Event Sourcing with Go and Redis
NOTE: This code is not tested, just an experiment

I thought you already heard about Event Sourcing in the past recent year.
But let's go through the definition again.

> Capture all changes to an application state as a sequence of events.
> Event Sourcing ensures that all changes to application state are stored as a sequence of events. - [Martin Fowler](https://martinfowler.com/eaaDev/EventSourcing.html)

If you know bitcoin/blockchain you will know it's quite similar with Event Sourcing.

> Your current balance (Application State) is calculated from a series of events in history (in the chain)
![Alt Text](https://thepracticaldev.s3.amazonaws.com/i/ztik9xqelulsh4lx3kl9.png)

so you don't have a table like this in database

|user_id|balance|
|----|----|
| 10 | 100$|
| 7  | 200$|

now you have

|events|
|------|
|user x top-up event|
|user buy 5 items event|
|user y top-up event|

I've read many articles/blog posts about Event Sourcing so I try to make once.

## What we will build?
Let's say you have an e-commerce website and users can buy items from your website.
Source: https://github.com/felixvo/lmax

Entities:  
- `User` will have `balance`.
- `Item` will have `price` and number of `remain` items in the warehouse.

Events:  
- `Topup`: increase user balance
- `AddItem`: add more item to warehouse
- `Order`: buy items

## Directory Structure

```
├── cmd
│   ├── consumer       # process events
│   │   ├── handler    # handle new event base on event Type
│   │   └── state
│   └── producer       # publish events
└── pkg
    ├── event          # event definition
    ├── snapshot       # snapshot state of the app
    ├── user           # user domain
    └── warehouse      # item domain
```

## Architecture
![Alt Text](https://thepracticaldev.s3.amazonaws.com/i/ui1bv7ili5wag324ucil.png)

- Event storage: [Redis Stream](https://redis.io/topics/streams-intro)  
>Entry IDs
>The entry ID returned by the XADD command, and identifying univocally >each entry inside a given stream, is composed of two parts:
>`<millisecondsTime>-<sequenceNumber>`
> I use this `Entry ID` to keep track of processed event

- The consumer will consume events and build the application state
- `snapshot` package will take the application state and save to redis every 30s. Application state will restore from this if our app crash

## Run

### Producer

First, start the producer to insert some events to `redis stream`
![Alt Text](https://thepracticaldev.s3.amazonaws.com/i/v0scin2iq1cdcnu9nkaw.png)

### Consumer

Now start the consumer to consume events

![Alt Text](https://thepracticaldev.s3.amazonaws.com/i/eueho865f5couulhbnal.png)
Because the consumer consumes the events but not backup the state yet.
If you wait for more than 30s, you will see this message from console

![Alt Text](https://thepracticaldev.s3.amazonaws.com/i/0qkcmgzgwlhctk51q3bj.png)

Now if you stop the app and start it again, the application state will restore from the latest snapshot, not reprocess the event again

![Alt Text](https://thepracticaldev.s3.amazonaws.com/i/vxxwbn2hfho3qj1hgi3b.png)

Thank you for reading!
I hope the source code is clean enough for you to understand :scream:

## Thoughts


