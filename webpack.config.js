module.exports = {
  entry: __dirname + '/assets/js/main.js',
  output: {
    path: __dirname + "/assets/js",
    filename: 'bundle.js'
  },
  module: {
    loaders: [
      {
        test: /.jsx?$/,
        loader: 'babel-loader',
        exclude: /node_modules/,
        query: {
          presets: ['es2015', 'react']
        }
      }
    ]
  }
}
