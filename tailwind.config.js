/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.html'],
  theme: {
    extend: {
      colors: {
        'primary': '#C778DD',
        'primary-light': '#c470db1a',
        'secondary': '#282C33',
        'tertiary': '#ABB2BF'
      
      },
      fontFamily: {
        'fira-code': ['Fira Code'],
      }
    },
  },
  plugins: [],
}

