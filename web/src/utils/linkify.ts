import linkify from 'linkifyjs/html'

export function linkifyHTMLText(text: string): string {
  return linkify(text, {
    defaultProtocol: 'https',
    attributes: {
      rel: 'noreferrer noopener',
    },
    className: 'text-blue hover:underline',
    nl2br: true,
  })
}
