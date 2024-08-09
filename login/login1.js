const express = require('express');
const axios = require('axios')
const session = require('express-session');
const Keycloak = require('keycloak-connect');
const cors = require('cors');
const bodyParser = require('body-parser');

const app = express();
app.use(cors());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

const memoryStore = new session.MemoryStore();
const keycloak = new Keycloak({ store: memoryStore });


app.use(session({
  secret: 'mysecret',
  resave: false,
  saveUninitialized: true,
  store: memoryStore
}));



app.set( 'trust proxy', true );
// Middleware to protect routes
app.use(keycloak.middleware());

// Public route (no authentication required)
app.post('/google-login', async(req, res) => {
  const authUrl = keycloak.authUrl + '/realms/E-Commerce/protocol/openid-connect/auth?client_id=apisix&redirect_uri=http://localhost:52000/home&response_type=code&scope=openid email profile';
  console.log('doing...')
  res.json({ authUrl: authUrl });
  
});

// Protected route (requires authentication)
app.post('/add', keycloak.protect(), async (req, res) => {
  const { username, email } = req.body;
  const token = req.kauth.grant.access_token.content;
  const getCompany = token['company'];

  const url = 'http://localhost:8080/admin/realms/E-Commerce/users';
  const headers = {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${req.kauth.grant.access_token.token}`,
  };
  const data = {
    username: username,
    email: email,
    enabled: true,
    groups: getCompany
  };

  try {
    const response = await axios.post(url, data, { headers });
    res.status(200).json("User created");
  } catch (error) {
    console.error('Error creating user:', error);
    res.status(500).send('Internal Server Error');
  }
});

// Logout route
app.post('/logout', keycloak.protect(), (req, res) => {
  keycloak.logout(req.kauth.grant.access_token.token)
    .then(() => {
      res.json({ "msg": "Logout success" });
    })
    .catch((error) => {
      console.error('Error logging out:', error);
      res.status(500).send('Internal Server Error');
    });
});

// Example of protecting a route
app.get('/get-user', keycloak.protect(), async (req, res) => {
  const token = req.kauth.grant.access_token.content;
  const company = token['company'].replace(/\//g, '');

  try {
    const response = await axios.get(`http://localhost:8080/admin/realms/E-Commerce/groups?search=${company}`, {
      headers: { 'Authorization': `Bearer ${req.kauth.grant.access_token.token}` }
    });

    const groupId = response.data[0]?.id;
    if (!groupId) {
      return res.status(404).send('Group not found');
    }

    const usersResponse = await axios.get(`http://localhost:8080/admin/realms/E-Commerce/groups/${groupId}/members`, {
      headers: { 'Authorization': `Bearer ${req.kauth.grant.access_token.token}` }
    });

    const users = usersResponse.data.map(({ id, username, email }) => ({ id, username, email }));
    res.json(users);
  } catch (error) {
    console.error('Error fetching users:', error);
    res.status(500).send('Internal Server Error');
  }
});

PORT = 3000;
app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}`);
});
