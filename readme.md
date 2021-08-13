# SlackSSHLoginNotifier
## 概要
SSHでログインされたとき、Slackでそれを監視したい

## 目次
<!-- TOC -->

- [SlackSSHLoginNotifier](#slacksshloginnotifier)
    - [概要](#概要)
    - [目次](#目次)
    - [設定](#設定)

<!-- /TOC -->

## 設定

1. バイナリを適当に配置し、PATHを通す。

2. 以下を追記

`/etc/ssh/sshrc `

```
SSHSlackHookURL="HOOK URI" SlackSSHLoginNotifier
```

