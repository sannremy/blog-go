const path = require('path');
const merge = require('webpack-merge');
const common = require('./webpack.common.js');

const NodemonPlugin = require('nodemon-webpack-plugin');
const LiveReloadPlugin = require('webpack-livereload-plugin');

module.exports = merge(common, {
  mode: 'development',
  plugins: [
    new NodemonPlugin({
      watch: path.resolve('./src'),
      ignore: ['./src/static', './src/views'],
      verbose: true,
      script: './src/app.js'
    }),
    new LiveReloadPlugin()
  ]
});
