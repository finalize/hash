const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");

module.exports = {
  mode: "development",
  entry: "./src/index.tsx",
  output: {
    filename: "bundle.js",
    path: __dirname + "/dist"
  },

  devtool: "source-map",

  resolve: {
    extensions: [".jsx", ".ts", ".tsx", ".js", ".json"]
  },

  module: {
    rules: [
      {
        test: /\.(ts|tsx)$/,
        exclude: /node_modules/,
        use: {
          loader: "ts-loader"
        }
      },
      {
        test: /\.(gif|png|jpg|eot|wof|woff|woff2|ttf|svg)$/,
        loader: "url-loader"
      }
    ]
  },
  devServer: {
    openPage: __dirname + "/dist/index.html", //自動で指定したページを開く
    contentBase: path.join(__dirname, "public"), // HTML等コンテンツのルートディレクトリ
    watchContentBase: true, //コンテンツの変更監視をする
    historyApiFallback: true,
    port: 3000
  },
  plugins: [
    //以下追記
    new HtmlWebpackPlugin({
      template: __dirname + "/dist/index.html",
      filename: "index.html"
    })
  ]
};
