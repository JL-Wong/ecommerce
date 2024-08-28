const express = require('express');
const UAParser = require('ua-parser-js');

const app = express();

app.get('/login', (req, res) => {
    const parser = new UAParser();
    const ua = req.headers['user-agent'];
    const deviceInfo = parser.setUA(ua).getResult();

    // Extract platform details
    const platform = deviceInfo.os.name;
    const browser = deviceInfo.browser.name;
    const device = deviceInfo.device.type || 'Desktop';

    console.log(`Platform: ${platform}`);
    console.log(`Browser: ${browser}`);
    console.log(`Device: ${device}`);

    res.send(`Logged in from ${platform} on a ${device} using ${browser}`);
});

app.listen(3000, () => {
    console.log('Server is running on port 3000');
});
