# Beego-LineBot: using Beego framework as LineBot handler
Can try it first.
<img src="http://qr-official.line.me/L/93By7CZ7co.png">
<a href="https://line.me/R/ti/p/%40orx4083y"><img height="36" border="0" alt="加入好友" src="https://scdn.line-apps.com/n/line_add_friends/btn/zh-Hant.png"></a>

# Installation and Usage
### 1.Deploy on your Heroku
<a href="https://heroku.com/deploy">
  <img src="https://www.herokucdn.com/deploy/button.svg" alt="Deploy">
</a>
### 2.Create LineBot account
If you don't have one, can create from [here](https://business.line.me/zh-hant/).
After create account, don't forget to open the Bot feature.

### 3.Setting LineBot Channel information
- `Webhook URL`: https://{YOUR_HEROKU_APP_URL}/callback

### 3.Setting Heroku Config Variables
You will need below message to run this app:
- `ChannelAccessToken` : You can find it in LineBot dashboard
- `ChannelSecret` : You can find it in LineBot dashboard
- `GoogleMapNearbySearchKey` : You can get it from [Google Places API](https://developers.google.com/places/)
