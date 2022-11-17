const path = require("path")
const webpack = require("webpack")

module.exports = {
  entry: {
    index: "./sender.js",
  },
  output: {
    filename: "sender.bundle.js",
    path: path.resolve(__dirname),
  },
  plugins: [
    new webpack.DefinePlugin({
      ESP_EYE_IP: JSON.stringify(process.env.ESP_EYE_IP),
    }),
  ],
}