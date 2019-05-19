module.exports = {
  root: true,
  env: {
    browser: true,
    node: true,
    jest: true,
  },
  globals: {
    $nuxt: true
  },
  parserOptions: {
    parser: 'babel-eslint',
  },
  extends: [
    'plugin:vue/recommended',
    'eslint:recommended',
    'plugin:prettier/recommended',
    'prettier/vue'
  ],
  plugins: ['vue'],
  rules: {
    'vue/component-name-in-template-casing': ['error', 'PascalCase'],
    'prettier/prettier': [
      'error',
      {
        semi: false,
        singleQuote: true,
        trailingComma: 'es5',
        printWidth: 80,
      },
    ],
  },
}
