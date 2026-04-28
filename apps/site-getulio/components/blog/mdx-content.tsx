import { MDXRemote } from 'next-mdx-remote/rsc';

const components = {
  h2: (props: React.HTMLAttributes<HTMLHeadingElement>) => (
    <h2 className="heading-section text-2xl md:text-3xl text-ink mt-12 mb-4" {...props} />
  ),
  h3: (props: React.HTMLAttributes<HTMLHeadingElement>) => (
    <h3 className="heading-section text-xl md:text-2xl text-ink mt-10 mb-3" {...props} />
  ),
  p: (props: React.HTMLAttributes<HTMLParagraphElement>) => (
    <p className="font-serif text-ink-soft leading-relaxed text-lg my-5" {...props} />
  ),
  ul: (props: React.HTMLAttributes<HTMLUListElement>) => (
    <ul className="list-disc pl-6 space-y-2 font-serif text-ink-soft my-5" {...props} />
  ),
  ol: (props: React.HTMLAttributes<HTMLOListElement>) => (
    <ol className="list-decimal pl-6 space-y-2 font-serif text-ink-soft my-5" {...props} />
  ),
  blockquote: (props: React.HTMLAttributes<HTMLQuoteElement>) => (
    <blockquote
      className="border-l-2 border-bordo pl-6 py-2 my-8 italic font-serif text-ink-soft text-xl"
      {...props}
    />
  ),
  a: ({ href = '#', ...props }: React.AnchorHTMLAttributes<HTMLAnchorElement>) => (
    <a
      href={href}
      target={href.startsWith('http') ? '_blank' : undefined}
      rel={href.startsWith('http') ? 'noreferrer' : undefined}
      className="text-ink underline decoration-bordo underline-offset-4 hover:text-bordo"
      {...props}
    />
  ),
  strong: (props: React.HTMLAttributes<HTMLElement>) => (
    <strong className="text-ink font-semibold" {...props} />
  ),
  // Imagens vêm com URL absoluto reescrito (rewriteAssetPaths) apontando para
  // plenyasaude.com.br — servimos via <img> nativo (sem next/image) para evitar
  // dependência de remotePatterns + tirar custo de otimização extra: a Plenya
  // já serve as imagens otimizadas pelo seu próprio next/image+CDN.
  img: ({ src = '', alt = '' }: React.ImgHTMLAttributes<HTMLImageElement>) => (
    <figure className="my-10 -mx-4 md:mx-0">
      {/* eslint-disable-next-line @next/next/no-img-element */}
      <img src={typeof src === 'string' ? src : ''} alt={alt} loading="lazy" className="w-full h-auto" />
      {alt && (
        <figcaption className="font-sans text-ink-muted text-sm mt-3 italic px-4 md:px-0">
          {alt}
        </figcaption>
      )}
    </figure>
  ),
};

export function MdxContent({ source }: { source: string }) {
  return <MDXRemote source={source} components={components} />;
}
