cat ./data/input.mp3 | ffmpeg -i pipe:0 -c:a libvorbis -b:a 320k -bufsize 320k -analyzeduration 0 -loglevel 0 -f ogg pipe: