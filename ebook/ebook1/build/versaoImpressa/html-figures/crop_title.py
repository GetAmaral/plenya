from PIL import Image
im = Image.open('cap13-fig01/preview.png')
crop = im.crop((20, 50, 1400, 130))
crop.save('/tmp/cap13_title_crop.png')

orig = Image.open('cap13-fig01/_reference.png')
ocrop = orig.crop((20, 50, 1400, 130))
ocrop.save('/tmp/cap13_title_orig.png')

# Combined
W = 1380
out = Image.new("RGB", (W, 170), "white")
out.paste(ocrop, (0, 0))
out.paste(crop, (0, 85))
out.save('/tmp/cap13_title_compare.png')
