/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./template/**/*.{templ, html}"],
  theme: {
    extend: {
      screens: {
        print: {raw: "print"},
        screen: {raw: "screen"}
      }
    },
  },
  plugins: [],
}

