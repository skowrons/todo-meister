// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Todo Meister',
  tagline: 'Make TODO: fun again!',

  // TODO: set production url
  url: 'https://your-docusaurus-test-site.com',
  baseUrl: '/',

  organizationName: 'skowrons',
  projectName: 'todo-meister',

  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',

  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
        },
        blog: {
          showReadingTime: true,
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      // TODO: replace with own social card
      image: 'img/docusaurus-social-card.jpg',
      navbar: {
        title: 'Todo Meister',
        items: [
          {
            type: 'docSidebar',
            sidebarId: 'docsSidebar',
            position: 'left',
            label: 'Documentation',
          },
          {
            href: 'https://github.com/skowrons/todo-meister/releases',
            label: 'Version History',
            position: 'left',
          },
          {
            href: 'https://github.com/skowrons/todo-meister',
            label: 'GitHub',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        links: [
          {
            title: 'Docs',
            items: [
              {
                label: 'Intro',
                to: '/docs/intro',
              },
              {
                label: 'Getting Started',
                to: '/docs/getting-started',
              },
            ],
          },
          {
            title: 'More',
            items: [
              {
                label: 'Maximilian Skowron',
                href: "https://www.skowron.one",
              },
              {
                label: 'GitHub',
                href: 'https://github.com/skowrons/todo-meister',
              },
            ],
          },
        ],
        copyright: `Copyright © ${new Date().getFullYear()} Maximilian Skowron`,
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
      },
    }),
};

module.exports = config;
