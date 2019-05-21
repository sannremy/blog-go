const path = require('path');
const merge = require('webpack-merge');
const common = require('./webpack.common.js');

const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const LiveReloadPlugin = require('webpack-livereload-plugin');

module.exports = merge(common, {
  mode: 'development',
  output: {
    filename: '[name]-dev.js',
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: '[name]-dev.css',
    }),
    new LiveReloadPlugin()
  ]
});
