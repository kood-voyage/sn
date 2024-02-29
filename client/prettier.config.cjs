module.exports = {
	useTabs: true,
	singleQuote: true,
	trailingComma: "none",
	printWidth: 100,
	tailwindConfig: './styles/tailwind.config.js',
	overrides: [
		{
			files: "*.svelte",
			options: {
				parser: "svelte"
			}
		}
	],
	  plugins: [
    "prettier-plugin-tailwindcss" // MUST come last
  ]
}
