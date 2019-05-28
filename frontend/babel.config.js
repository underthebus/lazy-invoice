const plugins = [
  'react-hot-loader/babel',
];

const presets = [
  '@babel/preset-react',
  [
    '@babel/env',
    {
      useBuiltIns: 'usage',
      corejs: '3.1.2',
    },

  ],
];

module.exports = { presets, plugins };
