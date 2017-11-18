mkdir /koolshare/qiandao;
cd /koolshare/qiandao;
wget -O qiandao-arm https://raw.githubusercontent.com/Caeseason/qiandao/master/v1.3/qiandao-arm;
wget -O qiandao.sh https://raw.githubusercontent.com/Caeseason/qiandao/master/v1.3/qiandao.sh;
wget -O cookie.conf https://raw.githubusercontent.com/Caeseason/qiandao/master/v1.3/cookie.conf;
chmod 755 /koolshare/qiandao/qiandao-arm;
chmod 755 /koolshare/qiandao/qiandao.sh;
cd /jffs/scripts
wget -O init-start https://raw.githubusercontent.com/Caeseason/qiandao/master/init-start;
chmod 755 /jffs/scripts/init-start;
./init-start;
