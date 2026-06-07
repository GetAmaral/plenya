from PIL import Image
im = Image.open('cap13-fig01/preview.png')
crop = im.crop((30, 830, 200, 920))
crop = crop.resize((crop.width*3, crop.height*3))
crop.save('/tmp/warning_zoom_mine.png')

orig = Image.open('cap13-fig01/_reference.png')
ocrop = orig.crop((30, 830, 200, 920))
ocrop = ocrop.resize((ocrop.width*3, ocrop.height*3))
ocrop.save('/tmp/warning_zoom_orig.png')

cmb = Image.new("RGB", (max(ocrop.width, crop.width)*2 + 30, max(ocrop.height, crop.height) + 30), "white")
cmb.paste(ocrop, (10, 10))
cmb.paste(crop, (ocrop.width + 30, 10))
cmb.save('/tmp/warning_compare.png')
print("done")
