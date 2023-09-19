export const module = {
    rules: [
        {
            test: /\.go$/,
            use: 'golang-wasm-async-loader',
        },
    ],
};
  