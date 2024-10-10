/** @type {import('tailwindcss').Config} */
export default {
    content: ['./src/**/*.{html,js,svelte,ts}'],  theme: {
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

