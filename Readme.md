将 Gitlab push事件推送到 mattermonst。

## Usage
```
./gitlab_webhook --port=9090 --webhook=http://mm.test.com/hooks/xxxxxxxxx >/dev/null 2>&1 &
```

## 使用服务的方式启动
新建 `/usr/lib/systemd/system/gitlabhook.service` 服务文件，如下：
```
[Unit]
Description=gitlab hook deamon
Wants=network-online.target
After=network-online.target

[Service]
Type=simple
Environment="MMHOOK=http://mm.test.com/hooks/xxxxxxxxx"
# 注意：要用 ${MMHOOK}，不要用$MMHOOK
# 参见：https://www.qiansw.com/how-to-use-environment-variables-in-systemd-1.html
ExecStart=/usr/local/bin/gitlab_webhook --port=9090 --webhook=${MMHOOK}

[Install]
WantedBy=multi-user.target
```

## 参考
- [gitlab-webhook-robot](https://github.com/EalenXie/gitlab-webhook-robot)
- [go-gitlab-webhook](https://github.com/soupdiver/go-gitlab-webhook)