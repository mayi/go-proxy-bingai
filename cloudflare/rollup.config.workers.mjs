import compiler from '@ampproject/rollup-plugin-closure-compiler';

// rollup.config.mjs
// ---cut-start---
/** @type {import('rollup').RollupOptions} */
// ---cut-end---
export default {
    input: 'cloudflare/worker.js',
    output: [
        {
            file: 'worker.js',
            format: 'es',
        }
    ],
    plugins:[
        obfuscator({
			options: {
				// Your javascript-obfuscator options here
				// See what's allowed: https://github.com/javascript-obfuscator/javascript-obfuscator
                splitStrings: true,
                splitStringsChunkLength: 5,
			},
		}),
        compiler({
            compilation_level:"ADVANCED"
        }),
    ]
};