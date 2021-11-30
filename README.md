# go-reservation-api

# Getting started

## Prerequisite

* Install Go, Serverless CLI, AWS CLI

## Build and deploy backend

In the root directory of this project:

* `make build`
* `sls deploy --stage dev`

## Routes

POST https://45y1x9o537.execute-api.us-east-1.amazonaws.com/dev/room/reservations
GET  https://45y1x9o537.execute-api.us-east-1.amazonaws.com/dev/room/{roomId}/reservations
DEL  https://45y1x9o537.execute-api.us-east-1.amazonaws.com/dev/room/{roomId}/reservations/{reservationId}

# Design choices

Hi, to address the problem I took the decision to go with a Serverless architecture that will speed some processes and help me deliver the feature.
I went with Go but didn't decide to use PostgreSQL, I opted for the DynamoDB which fits well in the serverless decision and also takes care of some performance issues.

The API consistents of three endpoints to make a reservation, delete a reservation, and check the room hours. The API structure doesn't look similar to a common one that follows
DDD standards, that is because I prefer to set up in an elegant way to work with serverless APIs.

To take care of the infrastructure part, I'm using the SLS framework which allows me a simple and elegant solution to control my resources.

Why serverless? My first thought was to use EC2 with some Containers but I will need to care but the scalability issue and also spend more time configuring the infrastructure part.
Going serverless with Lambdas + API Gateway, the cloud will help me take care of part of the responsibility and without spending too much time setting up environments.

My regret of using the SLS framework is that my idea in the first moment was to use ElasticCache with Redis to put an extra layer when checking for the room availabilities. Unfortonelly, the setup is more complicated than I imagined. In that case, using TF or Cloudformation could be a better option.


# CI/CD

Using git pipelines is a great way to have an elegant solution to deploy applications. As I'm using Github, I opted for Actions.
How the pipe works?

1. every time a change goes to the Main Branch, it will trigger the Action.
2. The action first builds the project ( I'm also using make to abstract some system stuff)
3. The action calls the deploy functionality from the SLS framework
4. With the secrets well setup, the projects go to the AWS

# Improvements

1. DynamoDB will handle a lot of scalability issues, but when including the data, I would certainly send a message to a queue and properly handle that to a cache layer.
2. The endpoint to see if the room is available basically returns all the date times. I think another endpoint using a similar logic that I have inside the put to return a bool if the date is available on a certain date is a good idea.
3. Improve the errors and checks when hitting the APIs
4. There is no security in the API, it needs an auth layer.
5. I wrote just unit tests, integrations tests are needed.