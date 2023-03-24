/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './src/**/*.{js,jsx}',
    './public/**/*.html',
    'node_modules/flowbite-react/**/*.{js,jsx,ts,tsx}'
  ],
  theme: {
    extend: {},
    colors : {
      'white':'#FFFFFF',
      'primary':'#0073FF',
      'primary-hover':'#095DC3',
      'secondary' : '#FFD600',
      'dark' : '#34364A',
      'mute' : '#656565',
      'base' : '#F6F8FD'
    },
    fontFamily: {
      inter: ['Inter', 'sans-serif']
    },
  },
  plugins: [
    require('flowbite/plugin')
  ],
}
