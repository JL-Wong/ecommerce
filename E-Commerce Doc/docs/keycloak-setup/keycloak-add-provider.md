---
sidebar_position: 2
---

# Keycloak Add Provider

## Configure Google

1. First go to [Google Cloud Console](https://console.cloud.google.com)

2. Then go to `APIs and services` --> `Crendentials` --> `Create Credentials`

    1. Choose `OAuth client ID`

    2. Select your application type as `Web application`

    3. Give a name

    4. Under the `Authorised redirect URIs`, add in the URI that Keycloak provide

3. Then configure the `OAuth consent screen`, just follow the instruction to proceed

4. Select the scopes of `email, profile and openid`

5. Add in your test users account

:::note
Remember to create a project first in the Google
:::

## Configure Keycloak's Provider

1. Go the the `Identity provides` in Keycloak

![identity-provider](/img/keycloak-add-provider.png)

2. Then add in the `Google` provider

3. The Client ID and Client Secret is from the Provider which is Google after completed the **Configure Google** steps

4. Copy the `Redirect URI` given by Keycloak and paste it into the Google `Authorised redirect URIs`

## After add in Provider

1. Go to `Authentication`

2. Click into the `browser`, change the `Identity Provider Redirector` to `Required`, then go into it setting under `Default Identity Provider` put in `google`

3. Then make a copy of the `browser` or you can create a flow by yourself

4. Inside the new flow, put in the required flow as per below image:

![keycloak-flow](/img/keycloak-flow.png)

5. Then at the `User session count limiter` go into it setting to set the concurrent user

6. After that go back to the `Identity Provider`, add in the own created flow into it

![provider-flow](/img/provider-flow.png)




