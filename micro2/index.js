// screen2-service/index.js
const express = require('express');
const cors = require('cors');
const app = express();
const port = 3002;
// const host = '192.168.31.85';
const host = '127.0.0.1';

app.use(cors());
app.use(express.json());

app.get('/screen2', (req, res) => {
  res.json({ message: 'This is data for Screen 2' });
});

app.post('/screen2', (req, res) => {
  const { email, password } = req.body;
  console.log('Received POST data:', req.body);

  // For demonstration, simply echoing the received data back in the response
  res.json({
    message: 'POST request to /screen2 received',
    data: {
      email,
      password
    }
  });
});


app.listen(port, host,() => {
  console.log(`Screen 2 service running at http://${host}:${port}`);
});
