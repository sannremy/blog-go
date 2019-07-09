const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const CopyPlugin = require('copy-webpack-plugin');
const autoprefixer = require('autoprefixer');

module.exports = {
  entry: './web/static/scripts/main.js',
  output: {
    filename: '[name]-[hash].js',
    path: path.resolve(__dirname, '..', 'dist')
  },
  plugins: [
    new CopyPlugin([
      {
        from: './web/static/assets',
        to: 'assets',
      },
    ]),
  ],
  module: {
    rules: [
      // Babel config
      {
        test: /\.m?js$/,
        exclude: /(node_modules|bower_components)/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: ['@babel/preset-env'],
          },
        },
      },
      // SCSS transpiler
      {
        test: /\.(sa|sc|c)ss$/,
        use: [
          MiniCssExtractPlugin.loader,
          'css-loader',
          {
            loader: 'postcss-loader',
            options: {
              ident: 'postcss',
              plugins: [
                autoprefixer({
                  'browsers': ['> 1%', 'last 2 versions']
                }),
              ]
            }
          },
          'sass-loader',
        ],
      },
    ]
  }
};
