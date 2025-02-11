export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'maaddii-food-recipe',
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: 'A food recipe application' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
    'assets/css/tailwind.css' // Add global Tailwind CSS
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    // Add any necessary plugins here
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/tailwindcss
    '@nuxtjs/tailwindcss',
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // Add any necessary modules here
  ],

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
    // Add any necessary build configurations here
  },

  publicRuntimeConfig: {
    apiUrl: process.env.API_URL || 'http://localhost:8080' // Set the API endpoint
  }
}
