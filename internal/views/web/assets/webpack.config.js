const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const CssMinimizerPlugin = require('css-minimizer-webpack-plugin'); // Import the plugin
const WebpackAssetsManifest = require('webpack-assets-manifest');

module.exports = {
  mode: 'production',
  entry: './src/js/app.js',
  output: {
    path: path.resolve(__dirname, 'dist'),   // Output directory
    filename: 'js/bundle.[contenthash].js',  // Output JS to dist/js
    publicPath: '/static/', // Ensure Webpack understands the path where the assets will be served
  },
  module: {
    rules: [
      {
        test: /\.scss$/, // Process SCSS files
        use: [
          MiniCssExtractPlugin.loader, // Extract CSS into separate files (for production)
          'css-loader', // Interpret @import and url() like import/require() and resolves them
          'sass-loader', // Compile SCSS to CSS
        ],
      },
    ],
  },
  plugins: [
    new CleanWebpackPlugin(),  // Clean the output directory
    new MiniCssExtractPlugin({
      filename: 'css/style.[contenthash].css',  // Extract CSS to dist/css
    }),
    new WebpackAssetsManifest({
      output: 'manifest.json', // Output manifest file
      publicPath: true, // Include the public path in the manifest
    }),
  ],
  optimization: {
    minimize: true,  // Minify the output
    minimizer: [
      '...',
      new CssMinimizerPlugin(),  // Add this for CSS minification
    ],
  },
  mode: 'production',
};

