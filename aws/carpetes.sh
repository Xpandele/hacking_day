#!/bin/bash

function createDirs() {
        for element in $@; do
                mkdir $element
        done
}

function createScript() {
        file=$1
        content=$2

        echo "#!/bin/bash" > $file
        echo $content >> $file
        chmod +x $file
}

# Creating youtubers folder and sub directories
mkdir Youtubers
cd Youtubers
youtubers=("Alias" "Llocs" "Anys")
createDirs ${youtubers[@]}

# Creating areas subdirectory
cd Alias
createScript "Vegetta777" "echo ves al gimnas"
createScript "elRubiusOMG" "echo ves a l'aula d'informatica 1"
createScript "Willyrex" "echo ves al labavo de nois davant de la sala de profes"
createScript "Mangelrogel" "echo ves a la sala de profes"
createScript "aLexBY11" "echo ves a l'aula d'informatica 2"
createScript "TheGrefg" "echo ves al labavo d'abaix"
createScript "Ampeterby7" "echo ves a consergeria"
createScript "Auronplay" "echo ves a secretaria"
createScript "Jordi_Wild" "echo ves al taller"
createScript "Fernanfloo" "echo ves als barracons"


