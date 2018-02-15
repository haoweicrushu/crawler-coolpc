if [ ! -f ./config/mongodb.conf ]; then
  echo "setting your mongodb config first"

  echo "vim ./config/mongodb.conf"
  cp ./config/mongodb.example ./config/mongodb.conf
fi

