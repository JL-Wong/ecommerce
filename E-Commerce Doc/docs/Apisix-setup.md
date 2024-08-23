---
sidebar_position: 2
---

# Apisix Configuration

1. Launch the `docker-apisix and dashboard` in your docker
2. Go to the apisix dashboard by clicking the docker port 9000
3. Login with username `admin` and password `admin`
4. To create the gateway, first need to create an `Upstream` then create the `Route`
5. To create `Upstream` click `Create` fill up the required field

:::info
To use **localhost ip** put the Host ip as below image.
![localhost](/img/keycloak-localhost.png)
:::

6. Then create the route, fill up the required field

:::tip
Remember to enable the CORS in the plugin section for cross platform communication
:::

7. That all to configure the gateway