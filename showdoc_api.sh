#! /bin/bash
#
# 文档说明： https://www.showdoc.cc/page/741656402509783 
# 
api_key="f8e2d90984fdbb582f0436e7b7b40dba1659615504" 			#api_key
api_token="2dd4c43df53e70c60eb074cbf0cd16f21147398481" 	#api_token
url="https://www.showdoc.cc/server/?s=/api/open/fromComments" #同步到的url。使用www.showdoc.cc的不需要修改，使用开源版的请修改
#
#
#
#
#
#
# 如果第一个参数是目录，则使用参数目录。若无，则使用脚本所在的目录。
if [[ -z "$1" ]] || [[ ! -d "$1" ]] ; then #目录判断，如果$1不是目录或者是空，则使用当前目录
	curren_dir=$(dirname $(readlink -f $0))
else
	curren_dir=$(cd $1; pwd)
fi
echo "$curren_dir"
 
 
# 递归搜索文件
searchfile() {
	
	old_IFS="$IFS"
	IFS=$'\n'            #IFS修改
	for chkfile in $1/*
	do
		filesize=`ls -l $chkfile | awk '{ print $5 }'`
		maxsize=$((1024*1024*1))  # 1M以下的文本文件才会被扫描
		if [[ -f "$chkfile" ]] &&  [ $filesize -le $maxsize ] && [[ -n $(file $chkfile | grep text) ]] ; then # 只对text文件类型操作
			echo "正在扫描 $chkfile"
			result=$(sed -n -e '/\/\*\*/,/\*\//p' $chkfile | grep showdoc) # 正则匹配
		if [ ! -z "$result" ] ; then 
					txt=$(sed -n -e '/\/\*\*/,/\*\//p' $chkfile)
					#echo "sed -n -e '/\/\*\*/,/\*\//p' $chkfile"
					#echo $result
					if  [[ $txt =~ "@url" ]] && [[ $txt =~ "@title" ]]; then
						echo -e "\033[32m $chkfile 扫描到内容 , 正在生成文档 \033[0m "
						# 通过接口生成文档
curl -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8'  "${url}" --data-binary @- <<CURL_DATA
from=shell&api_key=${api_key}&api_token=${api_token}&content=${txt}
CURL_DATA
					fi
			fi
		fi

		if [[ -d $chkfile ]] ; then
			searchfile $chkfile
		fi
	done
	IFS="$old_IFS"
}


#执行搜索
searchfile $curren_dir


#
sys=$(uname)
if [[ $sys =~ "MS"  ]] || [[ $sys =~ "MINGW"  ]] || [[ $sys =~ "win"  ]] ; then 
	read -s -n1 -p "按任意键继续 ... " # win环境下为照顾用户习惯，停顿一下
fi

