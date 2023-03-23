MYSQL_MESSAGE=$(cat << EOM
hello how are you

    ;alkdjf;alskjdf
EOM

)

echo $MYSQL_MESSAGE > mysql_message.txt
