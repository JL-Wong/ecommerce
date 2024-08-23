---
sidebar_position: 1
---

# Keycloak Create Realm

1. Start your keycloak in docker, go to it dashboard by clicking on the docker Port
2. Login with username `admin` and password `admin`
3. First create realm

![create-realm](/img/create-realm.png)

5. Input the realm name and create
6. Next create `Client`
7. Ensure the `Client type` is `OpenID Connect`
8. Input the name and next
9. Make sure the `CLient authentication` is `on`, tick the `Service accounts roles` and next
10. Under `Valid redirect URIs` put in your web URL like **http://example.com/***
11. Then under `Web origins` put in the request url(the url that request to keycloak, example my frontend is **http://localhost:52000**, so will put that into it) or * to allow all so that CORS is allowed
12. Save
