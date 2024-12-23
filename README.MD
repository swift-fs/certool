# 证书自动部署

只支持阿里云 CDN.

## 参考文档

- [设置 CDN 域名证书](https://api.aliyun.com/document/Cdn/2018-05-10/SetCdnDomainSSLCertificate)
- [查询用户证书或者订单列表](https://api.aliyun.com/document/cas/2020-04-07/ListUserCertificateOrder)
- [删除证书](https://api.aliyun.com/document/cas/2020-04-07/DeleteUserCertificate)

## 修改配置文件

找到`manifest/config/config.example.yaml`文件修改配置后同目录创建一个`config.yaml`文件:

```yaml

logger:
  level: "all"
  stdout: true
  path: "./logs/"
  file: "{Y-m-d}.log"

aliyun:
  cdnAccessKeyId: "cdn的access_id"****
  cdnAccessKeySecret: "cdn的access_key"
  certAccessKeyId: "证书管理的access_id"
  certAccessKeySecret: "证书管理的access_key"

```

## 构建产物

下载[gf 命令行](https://goframe.org/pages/viewpage.action?pageId=1115782)后进行构建

```bash
gf build
```

会在 bin 目录下生成 linux 的二进制文件`certool`,也可以自行修改`hack/config.yaml`构建其他平台二进制文件.

## 使用

```bash
certool cert -domain=需要部署的cdn域名 -path=证书所在目录 -del=是否删除过期证书
```

- `domain`需要部署的 cdn 加速域名
- `path`证书所在的绝对路径,公钥名`fullchain/pem`,私钥名`privkey.pem`
- `del`是否删除阿里云上传证书的过期证书,默认 true

案例:

```bash

sudo ./certool cert -domain=cdn.xxx.com -path=/home/ecs-user/certs/xxx.com/cdn
```

### 最佳搭档

使用 1panel 的证书管理服务,证书保存到本地,在申请证书后执行该脚本.

### TODO

- [ ] 定时任务(暂时结合 1panel 没必要,毕竟自用)
