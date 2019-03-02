const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

// icon svg from https://www.flaticon.com/authors/pixel-perfect

module.exports = {
  entry: './src/static/scripts/main.js',
  output: {
    filename: 'main.js',
    path: path.resolve(__dirname, 'dist')
  },
  plugins: [
    new MiniCssExtractPlugin()
  ],
  module: {
    rules: [
      // Use of ES6 class 'this'
      {
        test: /\.m?js$/,
        exclude: /(node_modules|bower_components)/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: [
              '@babel/preset-env'
            ],
            plugins: [
              '@babel/plugin-proposal-class-properties'
            ]
          }
        }
      },
      // SCSS transpiler
      {
        test: /\.(sa|sc|c)ss$/,
        use: [
          MiniCssExtractPlugin.loader,
          'css-loader',
          'sass-loader',
        ],
      },
      // Copy images
      {
        test: /\.(png|svg|jpg|gif)$/,
        use: [
          {
            loader: 'file-loader',
            options: {
              name: 'assets/images/[name]-[hash].[ext]',
            }
          }
        ]
      },
      // Copy fonts
      {
        test: /\.(woff|woff2|eot|ttf|otf)$/,
        use: [
          {
            loader: 'file-loader',
            options: {
              name: 'assets/fonts/[name]-[hash].[ext]',
            }
          }
        ]
      }
    ]
  }
};
