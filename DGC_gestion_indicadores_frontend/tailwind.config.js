/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      {
        myLigthTheme: {
          "primary": "#F35353",
          "secondary": "#B8B8B8",
          "accent": "#B7A38F",
          "neutral": "#87BED4",
          "base-100": "#F0F0F0",
          "info": "#509BCE",
          "success": "#41DC91",
          "warning": "#F4C015",
          "error": "#E74B72",
        },
        dark: {
          "primary": "#990B0B",
          "secondary": "#525252",
          "accent": "#705C48",
          "neutral": "#316E87",
          "base-100": "#0A0A0A",
          "info": "#2D719F",
          "success": "#167949",
          "warning": "#E9B50C",
          "error": "#C51B46",
        },
      },
    ],
    darkTheme: "dark",
    base: true
  },
}

