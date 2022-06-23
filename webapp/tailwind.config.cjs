/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'class',
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        'primary-50': "#EBF8FF",
        'primary-100': "#BEE3F8",
        'primary-200': "#90CDF4",
        'primary-300': "#63B3ED",
        'primary-400': "#4299E1",
        'primary-500': "#3182CE",
        'primary-600': "#2B6CB0",
        'primary-700': "#2C5282",
        'primary-800': "#2A4365",
        'primary-900': "#1A365D",
      }
    }
  },
  plugins: []
};