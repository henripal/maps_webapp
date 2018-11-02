# Sample Basic Web App with Deep Learning Backend

See it in action [here](http://yourcityfrom.space/) !

Type in your address, town or location, the webapp gets a satellite picture of your city/neighborhood. Our deep learning algorithm then tries to guess its country by looking at the pixels. The algorithm doesn't generalize very well at the moment but I'm working on it!

## Background

This project was made as part of the [fast.ai course](https://course.fast.ai/), Fall 2018.

It's a simple, single page web application with a VueJS [frontend](./frontend), a [backend API](./backend) with the Google Maps Static API using Go, and a [backend deep learning API](./model_backend).

The deep learning API is a simple resnet34 trained on satellite images downloaded using a Golang tool. The training was done using the [fastai library](https://docs.fast.ai/). 

The design of the DL backend was inspired by @simonw's [cougar or not](https://github.com/simonw/cougar-or-not) app and uses [Starlette](https://www.starlette.io/) for async serving.

The entire app is served on a [Digital Ocean](http://digitalocean.com) docker droplet and is run using docker-compose.




