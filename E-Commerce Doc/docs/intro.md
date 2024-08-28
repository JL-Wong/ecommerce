---
sidebar_position: 1
---

# E-Commerce Project Intro

## Getting Started

The Objective of the project:
- to handle Order and Product management on web app
- to simplify the SQL desktop app on E-Commerce handling

## What you'll need

### 1. Docker

- Download the [Docker](https://www.docker.com/), just follow the instruction as per website.

### 2. Apache Apisix

First clone the apisix-docker repository:

```bash
git clone https://github.com/apache/apisix-docker.git
cd apisix-docker/example
```
Now, use `docker-compose` to start APISIX.

```bach
docker-compose -p docker-apisix up -d
```

Above will start the apisix in docker. In order to use the dashboard of the apisix, perform the below steps:

First clone the apisix-dashboard repository:

```bash
git clone https://github.com/apache/apisix-dashboard.git
```

Locate the conf.yaml file in the api folder: `xxxx\apisix-dashboard\api\conf`

First comment out the **allow_list** to allow any ip to access

![apisix-dashboard-conf1](/img/apisix-dashboard-conf1.png)

Then change the etcd's endpoint to your local ip address

![apisix-dashboard-conf](/img/apisix-dashboard-conf.png)

After that, comment out the **oidc** part

![apisix-dashboard-conf2](/img/apisix-dashboard-conf2.png)

Then replace the `CONFIG_FILE` with the path
```bash
docker pull apache/apisix-dashboard
docker run -d --name dashboard \
           -p 9000:9000        \
           -v <CONFIG_FILE>:/usr/local/apisix-dashboard/conf/conf.yaml \
           apache/apisix-dashboard
```

Now your docker should have something like below:
![Apisix](/img/apisix.png)

### 3. Keycloak
Pull image from docker hub

```bash
docker run -d -p 8080:8080 -e KEYCLOAK_ADMIN=admin -e KEYCLOAK_ADMIN_PASSWORD=admin quay.io/keycloak/keycloak:25.0.4 start-dev
```
Then you should have a keycloak in your docker like below:
![Keycloak](/img/keycloak.png)

## What to do to achieve

1. To generate a **General JSON** that would be use when SQL acc fetch data from the web app

    1. The **General JSON** should also be reusable for all the platform

2. Design a suitable **DB** to store big data, must consider the performance when query