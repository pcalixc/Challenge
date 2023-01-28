module.exports = {
  plugins: {
    tailwindcss: {},
    autoprefixer: {}
  },
  webpack: {
    configure: {
      experiments: {
        topLevelAwait: true,
      },
    },
  }
}
