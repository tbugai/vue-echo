const path = require('path');

module.exports = {
  outputDir: path.resolve(__dirname, '../dist/public'),
  assetsDir: 'assets',
  lintOnSave: false,
  devServer: {
    port: process.env.FRONTEND_PORT,
    proxy: {
      '/api': {
        target: 'http://localhost:' + process.env.BACKEND_PORT,
        ws: true,
        changeOrigin: true
      }
    }
  }
}
