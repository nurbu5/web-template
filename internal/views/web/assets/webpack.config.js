const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const CssMinimizerPlugin = require('css-minimizer-webpack-plugin'); // Import the plugin

module.exports = {
  mode: 'production',
  entry: './src/js/app.js',
  output: {
    // TODO: Add a hash to the filename
    // filename: 'js/bundle.[contenthash].js',  // Output JS to dist/js
    filename: 'js/bundle.js',
    path: path.resolve(__dirname, 'dist'),   // Output directory
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
      // TODO: Add a hash to the filename
      // filename: 'css/style.[contenthash].css',  // Extract CSS to dist/css
      filename: 'css/style.css',
    }),
  ],
  optimization: {
    minimize: true,  // Minify the output
    minimizer: [
      '...',
      new CssMinimizerPlugin(),  // Add this for CSS minification
    ],
  },
};

