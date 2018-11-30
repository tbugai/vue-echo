const path = require('path');

module.exports = {
  outputDir: path.resolve(__dirname, '../dist/public'),
  assetsDir: 'assets',
  lintOnSave: false,
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:8000',
        ws: true,
        changeOrigin: true
      }
    }
  }
}
