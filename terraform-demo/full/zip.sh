ZIP_PATH=$(pwd)/$1
SRC_PATH=$(pwd)/$2

rm $ZIP_PATH
cd $SRC_PATH

zip -r -X $ZIP_PATH . %1>/dev/null %2>/dev/null
echo "{ \"output_base64sha256\": \"$(cat "$ZIP_PATH" | shasum -a 256 | cut -d " " -f 1 | xxd -r -p | base64)\", \"output_md5\": \"$(cat "$ZIP_PATH" | md5)\", \"output_path\": \"$ZIP_PATH\" }"