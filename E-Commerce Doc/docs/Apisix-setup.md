---
sidebar_position: 2
---

# Apisix Configuration

1. Launch the `docker-apisix and dashboard` in your docker

2. Go to the apisix dashboard by clicking the docker port 9000

3. Login with username `admin` and password `admin`

4. To create the gateway, first need to create an `Upstream` then create the `Route`

5. `Upstream` is where your backend locate, the purpose of create is for easier management on the backend and reusable if frontend required

6. `Route` is where the frontend will use to navigate to the backend

7. To create the `Upstream`, go to it and click create, then fill up the red rectangle box, **Host** is where your backend ip, **Port** is which port your backend running on(if applicable)

![apisix-upstream](/img/apisix-upstream.png)

:::info
To use **localhost ip** put the Host ip as below image.
![localhost](/img/keycloak-localhost.png)
:::

6. Then create the `Route`, fill up the required field as per the red box, **Host** is where your frontend ip

![apisix-route](/img/apisix-route.jpeg)

7. After that, choose the correct `Upstream` 

8. Remember to enable the CORS in the plugin section for cross platform communication

7. That all to configure the gateway