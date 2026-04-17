import Image from 'next/image';

const tiles = [
  { src: '/images/lifestyle-movement.jpg', alt: 'Pessoa correndo em trilha ao ar livre' },
  { src: '/images/lifestyle-nature.jpg', alt: 'Meditação ao pôr do sol' },
  { src: '/images/lifestyle-wellness.jpg', alt: 'Prática de mindfulness' },
];

export function LifestyleGrid() {
  return (
    <section className="bg-cream pb-8">
      <div className="site-container">
        <div className="grid md:grid-cols-3 gap-3">
          {tiles.map((tile) => (
            <div key={tile.src} className="relative aspect-[4/3] overflow-hidden">
              <Image
                src={tile.src}
                alt={tile.alt}
                fill
                className="object-cover"
                sizes="(min-width: 768px) 33vw, 100vw"
              />
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}
