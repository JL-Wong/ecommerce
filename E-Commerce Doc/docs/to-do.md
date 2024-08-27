# To Do List

1. Keycloak have `check_session_iframe`, based on documentation is used to check for the SSO session, can research on this to enhance the session control function

2. Research on how to store and query big data without affecting the performance, so method for further explore

    1. Columnar Databases

    2. Presto (or Trino)

3. Research on how to enable different access token when login from the desktop app and also limit to one login per user only(no concurrent login). Below is some idea

    1. use different client for different device

4. Research on the scope for user management, when the admin can assign different scope to the users added 

5. Research on how to assign store to particular user. Below is some idea

    1. inside the group create sub-group

6. Research on the NATS.io message queue, on how the load balancing work, others feature that can handle large API call



:::warning
Keycloak js adaptor is deprecated and will be removed in the future, if wish to use then please consult with ShengLi 
- https://www.npmjs.com/package/keycloak-connect
- https://www.keycloak.org/2022/02/adapter-deprecation
:::

