import linkify from 'linkifyjs/html'

export function linkifyHTMLText(text: string): string {
  return linkify(text, {
    defaultProtocol: 'https',
    attributes: {
      rel: 'noreferrer noopener',
      target: '_blank',
    },
    className: 'text-blue hover:underline',
    nl2br: true,
  })
}
