# Beego-LineBot: using Beego framework as LineBot handler
Try it by yourself : 
<a href="http://qr-official.line.me/M/93By7CZ7co.png" target="_blank"><img height="36" border="0" alt="加入好友" src="https://scdn.line-apps.com/n/line_add_friends/btn/zh-Hant.png"></a>

# Installation and Usage
### 1. Deploy on your Heroku
<a href="https://heroku.com/deploy">
  <img src="https://www.herokucdn.com/deploy/button.svg" alt="Deploy">
</a>

### 2. Create LineBot account
You can create one from [here](https://business.line.me/zh-hant/)


### 3. Setting LineBot Channel information
- `Webhook URL`: https://{YOUR_HEROKU_APP_URL}/callback

### 4. Setting Heroku Config Variables
Go to Your Heroku APP -> Settings -> Config Variables -> Reveal Config Vars
Add below variables:
- `ChannelAccessToken` : You can find it in LineBot dashboard
- `ChannelSecret` : You can find it in LineBot dashboard
- `GoogleMapNearbySearchKey` : You can get it from [Google Places API](https://developers.google.com/places/)
