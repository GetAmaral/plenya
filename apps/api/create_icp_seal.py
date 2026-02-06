from PIL import Image, ImageDraw, ImageFont
import os

# Criar imagem para selo ICP-Brasil
width = 350
height = 100
bg_color = (0, 120, 0)  # Verde oficial ICP-Brasil
text_color = (255, 255, 255)  # Branco

# Criar imagem
img = Image.new('RGB', (width, height), bg_color)
draw = ImageDraw.Draw(img)

# Adicionar texto "ICP-BRASIL"
try:
    # Tentar usar fonte padrão
    font = ImageFont.truetype("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 40)
except:
    font = ImageFont.load_default()

text = "ICP-BRASIL"
# Calcular posição para centralizar
bbox = draw.textbbox((0, 0), text, font=font)
text_width = bbox[2] - bbox[0]
text_height = bbox[3] - bbox[1]
x = (width - text_width) // 2
y = (height - text_height) // 2 - 5

draw.text((x, y), text, fill=text_color, font=font)

# Salvar
img.save('icp-brasil-seal.png')
print("Selo ICP-Brasil criado com sucesso!")
