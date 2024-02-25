import resolve from '@rollup/plugin-node-resolve';
import commonjs from '@rollup/plugin-commonjs';

export default {
  input: 'src/main.js',
  output: {
    format: 'iife',
    file: 'public/bundle.js',
  },
  plugins: [resolve(), commonjs()],
};