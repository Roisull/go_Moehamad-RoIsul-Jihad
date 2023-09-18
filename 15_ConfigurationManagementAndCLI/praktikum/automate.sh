#!/bin/bash
echo "======================"
echo "| Scripts Running... |" 
echo "======================"
# 1. membuat struktur folder
folder_name="$1 at $(date +'%a %b %d %H:%M:%S WIB %Y')"

mkdir "$folder_name"

mkdir "$folder_name/about_me/"
mkdir "$folder_name/about_me/personal/"
mkdir "$folder_name/about_me/professional/"

# 2. membuat fie facebook.txt dan linkedin.txt
echo "https://www.facebook.com/$2" > "$folder_name/about_me/personal/facebook.txt"
echo "https://www.linkedin.com/in/$3" > "$folder_name/about_me/professional/linkedin.txt"

# 3. membuat file list_of_my_friends.txt
mkdir "$folder_name/my_friends"
echo "
1) Achmad Miftahul - amulum
2) Achmad Rafiq - achmadrafiq97
3) Adiststi - adiststi
4) Agung - agungajin19
5) Azzahra - al7262
6) Charisma - fadzrichar
7) Farida - ulfarida
8 ) Garry - garryarielcussoy
9) Hamdi - hamdiranu
10) Hedy Gading - Gading09
11) Ilham - aamfatur
12) Lelianto - Lelianto
13) Mohammad - daffa99
14) Muhammad Fadhil - fabdulkarim
15) Ofbimon - bimon-alta
16) Purnama Syafitri - pipitmnr
17) Setyo - setyoyo07
18) Wildan - wiflash
19) Willy - sumarnowilly94
20) Woka - woka20
" > "$folder_name/my_friends/list_of_my_friends.txt"

# # 4. membuat file about_this_laptop.txt
mkdir "$folder_name/my_system_info/"
path="$folder_name/my_system_info/about_this_laptop.txt"

echo "My Username: $(whoami)" > "$path"
echo "With host: " >> "$path"
uname -a >> "$path"

# # 5. membuat file internet_connection.txt
ping -c 3 google.com > "$folder_name/my_system_info/internet_connection.txt"

# # Menampilkan struktur folder
tree "$folder_name"

