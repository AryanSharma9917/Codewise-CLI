import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'Codewise CLI',
  description: 'Documentation for the Codewise DevOps command-line utility',
  base: '/Codewise-CLI/',
  cleanUrls: true,
  lastUpdated: true,
  head: [
    ['link', { rel: 'icon', href: '/Codewise-CLI/logo.png' }],
    ['meta', { name: 'theme-color', content: '#10b981' }]
  ],
  themeConfig: {
    logo: '/logo.png',
    siteTitle: 'Codewise CLI',
    search: { provider: 'local' },
    nav: [
      { text: 'Guide', link: '/guide/quick-start' },
      { text: 'Commands', link: '/commands/' },
      { text: 'Workflows', link: '/workflows/docker-to-kubernetes' },
      { text: 'Troubleshooting', link: '/troubleshooting' }
    ],
    sidebar: {
      '/guide/': [
        {
          text: 'Guide',
          items: [
            { text: 'Installation', link: '/guide/installation' },
            { text: 'Quick start', link: '/guide/quick-start' },
            { text: 'Configuration', link: '/guide/configuration' },
            { text: 'How Codewise works', link: '/guide/architecture' }
          ]
        }
      ],
      '/commands/': [
        {
          text: 'Command reference',
          items: [
            { text: 'Overview', link: '/commands/' },
            { text: 'Config and environments', link: '/commands/config-env' },
            { text: 'Docker', link: '/commands/docker' },
            { text: 'Kubernetes and Helm', link: '/commands/kubernetes-helm' },
            { text: 'Deploy', link: '/commands/deploy' },
            { text: 'Encode', link: '/commands/encode' },
            { text: 'Init and templates', link: '/commands/generators' }
          ]
        }
      ],
      '/workflows/': [
        {
          text: 'Workflows',
          items: [
            { text: 'Docker to Kubernetes', link: '/workflows/docker-to-kubernetes' },
            { text: 'Helm deployment', link: '/workflows/helm' },
            { text: 'GitOps with Argo CD', link: '/workflows/gitops' },
            { text: 'Preview environments', link: '/workflows/preview-environments' }
          ]
        }
      ]
    },
    socialLinks: [
      { icon: 'github', link: 'https://github.com/aryansharma9917/codewise-cli' }
    ],
    editLink: {
      pattern: 'https://github.com/aryansharma9917/codewise-cli/edit/main/website/:path',
      text: 'Edit this page on GitHub'
    },
    footer: {
      message: 'Released under the MIT License.',
      copyright: 'Copyright © 2026 Codewise CLI contributors'
    },
    outline: { level: [2, 3] }
  }
})
