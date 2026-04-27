import Image from 'next/image';
import { MDXRemote } from 'next-mdx-remote/rsc';
import { Link } from '@/lib/i18n/navigation';

const components = {
  h2: (props: React.HTMLAttributes<HTMLHeadingElement>) => (
    <h2 className="heading-section text-2xl md:text-3xl text-petrol mt-12 mb-4" {...props} />
  ),
  h3: (props: React.HTMLAttributes<HTMLHeadingElement>) => (
    <h3 className="heading-section text-xl md:text-2xl text-petrol mt-10 mb-3" {...props} />
  ),
  p: (props: React.HTMLAttributes<HTMLParagraphElement>) => (
    <p className="text-petrol/85 leading-relaxed text-lg my-5" {...props} />
  ),
  ul: (props: React.HTMLAttributes<HTMLUListElement>) => (
    <ul className="list-disc pl-6 space-y-2 text-petrol/85 my-5" {...props} />
  ),
  ol: (props: React.HTMLAttributes<HTMLOListElement>) => (
    <ol className="list-decimal pl-6 space-y-2 text-petrol/85 my-5" {...props} />
  ),
  blockquote: (props: React.HTMLAttributes<HTMLQuoteElement>) => (
    <blockquote
      className="border-l-2 border-gold pl-6 py-2 my-8 italic text-petrol/85 text-xl"
      {...props}
    />
  ),
  a: ({ href = '#', ...props }: React.AnchorHTMLAttributes<HTMLAnchorElement>) => {
    const isInternal = href.startsWith('/');
    if (isInternal) {
      return (
        <Link
          href={href as never}
          className="text-petrol underline decoration-gold underline-offset-4 hover:text-gold"
          {...(props as object)}
        />
      );
    }
    return (
      <a
        href={href}
        target="_blank"
        rel="noreferrer"
        className="text-petrol underline decoration-gold underline-offset-4 hover:text-gold"
        {...props}
      />
    );
  },
  strong: (props: React.HTMLAttributes<HTMLElement>) => <strong className="text-petrol font-semibold" {...props} />,
  img: ({ src = '', alt = '' }: React.ImgHTMLAttributes<HTMLImageElement>) => (
    <figure className="my-10 -mx-4 md:mx-0">
      <Image
        src={typeof src === 'string' ? src : ''}
        alt={alt}
        width={1024}
        height={1024}
        sizes="(min-width: 768px) 720px, 100vw"
        className="w-full h-auto"
      />
      {alt && (
        <figcaption className="text-petrol/60 text-sm mt-3 italic px-4 md:px-0">
          {alt}
        </figcaption>
      )}
    </figure>
  ),
};

export function MdxContent({ source }: { source: string }) {
  return <MDXRemote source={source} components={components} />;
}
