#!/bin/bash

# Option 1: PNG zu JPEG konvertieren (meist dramatische Größenreduktion)
for i in `find posts -type f ! -name "thumb-*.png" -name "*.png" | grep -v thumbs`;
do
  filefull="${i##*/}"
  file="${filefull%.[^.]*}"
  ext="${filefull:${#file} + 1}"
  dir=$(echo ${i} | sed -e "s/${file}\.${ext}//")
  source="${dir}${file}.${ext}"
  
  if [ ! -d ${dir}thumbs/ ];
  then
    mkdir ${dir}thumbs/
  fi
  
  # Als JPEG speichern für dramatische Größenreduktion
  dest="${dir}thumbs/${file}.jpg"
  
  if [ -f ${dest} ];
  then
    continue;
  fi;
  
  # PNG zu JPEG mit weißem Hintergrund (für Transparenz)
  convert ${source} -background white -flatten -quality 85 -strip ${dest};
  echo "convert ${source} -background white -flatten -quality 85 -strip ${dest};"
done

echo -e "\n=== webp Konvertierung ==="
# for i in posts/**/*.{jpg,jpeg,png,JPG,JPEG,PNG}; do
for i in `find posts -type f ! -name "*.svg" ! -name "*.webp" ! -name "*.gif" | grep -v thumbs`; do 
  # echo "i: $i"
  [ ! -f "$i" ] && continue
  
  filename=$(basename "$i")
  name="${filename%.*}"
  dir=$(dirname "$i")
  output="${dir}/${name}.webp"
  
  [ -f "$output" ] && continue
  
  # Dateigröße prüfen
  size=$(stat -f%z "$i" 2>/dev/null || stat -c%s "$i" 2>/dev/null)
  
  # Qualität basierend auf Dateigröße anpassen
  if [ "$size" -gt 1000000 ]; then  # > 1MB
    quality=75
  elif [ "$size" -gt 500000 ]; then # > 500KB
    quality=80
  else
    quality=85
  fi
  
  echo "Konvertiere: $i (${size} bytes) -> $output (quality: $quality)"
  cwebp -q $quality "$i" -o "$output"
done

# echo "Generate Thumbnails" 
# for i in `find posts -type f ! -name "*.svg" ! -name "*.webp" | grep -v thumbs`;
# do
#   # echo $i;
#   filefull="${i##*/}"
#   file="${filefull%.[^.]*}"
#   ext="${filefull:${#file} + 1}"
#   dir=$(echo ${i} | sed -e "s/${file}\.${ext}//")
#   source="${dir}${file}.${ext}"
#   if [ ! -d ${dir}thumbs/ ];
#   then
#     mkdir ${dir}thumbs/
#   fi
#   dest="${dir}thumbs/${file}-600x600.${ext}"
#   if [ -f ${dest} ];
#   then
#     continue;
#   fi;
#   convert ${source} -thumbnail 600x600 ${dest}
#   echo "convert ${source} -thumbnail 600x600 ${dest}";
# done


# Option 2: Wenn Sie PNG behalten möchten - mit pngquant (installieren mit: sudo apt install pngquant)
# for i in `find posts -type f ! -name "thumb-*.png" -name "*.png" | grep -v thumbs`;
# do
#   filefull="${i##*/}"
#   file="${filefull%.[^.]*}"
#   ext="${filefull:${#file} + 1}"
#   dir=$(echo ${i} | sed -e "s/${file}\.${ext}//")
#   source="${dir}${file}.${ext}"
#   
#   if [ ! -d ${dir}thumbs/ ];
#   then
#     mkdir ${dir}thumbs/
#   fi
#   dest="${dir}thumbs/${file}.${ext}"
#   
#   if [ -f ${dest} ];
#   then
#     continue;
#   fi;
#   
#   # PNG verlustbehaftet komprimieren (kann 50-80% Größenreduktion bringen)
#   pngquant --quality=65-80 --output ${dest} ${source};
# done

# Option 3: Moderne WebP Format (beste Komprimierung)
# for i in `find posts -type f ! -name "thumb-*.png" -name "*.png" | grep -v thumbs`;
# do
#   filefull="${i##*/}"
#   file="${filefull%.[^.]*}"
#   ext="${filefull:${#file} + 1}"
#   dir=$(echo ${i} | sed -e "s/${file}\.${ext}//")
#   source="${dir}${file}.${ext}"
#   
#   if [ ! -d ${dir}thumbs/ ];
#   then
#     mkdir ${dir}thumbs/
#   fi
#   dest="${dir}thumbs/${file}.webp"
#   
#   if [ -f ${dest} ];
#   then
#     continue;
#   fi;
#   
#   # Als WebP mit hoher Komprimierung
#   convert ${source} -quality 80 -define webp:lossless=false ${dest};
# done

# # Option 1: Einfache Qualitätsreduzierung (empfohlen für die meisten Fälle)
# for i in `find posts -type f ! -name "thumb-*.png" -name "*.png" | grep -v thumbs`;
# do
#   filefull="${i##*/}"
#   file="${filefull%.[^.]*}"
#   ext="${filefull:${#file} + 1}"
#   dir=$(echo ${i} | sed -e "s/${file}\.${ext}//")
#   source="${dir}${file}.${ext}"
  
#   if [ ! -d ${dir}thumbs/ ];
#   then
#     mkdir ${dir}thumbs/
#   fi
  
#   dest="${dir}thumbs/${file}.${ext}"
  
#   # if [ -f ${dest} ];
#   # then
#   #   continue;
#   # fi;
  
#   # Komprimierung mit reduzierter Qualität (80% ist ein guter Startwert)
#   echo "convert ${source} -quality 80 -strip ${dest};"
#   # convert ${source} -quality 60 -strip ${dest};
#   convert ${source} -quality 30 -strip -interlace Plane -resize '1000x1000>' ${dest};
#   # convert ${source} -strip -define png:compression-level=9 ${dest};
# done

# Option 2: Aggressivere Komprimierung
# convert ${source} -quality 60 -strip -interlace Plane ${dest};

# Option 3: Für PNG spezifisch optimiert
# convert ${source} -strip -define png:compression-level=9 ${dest};

# Option 4: Automatische Optimierung basierend auf Dateigröße
# convert ${source} -quality 85 -strip -resize '2000x2000>' ${dest};

####################################################################################################

# for i in `find posts -type f ! -name "thumb-*.png" -name "*.png" | grep -v thumbs`;
# do
#   filefull="${i##*/}"
#   file="${filefull%.[^.]*}"
#   ext="${filefull:${#file} + 1}"
#   dir=$(echo ${i} | sed -e "s/${file}\.${ext}//")
#   source="${dir}${file}.${ext}"
#   if [ ! -d ${dir}thumbs/ ];
#   then
#     mkdir ${dir}thumbs/
#   fi
#   dest="${dir}thumbs/${file}.${ext}"
#   if [ -f ${dest} ];
#   then
#     continue;
#   fi;
#   convert ${source} -thumbnail 200x200 ${dest};
# done



#################


#for i in `find static/images/galleries -type f ! -name "*-thumb.jpg" -name "*.jpg"`;
#do
#  echo $i;
#  if [ -f ${i%.*}-thumb.jpg ];
#  then
#    continue;
#  fi;
#  convert $i -thumbnail 300x300 ${i%.*}-thumb.jpg;
#done

# for i in `find static/images/galleries -type f ! -name "thumb-*.jpg" -name "*.jpg" | grep -v thumbs`;
# do
#   filefull="${i##*/}"
#   file="${filefull%.[^.]*}"
#   ext="${filefull:${#file} + 1}"
#   dir=$(echo ${i} | sed -e "s/${file}\.${ext}//")
#   source="${dir}${file}.${ext}"
#   if [ ! -d ${dir}thumbs/ ];
#   then
#     mkdir ${dir}thumbs/
#   fi
#   dest="${dir}thumbs/${file}.${ext}"
#   if [ -f ${dest} ];
#   then
#     continue;
#   fi;
#   convert ${source} -thumbnail 300x300 ${dest};
# done

# for i in `find static/img/posts -type f ! -name "*.svg" | grep -v thumbs`;
# do
#   # echo $i;
#   filefull="${i##*/}"
#   file="${filefull%.[^.]*}"
#   ext="${filefull:${#file} + 1}"
#   dir=$(echo ${i} | sed -e "s/${file}\.${ext}//")
#   source="${dir}${file}.${ext}"
#   if [ ! -d ${dir}thumbs/ ];
#   then
#     mkdir ${dir}thumbs/
#   fi
#   dest="${dir}thumbs/${file}.${ext}"
#   if [ -f ${dest} ];
#   then
#     continue;
#   fi;
#   convert ${source} -thumbnail 600x600 ${dest}
#   echo "convert ${source} -thumbnail 300x300 ${dest}";
# done
