# width
convert -quality 100 -resize 4096 origin_big.jpg origin_4k.jpg

convert -quality 100 -resize 2048 origin_light_big.jpg origin_light_2k.jpg
convert -quality 100 -resize 4096 origin_light_big.jpg origin_light_4k.jpg

for i in `ls *jpg`; do j=${i%.*}; convert -quality 100 i-jpg/$j.jpg i-png/$j.png ; done
for i in `ls *jpg`; do j=${i%.*}; convert -quality 100 i-jpg/$j.jpg i-web/$j.webp ; done