const express = require('express')
const axios = require('axios')
const cors = require('cors')
const bodyParser = require('body-parser')
const jwt = require('jsonwebtoken')
const { Issuer, Strategy } = require('openid-client')
const passport = require('passport')
const path = require('path');
const expressSession = require('express-session')
const keycloakConfig = require('./keycloak.json');

const app = express()
app.use(cors())
app.use(bodyParser.json())
app.use(bodyParser.urlencoded({ extended: true }))

const memoryStore = new expressSession.MemoryStore();

app.use(
  expressSession({
    secret: 'my_secret',
    resave: false,
    saveUninitialized: true,
    store: memoryStore
  })
)


// Asynchronous function to set up OIDC with Keycloak
const setupOIDC = async () => {
  const keycloakIssuer = await Issuer.discover(`${keycloakConfig['auth-server-url']}/realms/${keycloakConfig.realm}`)

  console.log('Discovered issuer %s %O', keycloakIssuer.issuer, keycloakIssuer.metadata)
}


PORT = 3000
app.listen(PORT, () => {
  setupOIDC()
  console.log(`Server is running on port ${PORT}`)
});