<?php
$xmlData = "
< xml ><ToUserName><![CDATA[ad775b217]]>< /ToUserName >
< FromUserName ><![CDATA[tWy3zC3xUgQMR5coXif5SA]]>< /FromUserName >
< CreateTime >1366181013< /CreateTim e>
< MsgType ><![CDATA[text]]>< /MsgType >
< Content ><![CDATA[我的测试]]>< /Content >
< MsgId >5867702771251151243< /MsgId >
< /xml >";

//第一种发送方式，也是推荐的方式：
$url = 'http://api.zkkj168.com:8001/device/api?imei=123-111-333-222';  //接收xml数据的文件
//$url = 'http://127.0.0.1:8001/device';
$header[] = "Content-type: text/xml";        //定义content-type为xml,注意是数组
$ch = curl_init ($url);
curl_setopt($ch, CURLOPT_URL, $url);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_HTTPHEADER, $header);
curl_setopt($ch, CURLOPT_POST, 1);
curl_setopt($ch, CURLOPT_POSTFIELDS, $xmlData);
$response = curl_exec($ch);
if(curl_errno($ch)){
    print curl_error($ch);
    }
    curl_close($ch);
?>
