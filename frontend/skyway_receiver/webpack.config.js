const path = require("path")
const webpack = require("webpack")

module.exports = {
  entry: {
    index: "./index.js",
  },
  output: {
    filename: "bundle.js",
    path: path.resolve(__dirname),
  },
  plugins: [
    new webpack.DefinePlugin({
      SW_WSURL: '"wss://webrtc.chofufes2021.gotti.dev/"',
      SKYWAY_APIKEY: JSON.stringify(process.env.SKYWAY_APIKEY),
      SKYWAY_DEBUG_LEVEL: "2",
      AYAME_SIGNALING_KEY: JSON.stringify(process.env.AYAME_SIGNALING_KEY),
      SENDER_TOKEN: JSON.stringify(process.env.SENDER_TOKEN),
    }),
  ],
}
