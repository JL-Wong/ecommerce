const express = require('express')
const axios = require('axios')
const cors = require('cors')
const bodyParser = require('body-parser')
const jwt = require('jsonwebtoken')
const keycloakConfig = require('./keycloak.json')

const app = express()
app.use(cors())
app.use(bodyParser.json())
app.use(bodyParser.urlencoded({ extended: true }))

let refreshToken = ''
const sessionStore = {}; // In-memory store for active sessions


app.post('/google-login', async (req, res) => {
  const baseUrl = `${keycloakConfig['auth-server-url']}realms/E-Commerce/protocol/openid-connect/auth`
  const clientID = 'apisix'
  const redirectUrl = 'http://localhost:52000/home'

  const authUrl = `${baseUrl}?client_id=${clientID}&redirect_uri=${redirectUrl}&response_type=code&scope=openid email profile`;
  console.log(authUrl)
  //send to frontend as JSON format
  res.json({ authUrl: authUrl });

})

//tis is use for the add user function
let access_token1

app.post('/exchange', async (req, res) => {
  const { code } = req.body;
  // console.log(code)
  const clientID = 'apisix'
  const clientSecret = 't8jdxl6oUfvsypAyGiZOeUseGlk0tjrF'
  const redirectUrl = 'http://localhost:52000/home'
  const authCode = code

  const url = `${keycloakConfig['auth-server-url']}realms/E-Commerce/protocol/openid-connect/token`;
  const headers = {
    'Content-Type': 'application/x-www-form-urlencoded'
  };
  const data = new URLSearchParams({
    grant_type: 'authorization_code',
    code: authCode,
    client_id: clientID,
    client_secret: clientSecret,
    redirect_uri: redirectUrl
  });

  try {
    const response = await axios.post(url, data, { headers });
    if (response.status === 200) {
      const { access_token, id_token, refresh_token } = response.data;
      
      //Update the refresh token 
      //   refreshToken = refresh_token
      //   console.log(response.data)
      //   const decoded = jwt.decode(access_token);
      // console.log(decoded)
      // const role = decoded['realm_access']['roles'].toString()
      // console.log(role.toString())
      access_token1 = access_token

      const access = checkUserAccess(access_token)
      console.log(access_token)
      res.json({ access_token, id_token, access});
    } else {
      res.status(response.status).send('Failed to get token');
    }
  } catch (error) {
    console.error('Error fetching token:', error);
    res.status(500).send('Internal Server Error');
  }
})

app.post('/logout', async (req, res) => {
  const { id_token } = req.body;
  console.log(id_token)
  const url = `${keycloakConfig['auth-server-url']}realms/E-Commerce/protocol/openid-connect/logout`;
  // console.log(url)
  const headers = {
    'Content-Type': 'application/x-www-form-urlencoded'
  };
  const data = new URLSearchParams({
    id_token_hint: id_token,

  });

  try {
    const response = await axios.post(url, data, { headers });
    if (response.status === 200) {
      console.log("logout success")
      res.json({ "msg": "Logout success" })
    } else {
      res.status(response.status).send('Failed to logout');
    }
  } catch (error) {
    console.error('Error fetching logout:', error);
    res.status(500).send('Internal Server Error');
  }
});

app.post('/add', async (req, res) => {
  const { username, email, realmRoles } = req.body;
  
  const decoded = jwt.decode(access_token1)
  // console.log(decoded)
  const getCompany = decoded['company']
  // console.log(getCompany)
  // console.log(access_token1)
  const url = `${keycloakConfig['auth-server-url']}admin/realms/E-Commerce/users`;
  const roleAssignmentUrl = `${keycloakConfig['auth-server-url']}admin/realms/E-Commerce/users/{userId}/role-mappings/realm`;
  const headers = {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${access_token1}`,
  };
  const data = {
    username: username,
    email: email,
    enabled: true,
    groups: getCompany,
  }

  try {
    const response = await axios.post(url, data, { headers });
    //there will be no response data, only status indicate success of fail
    if (response.status === 201) {
      
      const userID = await getSingleUserID(email)
      // console.log(`this is the add function userID ${userID}`)
      const rolesData = [{
        "id":realmRoles == 'Packer' ? keycloakConfig.Packer : keycloakConfig.Admin,
        "name":realmRoles
      }]
      // console.log(rolesData)
      const roleResponse = await axios.post(roleAssignmentUrl.replace('{userId}', userID), rolesData, { headers });
      // console.log(roleResponse.status)
      if (roleResponse.status === 204) { // Role assignment successful
        console.log("User created and roles assigned");
        res.json("User created and roles assigned");
      } else {
        res.status(roleResponse.status).send('Failed to assign roles');
      }
    } else {
      res.status(response.status).send('Failed to create user');
    }
  } catch (error) {
    console.error('Error:', error);
    res.status(500).send('Internal Server Error');
  }


});

app.get('/get-user',async(req,res)=>{
  const authHeader = req.headers['authorization']
  const token = authHeader.split(' ')[1]
  const decoded = jwt.decode(token)
  console.log(decoded)
  const getCompany = decoded['company']
  const company = getCompany.toString().replace(/\//g,'')


  const url = `${keycloakConfig['auth-server-url']}admin/realms/E-Commerce/groups?search=${company}`
  const headers = {
    'Authorization': `Bearer ${token}`,
  };

  try{
    const response = await axios.get(url, { headers });
    // console.log(response)
    const id = response.data[0]['id']
    // console.log(id)
    const url2 = `${keycloakConfig['auth-server-url']}admin/realms/E-Commerce/groups/${id}/members`
    if(!id){
      res.status(response.status).send('Failed to get users');
    }

    const response2 = await axios.get(url2,{ headers })
    // console.log(response2.data)
    const users = response2.data.map(({ id,username, email }) => ({ id,username, email }));
    res.json(users)

  }catch(error){
    console.log(error)
  }
})

app.put('/update-user',async(req,res)=>{
  const { id, username, email } = req.body;
  
  //this header include the Bearer so need to split
  const authHeader = req.headers['authorization']
  const token = authHeader.split(' ')[1]
  const url = `${keycloakConfig['auth-server-url']}admin/realms/E-Commerce/users/${id}`
  const headers = {
    'Authorization': `Bearer ${token}`,
  };
  const data = {
    email: email,
  }

  try {
    const response = await axios.put(url, data, { headers })
    if (response.status === 204) {
      // console.log("User updated")
      res.json('User updated')
    } else {
      res.status(response.status).send('Failed to update');
    }
  } catch (error) {
    console.log(error)
  }

})

app.delete('/delete-user',async(req,res)=>{
  const { id } = req.body;
  
  //this header include the Bearer so need to split
  const authHeader = req.headers['authorization']
  const token = authHeader.split(' ')[1]
  const url = `${keycloakConfig['auth-server-url']}admin/realms/E-Commerce/users/${id}`

  const headers = {
    'Authorization': `Bearer ${token}`,
  };

  try {
    const response = await axios.delete(url,{ headers })
    if (response.status === 204) {
      console.log("User deleted")
      res.json('User deleted')
    } else {
      res.status(response.status).send('Failed to delete');
    }
  } catch (error) {
    console.log(error)
  }

})

async function getSingleUserID(email){
  const url = `${keycloakConfig['auth-server-url']}admin/realms/E-Commerce/users?search=${email}`
  const headers ={
    'Authorization': `Bearer ${access_token1}`,
  }
  try {
    const response = await axios.get(url, { headers })

    if(response.status == 200){
      const userID = response.data[0]['id']
      // console.log(userID)
      return userID
    }
  } catch (error) {
    console.error('Fail to get user info:', error)
  }
}

function checkUserAccess(token){
  const decoded = jwt.decode(token)
  const getRoles = decoded['realm_access'].roles
  const checkRole = getRoles.includes('Admin')
  
  return checkRole
  
}

PORT = 3000
app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}`)
});