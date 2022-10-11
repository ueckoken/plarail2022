const path = require("path")
const webpack = require("webpack")

module.exports = {
  entry: {
    skyway_receiver: "./skyway_receiver.js",
    momo_sender: "./momo_sender.js",
  },
  output: {
    filename: "[name]_.js",
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
